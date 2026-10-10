package curator

import (
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

// CanonicalURL removes fragments and known tracking parameters only.
func CanonicalURL(raw string) (string, error) {
	u, e := url.Parse(raw)
	if e != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") || u.User != nil {
		return "", fmt.Errorf("invalid HTTP(S) URL")
	}
	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)
	if (u.Scheme == "https" && u.Port() == "443") || (u.Scheme == "http" && u.Port() == "80") {
		u.Host = u.Hostname()
		if strings.Contains(u.Host, ":") {
			u.Host = "[" + u.Host + "]"
		}
	}
	u.Fragment = ""
	u.RawFragment = ""
	q, e := url.ParseQuery(u.RawQuery)
	if e != nil {
		return "", e
	}
	for k := range q {
		low := strings.ToLower(k)
		if strings.HasPrefix(low, "utm_") || low == "gclid" || low == "fbclid" {
			q.Del(k)
		}
	}
	u.RawQuery = q.Encode()
	if u.Path == "" {
		u.Path = "/"
	}
	return u.String(), nil
}

// Deduplicate marks exact duplicates in input order, preserving rejected and superseded decisions.
// No semantic similarity or provenance is inferred.
func Deduplicate(r *Run) error {
	if e := Validate(r); e != nil {
		return e
	}
	targetRun := r
	trial := *r
	trial.Sources = append([]Source{}, r.Sources...)
	trial.Graph.Edges = append([]Edge{}, r.Graph.Edges...)
	r = &trial
	urls, hashes := map[string]string{}, map[string]string{}
	for i := range r.Sources {
		s := &r.Sources[i]
		u, e := CanonicalURL(s.URL)
		if e != nil {
			return e
		}
		h := ContentHash(s.Content)
		s.ContentHash = h
		if s.Status == "duplicate" || s.Status == "rejected" || s.Status == "superseded" {
			continue
		}
		target := urls[u]
		reason := "exact canonical URL"
		if target == "" && s.Content != "" {
			target = hashes[h]
			reason = "exact retained-text SHA256"
		}
		if target != "" {
			s.Status = "duplicate"
			s.DuplicateOf = target
			s.Reason = "Duplicate of " + target + ": " + reason
			r.Graph.Edges = append(r.Graph.Edges, Edge{From: s.ID, To: target, Type: "duplicates"})
		} else {
			urls[u] = s.ID
			if s.Content != "" {
				hashes[h] = s.ID
			}
		}
	}
	// Previously asserted coverage may no longer be valid after deduplication.
	r.Coverage = Analyze(r)
	r.RejectedSources = rejectionIndex(r)
	if e := Validate(r); e != nil {
		return e
	}
	*targetRun = trial
	return nil
}

// Rank returns selected sources only; weights are contractual and ties are stable by ID.
func Rank(r *Run) []Source {
	out := []Source{}
	for _, s := range r.Sources {
		if s.Status == "selected" {
			out = append(out, s)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := Score(out[i]), Score(out[j])
		if a == b {
			return out[i].ID < out[j].ID
		}
		return a > b
	})
	return out
}
func Score(s Source) float64 { return .5*s.Fit + .3*s.Evidence + .2*s.Utility }

