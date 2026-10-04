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
		for _, id := range c.ClaimIDs {
			if e := ref(id, "Claim"); e != nil {
				return e
			}
			if c.Status == "verified" && !verifiedClaim(claims[id], sources) {
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
		if kind == "Query" || kind == "Event" || kind == "Decision" {
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
					return fmt.Errorf("verified conclusion %s has unresolved contradiction", e.To)
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
func verifiedClaim(c Claim, sources map[string]Source) bool {
	supported := false
	for _, ev := range c.Evidence {
		s := sources[ev.SourceID]
		if ev.Verification != "verified" || s.Verification != "verified" {
			continue
		}
		if ev.Relation == "contradicts" {
			return false
		}
		if ev.Relation == "supports" && s.Status == "selected" {
			supported = true
		}
	}
	return supported
}

// ContentHash hashes exact retained text, not a semantic normalization.
func ContentHash(content string) string {
	h := sha256.Sum256([]byte(content))
	return hex.EncodeToString(h[:])
}
