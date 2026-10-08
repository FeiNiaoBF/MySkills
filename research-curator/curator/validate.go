package curator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func validateSemantic(r *Run) error {
	ids := map[string]string{}
	add := func(id, kind string) error {
		if _, ok := ids[id]; ok {
			return fmt.Errorf("duplicate ID %q", id)
		}
		ids[id] = kind
		return nil
	}
	if e := add(r.ID, "Goal"); e != nil {
		return e
	}
	for _, q := range r.Contract.Questions {
		if e := add(q.ID, "Question"); e != nil {
			return e
		}
	}
	if r.Contract.Output.TargetSources < 1 {
		return fmt.Errorf("contract.output.target_sources must be at least 1")
	}
	if r.Contract.Coverage.MinIndependentOrigins < 1 {
		return fmt.Errorf("contract.coverage.min_independent_origins must be at least 1")
	}
	sources := map[string]Source{}
	claims := map[string]Claim{}
	for _, s := range r.Sources {
		if e := add(s.ID, "Source"); e != nil {
			return e
		}
		sources[s.ID] = s
	}
	for _, c := range r.Claims {
		if e := add(c.ID, "Claim"); e != nil {
			return e
		}
		claims[c.ID] = c
	}
	for _, c := range r.Conclusions {
		if e := add(c.ID, "Conclusion"); e != nil {
			return e
		}
	}
	ref := func(id, kind string) error {
		if ids[id] != kind {
			return fmt.Errorf("%q is not a %s", id, kind)
		}
		return nil
	}
	date := func(s string) error {
		_, e := time.Parse(time.RFC3339, s)
		if e != nil {
			return fmt.Errorf("invalid RFC3339 date %q", s)
		}
		return nil
	}
	if e := date(r.Metadata.CreatedAt); e != nil {
		return e
	}
	if r.Metadata.CompletedAt != "" {
		if e := date(r.Metadata.CompletedAt); e != nil {
			return e
		}
		created, _ := time.Parse(time.RFC3339, r.Metadata.CreatedAt)
		completed, _ := time.Parse(time.RFC3339, r.Metadata.CompletedAt)
		if completed.Before(created) {
			return fmt.Errorf("completed_at precedes created_at")
		}
	}
	if (r.Metadata.Status == "finalized") != (r.Metadata.CompletedAt != "") {
		return fmt.Errorf("completed_at must be present exactly when status is finalized")
	}
	rejected := map[string]bool{}
	for _, item := range r.RejectedSources {
		s, ok := sources[item.ID]
		if !ok || s.Status == "selected" || rejected[item.ID] || item.Reason != s.Reason {
			return fmt.Errorf("invalid rejected_sources entry %s", item.ID)
		}
		rejected[item.ID] = true
	}
	if r.Metadata.Status == "finalized" {
		for _, s := range r.Sources {
			if s.Status != "selected" && !rejected[s.ID] {
				return fmt.Errorf("finalized run omits rejected source %s", s.ID)
			}
		}
	}
	for _, s := range r.Sources {
		u, e := url.Parse(s.URL)
		if e != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil {
			return fmt.Errorf("unsafe source URL %q", s.URL)
		}
		if _, e := CanonicalURL(s.URL); e != nil {
			return fmt.Errorf("invalid canonical source URL %q: %w", s.URL, e)
		}
		if e := date(s.RetrievedAt); e != nil {
			return e
		}
		if s.PublishedAt != "" {
			if e := date(s.PublishedAt); e != nil {
				return e
			}
		}
		if s.ContentHash != "" && s.ContentHash != ContentHash(s.Content) {
			return fmt.Errorf("source %s content hash mismatch", s.ID)
		}
		if s.Original && len(s.UpstreamIDs) > 0 {
			return fmt.Errorf("original source %s has upstream sources", s.ID)
		}
		if s.Status == "duplicate" && s.DuplicateOf == "" {
			return fmt.Errorf("duplicate %s lacks duplicate_of", s.ID)
		}
		if s.DuplicateOf != "" {
			if e := ref(s.DuplicateOf, "Source"); e != nil {
				return e
			}
			if s.DuplicateOf == s.ID {
				return fmt.Errorf("self duplicate %s", s.ID)
			}
			if s.Status != "duplicate" {
				return fmt.Errorf("duplicate_of requires duplicate status")
			}
		}
		seen := map[string]bool{}
		for _, id := range s.UpstreamIDs {
			if e := ref(id, "Source"); e != nil {
				return e
			}
			if id == s.ID || seen[id] {
				return fmt.Errorf("self/repeated upstream %s", id)
			}
			seen[id] = true
		}
	}
	for _, c := range r.Claims {
		for _, id := range c.QuestionIDs {
			if e := ref(id, "Question"); e != nil {
				return e
			}
		}
		for _, ev := range c.Evidence {
			if e := ref(ev.SourceID, "Source"); e != nil {
				return e
			}
			s := sources[ev.SourceID]
			if !strings.Contains(s.Content, ev.Quote) {
				return fmt.Errorf("claim %s quote absent from source %s", c.ID, s.ID)
			}
			if ev.Verification == "verified" && s.Verification != "verified" {
				return fmt.Errorf("verified evidence on unverified source %s", s.ID)
			}
		}
	}
	for _, c := range r.Conclusions {
		if c.Status == "verified" {
			for _, a := range r.Adjudications {
				if a.TargetType == "conclusion" && a.TargetID == c.ID && (a.Status != "resolved" || a.Outcome != "supports_conclusion") {
					return fmt.Errorf("verified conclusion %s has unresolved or adverse adjudication", c.ID)
				}
			}
		}
		for _, id := range c.ClaimIDs {
			if e := ref(id, "Claim"); e != nil {
				return e
			}
			if c.Status == "verified" && !verifiedClaim(claims[id], sources, r) {
				return fmt.Errorf("conclusion %s cites unsupported claim %s", c.ID, id)
			}
		}
	}
	for _, q := range r.Queries {
		if e := add(q.ID, "Query"); e != nil {
			return e
		}
		if e := ref(q.QuestionID, "Question"); e != nil {
			return e
		}
		if e := date(q.At); e != nil {
			return e
		}
		if strings.TrimSpace(q.Provider) == "" || strings.TrimSpace(q.Tool) == "" || strings.TrimSpace(q.RetrievalReference) == "" {
			return fmt.Errorf("query %s requires provider, tool and retrieval_reference", q.ID)
		}
		seenCandidates := map[string]bool{}
		for _, sid := range q.CandidateSourceIDs {
			if e := ref(sid, "Source"); e != nil {
				return e
			}
			if seenCandidates[sid] {
				return fmt.Errorf("query %s repeats candidate %s", q.ID, sid)
			}
			seenCandidates[sid] = true
		}
	}
	adjudicated := map[string]bool{}
	for _, a := range r.Adjudications {
		if e := add(a.ID, "Adjudication"); e != nil {
			return e
		}
		if a.TargetType != "claim" && a.TargetType != "conclusion" {
			return fmt.Errorf("adjudication %s has invalid target_type", a.ID)
		}
		if e := ref(a.TargetID, map[string]string{"claim": "Claim", "conclusion": "Conclusion"}[a.TargetType]); e != nil {
			return e
		}
		key := a.TargetType + "\x00" + a.TargetID
		if adjudicated[key] {
			return fmt.Errorf("multiple adjudications for conflict %s", a.TargetID)
		}
		adjudicated[key] = true
		if a.Status != "resolved" && a.Status != "unresolved" {
			return fmt.Errorf("adjudication %s has invalid status", a.ID)
		}
		if strings.TrimSpace(a.Rationale) == "" { return fmt.Errorf("adjudication %s requires rationale", a.ID) }
		if strings.TrimSpace(a.FinalEffect) == "" { return fmt.Errorf("adjudication %s requires final_effect", a.ID) }
		if a.Status == "resolved" {
			if len(a.Evidence) == 0 {
				return fmt.Errorf("resolved adjudication %s requires final_effect and evidence", a.ID)
			}
			if (a.TargetType == "claim" && a.Outcome != "supports_claim" && a.Outcome != "rejects_claim") || (a.TargetType == "conclusion" && a.Outcome != "supports_conclusion" && a.Outcome != "rejects_conclusion") {
				return fmt.Errorf("resolved adjudication %s has invalid outcome", a.ID)
			}
		} else if a.Outcome != "unresolved" {
			return fmt.Errorf("unresolved adjudication %s must have unresolved outcome", a.ID)
		}
		seenClaims := map[string]bool{}
		for _, cid := range a.ClaimIDs {
			if e := ref(cid, "Claim"); e != nil {
				return e
			}
			if seenClaims[cid] {
				return fmt.Errorf("adjudication %s repeats claim", a.ID)
			}
			seenClaims[cid] = true
		}
		seenSources := map[string]bool{}
		for _, sid := range a.SourceIDs {
			if e := ref(sid, "Source"); e != nil {
				return e
			}
			if seenSources[sid] {
				return fmt.Errorf("adjudication %s repeats source", a.ID)
			}
			seenSources[sid] = true
		}
		if a.TargetType == "claim" && !seenClaims[a.TargetID] {
			return fmt.Errorf("adjudication %s omits target claim", a.ID)
		}
		if len(seenClaims) == 0 || len(seenSources) < 2 {
			return fmt.Errorf("adjudication %s must identify conflicting claims and sources", a.ID)
		}
		targetClaims := map[string]bool{}
		if a.TargetType == "claim" {
			targetClaims[a.TargetID] = true
		} else {
			for _, conclusion := range r.Conclusions {
				if conclusion.ID == a.TargetID {
					for _, cid := range conclusion.ClaimIDs {
						targetClaims[cid] = true
					}
				}
			}
		}
		for cid := range seenClaims {
			if !targetClaims[cid] {
				return fmt.Errorf("adjudication %s includes claim outside its target", a.ID)
			}
		}
		if a.TargetType == "conclusion" {
			linked := false
			for _, edge := range r.Graph.Edges {
				if edge.To == a.TargetID && edge.Type == "contradicts" && seenClaims[edge.From] {
					linked = true
				}
			}
			if !linked {
				return fmt.Errorf("adjudication %s is not linked to a conclusion contradiction", a.ID)
			}
		}
		linkedSources := map[string]bool{}
		supported, contradicted := false, false
		for _, cid := range a.ClaimIDs {
			claim, ok := claims[cid]
			if !ok {
				continue
			}
			if a.TargetType == "claim" && cid != a.TargetID {
				continue
			}
			for _, ev := range claim.Evidence {
				if !seenSources[ev.SourceID] {
					continue
				}
				linkedSources[ev.SourceID] = true
				if ev.Relation == "supports" {
					supported = true
				}
				if ev.Relation == "contradicts" {
					contradicted = true
				}
			}
		}
		if a.TargetType == "conclusion" {
			for _, c := range r.Conclusions {
				if c.ID != a.TargetID {
					continue
				}
				for _, cid := range c.ClaimIDs {
					claim := claims[cid]
					for _, ev := range claim.Evidence {
						if seenSources[ev.SourceID] && ev.Relation == "supports" {
							supported = true
						}
						if seenSources[ev.SourceID] && ev.Relation == "contradicts" {
							contradicted = true
						}
					}
				}
			}
		}
		if len(linkedSources) != len(seenSources) {
			return fmt.Errorf("adjudication %s lists sources not connected to target claims", a.ID)
		}
		if !supported || !contradicted {
			return fmt.Errorf("adjudication %s does not cover both sides of a recorded conflict", a.ID)
		}
		verifiedDependency := false
		for _, ev := range a.Evidence {
			s, ok := sources[ev.SourceID]
			if !ok || !seenSources[ev.SourceID] || ev.Verification != "verified" || s.Verification != "verified" || s.Status != "selected" || ev.Quote == "" || ev.Locator == "" || !strings.Contains(s.Content, ev.Quote) {
				return fmt.Errorf("adjudication %s has invalid or non-selected evidence dependency", a.ID)
			}
			verifiedDependency = true
		}
		if a.Status == "resolved" && !verifiedDependency {
			return fmt.Errorf("resolved adjudication %s lacks verified selected evidence", a.ID)
		}
	}
	for _, ev := range r.Events {
		if ev.BeforeCount != nil || ev.AfterCount != nil || ev.CountScope != "" {
			if ev.BeforeCount == nil || ev.AfterCount == nil || ev.CountScope == "" {
				return fmt.Errorf("event counts require before_count, after_count and count_scope together")
			}
		}
		if e := add(ev.ID, "Event"); e != nil {
			return e
		}
		if e := date(ev.At); e != nil {
			return e
		}
	}
	for _, d := range r.Decisions {
		if e := add(d.ID, "Decision"); e != nil {
			return e
		}
		if e := ref(d.SourceID, "Source"); e != nil {
			return e
		}
		if e := date(d.At); e != nil {
			return e
		}
	}
	stageSeen := map[string]bool{}
	for _, st := range r.Stages {
		if stageSeen[st.Name] {
			return fmt.Errorf("repeated stage %s", st.Name)
		}
		stageSeen[st.Name] = true
		if st.Status == "verified" && len(st.EvidenceIDs) == 0 {
			return fmt.Errorf("verified stage %s has no evidence", st.Name)
		}
		for _, id := range st.EvidenceIDs {
			if ids[id] == "" {
				return fmt.Errorf("stage references missing evidence %s", id)
			}
		}
	}
	nodes := map[string]string{}
	for _, n := range r.Graph.Nodes {
		if nodes[n.ID] != "" {
			return fmt.Errorf("duplicate graph node %s", n.ID)
		}
		if ids[n.ID] != n.Type {
			return fmt.Errorf("graph node %s type mismatch", n.ID)
		}
		nodes[n.ID] = n.Type
	}
	for id, kind := range ids {
		if kind == "Event" || kind == "Decision" || kind == "Adjudication" {
			continue
		}
		if nodes[id] != kind {
			return fmt.Errorf("missing graph node %s", id)
		}
	}
	edges := map[string]bool{}
	key := func(from, to, kind string) string { return from + "\x00" + to + "\x00" + kind }
	for _, e := range r.Graph.Edges {
		a, b := nodes[e.From], nodes[e.To]
		valid := false
		switch e.Type {
		case "contains":
			valid = a == "Goal" && b == "Question"
		case "searched_by":
			valid = a == "Question" && b == "Query"
		case "candidate":
			valid = a == "Query" && b == "Source"
		case "addresses":
			valid = a == "Question" && b == "Claim"
		case "supports", "contradicts":
			valid = (a == "Source" && b == "Claim") || (a == "Claim" && b == "Conclusion")
		case "duplicates", "replaces":
			valid = a == "Source" && b == "Source"
		case "derives_from":
			valid = (a == "Source" && b == "Source") || (a == "Claim" && b == "Claim") || (a == "Conclusion" && b == "Claim")
		}
		if !valid || e.From == e.To {
			return fmt.Errorf("invalid edge %s %s %s", e.From, e.Type, e.To)
		}
		k := key(e.From, e.To, e.Type)
		if edges[k] {
			return fmt.Errorf("repeated graph edge")
		}
		edges[k] = true
		if a == "Source" && b == "Claim" && (e.Type == "supports" || e.Type == "contradicts") {
			matched := false
			for _, evidence := range claims[e.To].Evidence {
				if evidence.SourceID == e.From && evidence.Relation == e.Type {
					matched = true
					break
				}
			}
			if !matched {
				return fmt.Errorf("evidence edge lacks recorded quotation: %s %s %s", e.From, e.Type, e.To)
			}
		}
		if b == "Conclusion" && a == "Claim" {
			for _, conclusion := range r.Conclusions {
				if conclusion.ID != e.To {
					continue
				}
				if e.Type == "supports" && !contains(conclusion.ClaimIDs, e.From) {
					return fmt.Errorf("conclusion support absent from claim_ids")
				}
				if e.Type == "contradicts" && conclusion.Status == "verified" {
					a := resolvedConflict(r, "conclusion", e.To)
					if a == nil || a.Outcome != "supports_conclusion" {
						return fmt.Errorf("verified conclusion %s has unresolved or adverse adjudication", e.To)
					}
				}
			}
		}
		if a == "Conclusion" && e.Type == "derives_from" {
			for _, conclusion := range r.Conclusions {
				if conclusion.ID == e.From && !contains(conclusion.ClaimIDs, e.To) {
					return fmt.Errorf("conclusion derivation absent from claim_ids")
				}
			}
		}
		if e.Type == "addresses" && !contains(claims[e.To].QuestionIDs, e.From) {
			return fmt.Errorf("addresses edge absent from question_ids")
		}
		if a == "Source" && e.Type == "derives_from" && !contains(sources[e.From].UpstreamIDs, e.To) {
			return fmt.Errorf("provenance edge absent from upstream_ids")
		}
		if e.Type == "duplicates" && sources[e.From].DuplicateOf != e.To {
			return fmt.Errorf("duplicate edge differs from duplicate_of")
		}
	}
	for _, s := range r.Sources {
		for _, up := range s.UpstreamIDs {
			if !edges[key(s.ID, up, "derives_from")] {
				return fmt.Errorf("missing provenance edge %s→%s", s.ID, up)
			}
		}
		if s.DuplicateOf != "" && !edges[key(s.ID, s.DuplicateOf, "duplicates")] {
			return fmt.Errorf("missing duplicate edge for %s", s.ID)
		}
	}
	for _, q := range r.Queries {
		if !edges[key(q.QuestionID, q.ID, "searched_by")] {
			return fmt.Errorf("missing question query provenance %s", q.ID)
		}
		for _, sid := range q.CandidateSourceIDs {
			if !edges[key(q.ID, sid, "candidate")] {
				return fmt.Errorf("missing candidate edge %s→%s", q.ID, sid)
			}
		}
		for _, edge := range r.Graph.Edges {
			if edge.From == q.ID && edge.Type == "candidate" && !contains(q.CandidateSourceIDs, edge.To) {
				return fmt.Errorf("candidate graph edge absent from query %s", q.ID)
			}
			if edge.To == q.ID && edge.Type == "searched_by" && edge.From != q.QuestionID {
				return fmt.Errorf("query %s has an unrecorded question edge", q.ID)
			}
		}
	}
	for _, q := range r.Contract.Questions {
		if !edges[key(r.ID, q.ID, "contains")] {
			return fmt.Errorf("missing goal question edge %s", q.ID)
		}
	}
	for _, c := range r.Claims {
		for _, q := range c.QuestionIDs {
			if !edges[key(q, c.ID, "addresses")] {
				return fmt.Errorf("missing question claim edge %s→%s", q, c.ID)
			}
		}
		for _, ev := range c.Evidence {
			if !edges[key(ev.SourceID, c.ID, ev.Relation)] {
				return fmt.Errorf("missing %s evidence edge for %s", ev.Relation, c.ID)
			}
		}
	}
	for _, c := range r.Conclusions {
		for _, id := range c.ClaimIDs {
			if !edges[key(id, c.ID, "supports")] && !edges[key(c.ID, id, "derives_from")] {
				return fmt.Errorf("missing conclusion claim edge")
			}
		}
	}
	seenQ := map[string]bool{}
	for _, q := range r.Coverage.Questions {
		if e := ref(q.QuestionID, "Question"); e != nil {
			return e
		}
		if seenQ[q.QuestionID] {
			return fmt.Errorf("duplicate coverage question")
		}
		seenQ[q.QuestionID] = true
		for _, id := range q.ClaimIDs {
			if e := ref(id, "Claim"); e != nil {
				return e
			}
			if !contains(claims[id].QuestionIDs, q.QuestionID) {
				return fmt.Errorf("coverage claim answers another question")
			}
		}
	}
	if r.Coverage.Status == "verified" {
		computed := Analyze(r)
		if computed.Status != "verified" {
			return fmt.Errorf("verified coverage has unmet requirements")
		}
		if r.Coverage.SelectedSources != computed.SelectedSources || r.Coverage.TargetSources != computed.TargetSources || r.Coverage.TargetSourcesMet != computed.TargetSourcesMet || r.Coverage.IndependentSources != computed.IndependentSources || r.Coverage.MinIndependentOrigins != computed.MinIndependentOrigins {
			return fmt.Errorf("verified coverage summary differs from computed requirements")
		}
		if len(r.Coverage.Questions) != len(computed.Questions) {
			return fmt.Errorf("verified coverage missing questions")
		}
		for _, q := range r.Coverage.Questions {
			for _, cq := range computed.Questions {
				if q.QuestionID == cq.QuestionID {
					if q.IndependentSources != cq.IndependentSources {
						return fmt.Errorf("incorrect independent source count")
					}
					if len(q.ClaimIDs) != len(cq.ClaimIDs) {
						return fmt.Errorf("verified coverage claim IDs differ from evidence")
					}
					for _, id := range cq.ClaimIDs {
						if !contains(q.ClaimIDs, id) {
							return fmt.Errorf("verified coverage omits supported claim")
						}
					}
				}
			}
			if len(q.Gaps) > 0 {
				return fmt.Errorf("verified coverage contains gaps")
			}
		}
	}
	return nil
}
func contains(a []string, s string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}
func verifiedClaim(c Claim, sources map[string]Source, r *Run) bool {
	supported := false
	contradicted := false
	for _, ev := range c.Evidence {
		s := sources[ev.SourceID]
		if ev.Verification != "verified" || s.Verification != "verified" {
			continue
		}
		if ev.Relation == "contradicts" {
			contradicted = true
		}
		if ev.Relation == "supports" && s.Status == "selected" {
			supported = true
		}
	}
	if contradicted {
		a := resolvedClaimAdjudication(r, c.ID)
		if a == nil || a.Outcome != "supports_claim" {
			return false
		}
	}
	return supported
}

// ContentHash hashes exact retained text, not a semantic normalization.
func ContentHash(content string) string {
	h := sha256.Sum256([]byte(content))
	return hex.EncodeToString(h[:])
}