// Analyze checks coverage from recorded quotes and explicit provenance; it never fetches evidence.
func Analyze(r *Run) Coverage {
	result := Coverage{Status: "unverified", TargetSources: r.Contract.Output.TargetSources, MinIndependentOrigins: r.Contract.Coverage.MinIndependentOrigins, Questions: []QuestionCoverage{}, Warnings: []string{}}
	selectedIDs := map[string]bool{}
	for _, s := range r.Sources {
		if s.Status == "selected" {
			selectedIDs[s.ID] = true
		}
	}
	result.SelectedSources = len(selectedIDs)
	result.TargetSourcesMet = result.SelectedSources >= result.TargetSources
	sources := map[string]Source{}
	for _, s := range r.Sources {
		sources[s.ID] = s
	}
	warnings := map[string]bool{}
	conclusionIDs := map[string]bool{}
	for _, conclusion := range r.Conclusions {
		conclusionIDs[conclusion.ID] = true
	}
	for _, edge := range r.Graph.Edges {
		if edge.Type == "contradicts" && conclusionIDs[edge.To] {
			a := resolvedConflict(r, "conclusion", edge.To)
			if a == nil || a.Outcome != "supports_conclusion" {
				warnings["unresolved contradiction for conclusion "+edge.To] = true
			}
		}
	}
	adjacency := map[string][]string{}
	for _, edge := range r.Graph.Edges {
		if edge.Type == "derives_from" || edge.Type == "duplicates" || edge.Type == "replaces" {
			adjacency[edge.From] = append(adjacency[edge.From], edge.To)
		}
	}
	state := map[string]int{}
	var visit func(string, int)
	visit = func(id string, depth int) {
		if depth >= 256 {
			warnings["graph provenance depth budget exceeded"] = true
			return
		}
		if state[id] == 1 {
			warnings["graph provenance cycle involving "+id] = true
			return
		}
		if state[id] == 2 {
			return
		}
		state[id] = 1
		for _, to := range adjacency[id] {
			visit(to, depth+1)
		}
		state[id] = 2
	}
	for _, node := range r.Graph.Nodes {
		visit(node.ID, 0)
	}
	rootCache := map[string]map[string]bool{}
	operations := 0
	var roots func(string, map[string]bool) map[string]bool
	roots = func(id string, path map[string]bool) map[string]bool {
		out := map[string]bool{}
		operations++
		if operations > 100000 || len(path) >= 256 {
			warnings["provenance analysis budget exceeded; coverage remains unverified"] = true
			return out
		}
		if path[id] {
			warnings["provenance cycle involving "+id] = true
			return out
		}
		if cached, ok := rootCache[id]; ok {
			return cached
		}
		s, ok := sources[id]
		if !ok {
			warnings["missing upstream "+id] = true
			return out
		}
		path[id] = true
		defer delete(path, id)
		if s.DuplicateOf != "" {
			out = roots(s.DuplicateOf, path)
		} else if len(s.UpstreamIDs) == 0 {
			if s.Original {
				out[id] = true
			} else {
				warnings["unknown original evidence for "+id] = true
			}
		} else {
			for _, up := range s.UpstreamIDs {
				for root := range roots(up, path) {
					operations++
					if operations > 100000 {
						warnings["provenance analysis budget exceeded; coverage remains unverified"] = true
						return out
					}
					out[root] = true
				}
			}
		}
		rootCache[id] = out
		return out
	}
	// Scan all candidates so even rejected circular provenance remains visible.
	for _, s := range r.Sources {
		roots(s.ID, map[string]bool{})
	}
	allEvidence := map[string]bool{}
	complete := true
	for _, q := range r.Contract.Questions {
		c := QuestionCoverage{QuestionID: q.ID, ClaimIDs: []string{}, Gaps: []string{}}
		evidenceIDs := map[string]bool{}
		for _, claim := range r.Claims {
			if !contains(claim.QuestionIDs, q.ID) {
				continue
			}
			supported := false
			contradicted := false
			for _, ev := range claim.Evidence {
				s := sources[ev.SourceID]
				if ev.Verification != "verified" || s.Verification != "verified" || ev.Quote == "" || ev.Locator == "" || !strings.Contains(s.Content, ev.Quote) {
					continue
				}
				if ev.Relation == "contradicts" {
					contradicted = true
					continue
				}
				if ev.Relation == "supports" && s.Status == "selected" {
					supported = true
					evidenceIDs[s.ID] = true
					allEvidence[s.ID] = true
				}
			}
			if contradicted {
				adjudication := resolvedClaimAdjudication(r, claim.ID)
				if adjudication == nil {
					c.Gaps = append(c.Gaps, "unresolved contradictory evidence for "+claim.ID)
				}
				if adjudication != nil && adjudication.Outcome == "rejects_claim" {
					supported = false
					c.Gaps = append(c.Gaps, "adjudicated against claim "+claim.ID)
				}
				if adjudication != nil && adjudication.Outcome == "supports_claim" && supported {
					contradicted = false
				}
			}
			if supported {
				c.ClaimIDs = append(c.ClaimIDs, claim.ID)
			}
		}
		c.IndependentSources = countOrigins(evidenceIDs, roots)
		if len(c.ClaimIDs) == 0 {
			c.Gaps = append(c.Gaps, "no verified supported claim")
		}
		if c.IndependentSources == 0 {
			c.Gaps = append(c.Gaps, "no known independent original evidence")
		}
		if len(c.Gaps) > 0 {
			complete = false
		}
		result.Questions = append(result.Questions, c)
	}
	if len(r.Contract.Questions) == 0 {
		complete = false
		warnings["no coverage questions"] = true
	}
	result.IndependentSources = countOrigins(allEvidence, roots)
	if result.IndependentSources < r.Contract.Coverage.MinIndependentOrigins {
		complete = false
		warnings[fmt.Sprintf("minimum independent origins unmet: need %d independent evidence groups", r.Contract.Coverage.MinIndependentOrigins)] = true
	}
	if !result.TargetSourcesMet {
		complete = false
		warnings[fmt.Sprintf("target sources unmet: need %d selected sources", r.Contract.Output.TargetSources)] = true
	}
	for w := range warnings {
		result.Warnings = append(result.Warnings, w)
	}
	sort.Strings(result.Warnings)
	if len(result.Warnings) > 0 {
		complete = false
	}
	if complete {
		result.Status = "verified"
	}
	return result
}

