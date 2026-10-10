package curator_test

import (
	"encoding/json"
	"researchcurator/curator"
	"strings"
	"testing"
)

func addContradictingSource(r *curator.Run) string {
	s := r.Sources[0]
	s.ID = "s2"
	s.URL = "https://example.org/opposing"
	s.Content = "Opposing quotation retained for conflict adjudication."
	s.Status = "rejected"
	s.Reason = "rejected for low utility, not because evidence is false"
	s.Utility = 0
	s.UtilityReason = "outside preferred output format"
	addSource(r, s)
	r.Claims[0].Evidence = append(r.Claims[0].Evidence, curator.Evidence{SourceID: "s2", Quote: "Opposing quotation retained", Locator: "section 2", Relation: "contradicts", Verification: "verified"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "s2", To: "c1", Type: "contradicts"})
	return s.ID
}
func validResolution(sourceID string) curator.ConflictAdjudication {
	return curator.ConflictAdjudication{ID: "adj1", TargetType: "claim", TargetID: "c1", ClaimIDs: []string{"c1"}, SourceIDs: []string{"s1", sourceID}, Status: "resolved", Rationale: "The claim is bounded to the cited primary specification; the opposing source discusses a different scope.", Evidence: []curator.AdjudicationEvidence{{SourceID: "s1", Quote: "Synthetic quotation for testing only.", Locator: "section 1", Verification: "verified"}}, Outcome: "supports_claim", FinalEffect: "Retain c1 only in the documented scope."}
}
func TestUnresolvedClaimConflictCannotBeHiddenByRejectingSource(t *testing.T) {
	r := fixture()
	addContradictingSource(r)
	r.Contract.Coverage.MinIndependentOrigins = 1
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.Status != "unverified" || len(r.Coverage.Questions[0].Gaps) == 0 {
		t.Fatal("rejected source silently erased contradiction")
	}
	r.Adjudications = []curator.ConflictAdjudication{validResolution("s2")}
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.Status != "verified" {
		t.Fatalf("explicit supported resolution not applied: %+v", r.Coverage)
	}
	r.Adjudications = []curator.ConflictAdjudication{}
	r.Sources[1].Status = "rejected"
	if err := curator.Validate(r); err == nil {
		t.Fatal("stale verified coverage bypassed unresolved conflict")
	}
	if got := curator.Analyze(r); got.Status != "unverified" {
		t.Fatal("source rejection erased the conflict", got)
	}
}
func TestConflictResolutionRequiresRationaleEvidenceAndParticipants(t *testing.T) {
	cases := map[string]func(*curator.Run){
		"no rationale": func(r *curator.Run) {
			a := validResolution("s2")
			a.Rationale = ""
			r.Adjudications = []curator.ConflictAdjudication{a}
		},
		"no evidence": func(r *curator.Run) {
			a := validResolution("s2")
			a.Evidence = nil
			r.Adjudications = []curator.ConflictAdjudication{a}
		},
		"wrong participant": func(r *curator.Run) {
			a := validResolution("s2")
			a.SourceIDs = []string{"s1"}
			r.Adjudications = []curator.ConflictAdjudication{a}
		},
		"fabricated quote": func(r *curator.Run) {
			a := validResolution("s2")
			a.Evidence[0].Quote = "invented"
			r.Adjudications = []curator.ConflictAdjudication{a}
		},
		"rejected-only evidence": func(r *curator.Run) {
			a := validResolution("s2")
			a.Evidence[0].SourceID = "s2"
			a.Evidence[0].Quote = "Opposing quotation retained"
			r.Adjudications = []curator.ConflictAdjudication{a}
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			r := fixture()
			addContradictingSource(r)
			r.Contract.Coverage.MinIndependentOrigins = 1
			mutate(r)
			if err := curator.Validate(r); err == nil {
				t.Fatal("invalid adjudication accepted")
			}
		})
	}
}
func TestRejectedSourceAloneNeverResolvesConflict(t *testing.T) {
	r := fixture()
	addContradictingSource(r)
	r.Contract.Coverage.MinIndependentOrigins = 1
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	for _, outcome := range []string{"supports_claim", "rejects_claim"} {
		r.Adjudications = []curator.ConflictAdjudication{{ID: "adj", TargetType: "claim", TargetID: "c1", ClaimIDs: []string{"c1"}, SourceIDs: []string{"s1", "s2"}, Status: "resolved", Rationale: "rejected source", Evidence: []curator.AdjudicationEvidence{}, Outcome: outcome, FinalEffect: "resolved"}}
		if err := curator.Validate(r); err == nil {
			t.Fatal("rejection masqueraded as adjudication")
		}
	}
}
func TestQuantityRequirementsAreDistinct(t *testing.T) {
	r := fixture()
	r.Contract.Output.TargetSources = 2
	r.Contract.Coverage.MinIndependentOrigins = 1
	s := r.Sources[0]
	s.ID = "s2"
	s.URL = "https://example.org/second"
	s.Content = "Second original source with a distinct quotation."
	addSource(r, s)
	r.Claims[0].Evidence = append(r.Claims[0].Evidence, curator.Evidence{SourceID: "s2", Quote: "Second original source", Locator: "p1", Relation: "supports", Verification: "verified"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "s2", To: "c1", Type: "supports"})
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.IndependentSources != 2 || r.Coverage.SelectedSources != 2 || !r.Coverage.TargetSourcesMet {
		t.Fatalf("wrong split semantics: %+v", r.Coverage)
	}
	r.Contract.Output.TargetSources = 2
	r.Contract.Coverage.MinIndependentOrigins = 3
	r.Metadata.Status = "in_progress"
	r.Metadata.CompletedAt = ""
	r.Coverage = curator.Coverage{Status: "unverified", Questions: []curator.QuestionCoverage{}, Warnings: []string{}}
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if !r.Coverage.TargetSourcesMet || r.Coverage.IndependentSources != 2 || r.Coverage.Status != "unverified" {
		t.Fatal("target count incorrectly substituted for independent origins", r.Coverage)
	}
	r.Contract.Output.TargetSources = 3
	r.Contract.Coverage.MinIndependentOrigins = 1
	r.Metadata.Status = "in_progress"
	r.Metadata.CompletedAt = ""
	r.Coverage = curator.Coverage{Status: "unverified", Questions: []curator.QuestionCoverage{}, Warnings: []string{}}
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.Status != "unverified" || r.Coverage.TargetSourcesMet {
		t.Fatal("source target not enforced separately")
	}
	b, _ := json.Marshal(r)
	var obj map[string]any
	_ = json.Unmarshal(b, &obj)
	contract := obj["contract"].(map[string]any)
	if _, old := contract["quantity"]; old {
		t.Fatal("legacy ambiguous quantity remains")
	}
	legacy := strings.Replace(string(b), `"version":"2.0"`, `"version":"1.0"`, 1)
	if _, err := curator.Decode([]byte(legacy)); err == nil {
		t.Fatal("old run schema bypassed the split contract")
	}
}
func TestQueryCandidateProvenanceIsTypedAndComplete(t *testing.T) {
	r := fixture()
	r.Queries = []curator.Query{{ID: "query1", QuestionID: "q1", Text: "search synthetic", At: r.Metadata.CreatedAt, Provider: "web-search", Tool: "search_tool", RetrievalReference: "search-run-42", CandidateSourceIDs: []string{"s1"}}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "query1", Type: "Query", Label: "search synthetic"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "q1", To: "query1", Type: "searched_by"}, curator.Edge{From: "query1", To: "s1", Type: "candidate"})
	if err := curator.Validate(r); err != nil {
		t.Fatal(err)
	}
	for name, mutate := range map[string]func(*curator.Run){"no search edge": func(x *curator.Run) { x.Graph.Edges = x.Graph.Edges[1:] }, "no candidate edge": func(x *curator.Run) { x.Graph.Edges = x.Graph.Edges[:len(x.Graph.Edges)-1] }, "missing candidate": func(x *curator.Run) { x.Queries[0].CandidateSourceIDs = []string{} }, "missing provider": func(x *curator.Run) { x.Queries[0].Provider = "" }, "missing retrieval reference": func(x *curator.Run) { x.Queries[0].RetrievalReference = "" }} {
		t.Run(name, func(t *testing.T) {
			x := fixture()
			x.Queries = append([]curator.Query{}, r.Queries...)
			x.Graph.Nodes = append([]curator.Node{}, r.Graph.Nodes...)
			x.Graph.Edges = append([]curator.Edge{}, r.Graph.Edges...)
			mutate(x)
			if err := curator.Validate(x); err == nil {
				t.Fatal("incomplete provenance accepted")
			}
		})
	}
}
func TestVerifiedConclusionCannotBypassUnresolvedAdjudication(t *testing.T) {
	r := fixture()
	addContradictingSource(r)
	r.Contract.Coverage.MinIndependentOrigins = 1
	r.Conclusions = []curator.Conclusion{{ID: "end", Text: "claim conclusion", ClaimIDs: []string{"c1"}, Status: "verified"}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "end", Type: "Conclusion", Label: "conclusion"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "c1", To: "end", Type: "supports"})
	if err := curator.Validate(r); err == nil {
		t.Fatal("verified conclusion bypassed unresolved conflict")
	}
	r.Adjudications = []curator.ConflictAdjudication{validResolution("s2")}
	if err := curator.Validate(r); err != nil {
		t.Fatal(err)
	}
	r.Adjudications[0].Status = "unresolved"
	r.Adjudications[0].Outcome = "unresolved"
	r.Conclusions[0].Status = "unverified"
	if err := curator.Validate(r); err != nil {
		t.Fatal(err)
	}
	r.Conclusions[0].Status = "verified"
	if err := curator.Validate(r); err == nil {
		t.Fatal("unresolved record bypassed verified conclusion guard")
	}
}
func TestConclusionConflictRequiresItsOwnAdjudication(t *testing.T) {
	r := fixture()
	addContradictingSource(r)
	claimAdj := validResolution("s2")
	r.Adjudications = []curator.ConflictAdjudication{claimAdj}
	r.Conclusions = []curator.Conclusion{{ID: "end", Text: "bounded conclusion", ClaimIDs: []string{"c1"}, Status: "verified"}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "end", Type: "Conclusion", Label: "end"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "c1", To: "end", Type: "supports"}, curator.Edge{From: "c1", To: "end", Type: "contradicts"})
	if err := curator.Validate(r); err == nil {
		t.Fatal("claim adjudication incorrectly resolved conclusion conflict")
	}
	ca := claimAdj
	ca.ID = "adj-conclusion"
	ca.TargetType = "conclusion"
	ca.TargetID = "end"
	ca.Outcome = "supports_conclusion"
	ca.FinalEffect = "Retain conclusion with scope qualification"
	r.Adjudications = append(r.Adjudications, ca)
	if err := curator.Validate(r); err != nil {
		t.Fatal("separate conclusion adjudication rejected:", err)
	}
	r.Adjudications[1].Outcome = "rejects_conclusion"
	if err := curator.Validate(r); err == nil {
		t.Fatal("adverse conclusion adjudication allowed verified conclusion")
	}
}

func TestConflictAndProvenanceSchemaFieldsRequired(t *testing.T) {
	r := fixture()
	addContradictingSource(r)
	r.Contract.Coverage.MinIndependentOrigins = 1
	r.Queries = []curator.Query{{ID: "query1", QuestionID: "q1", Text: "q", At: r.Metadata.CreatedAt, Provider: "p", Tool: "t", RetrievalReference: "ref", CandidateSourceIDs: []string{"s1"}}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "query1", Type: "Query", Label: "q"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "q1", To: "query1", Type: "searched_by"}, curator.Edge{From: "query1", To: "s1", Type: "candidate"})
	r.Adjudications = []curator.ConflictAdjudication{validResolution("s2")}
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(r)
	for _, field := range []string{"adjudications", "output", "coverage", "provider", "tool", "retrieval_reference", "candidate_source_ids"} {
		if !strings.Contains(string(b), `"`+field+`"`) {
			t.Fatalf("missing %s", field)
		}
	}
}