func resolvedConflict(r *Run, targetType, targetID string) *ConflictAdjudication {
	for i := range r.Adjudications {
		a := &r.Adjudications[i]
		if a.TargetType == targetType && a.TargetID == targetID && a.Status == "resolved" {
			return a
		}
	}
	return nil
}
func resolvedClaimAdjudication(r *Run, claimID string) *ConflictAdjudication {
	return resolvedConflict(r, "claim", claimID)
}

// Sources sharing any upstream root belong to the same corroboration group.
func countOrigins(ids map[string]bool, roots func(string, map[string]bool) map[string]bool) int {
	groups := []map[string]bool{}
	keys := []string{}
	for id := range ids {
		keys = append(keys, id)
	}
	sort.Strings(keys)
	for _, id := range keys {
		set := map[string]bool{}
		for root := range roots(id, map[string]bool{}) {
			set[root] = true
		}
		if len(set) == 0 {
			continue
		}
		for i := 0; i < len(groups); {
			overlap := false
			for root := range set {
				if groups[i][root] {
					overlap = true
				}
			}
			if overlap {
				for root := range groups[i] {
					set[root] = true
				}
				groups = append(groups[:i], groups[i+1:]...)
				i = 0
			} else {
				i++
			}
		}
		groups = append(groups, set)
	}
	return len(groups)
}

// Finalize derives coverage and the rejection index; it does not invent research or stages.
func Finalize(r *Run) error {
	if e := Validate(r); e != nil {
		return e
	}
	trial := *r
	trial.Coverage = Analyze(r)
	trial.RejectedSources = rejectionIndex(r)
	trial.Metadata.Status = "finalized"
	if trial.Metadata.CompletedAt == "" {
		trial.Metadata.CompletedAt = time.Now().UTC().Format(time.RFC3339)
	}
	if e := Validate(&trial); e != nil {
		return e
	}
	*r = trial
	return nil
}

func rejectionIndex(r *Run) []RejectedSource {
	out := []RejectedSource{}
	for _, s := range r.Sources {
		if s.Status != "selected" {
			out = append(out, RejectedSource{ID: s.ID, Reason: s.Reason})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
