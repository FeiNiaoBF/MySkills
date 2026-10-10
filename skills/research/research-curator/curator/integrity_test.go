package curator_test

import (
	"fmt"
	"researchcurator/curator"
	"strings"
	"testing"
)

func TestGraphRejectsEvidenceWithoutRecordedQuote(t *testing.T) {
	for _, relation := range []string{"supports", "contradicts"} {
		t.Run(relation, func(t *testing.T) {
			r := fixture()
			s := r.Sources[0]
			s.ID = "other"
			s.URL = "https://example.org/other"
			addSource(r, s)
			r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: s.ID, To: "c1", Type: relation})
			if err := curator.Validate(r); err == nil {
				t.Fatal("unquoted graph evidence accepted")
			}
		})
	}
}
func TestConclusionSupportMustMatchCitations(t *testing.T) {
	r := fixture()
	other := r.Claims[0]
	other.ID = "c2"
	r.Claims = append(r.Claims, other)
	r.Conclusions = []curator.Conclusion{{ID: "end", Text: "provisional conclusion", ClaimIDs: []string{"c1"}, Status: "unverified"}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "end", Type: "Conclusion", Label: "provisional conclusion"}, curator.Node{ID: "c2", Type: "Claim", Label: other.Text})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: other.QuestionIDs[0], To: "c2", Type: "addresses"}, curator.Edge{From: other.Evidence[0].SourceID, To: "c2", Type: "supports"}, curator.Edge{From: "c1", To: "end", Type: "supports"}, curator.Edge{From: "c2", To: "end", Type: "supports"})
	if err := curator.Validate(r); err == nil {
		t.Fatal("uncited conclusion support accepted")
	}
}
func TestConclusionContradictionBlocksVerificationAndCoverage(t *testing.T) {
	r := fixture()
	r.Conclusions = []curator.Conclusion{{ID: "end", Text: "contested conclusion", ClaimIDs: []string{"c1"}, Status: "unverified"}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "end", Type: "Conclusion", Label: "contested conclusion"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "c1", To: "end", Type: "supports"}, curator.Edge{From: "c1", To: "end", Type: "contradicts"})
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.Status != "unverified" || !strings.Contains(strings.Join(r.Coverage.Warnings, " "), "conclusion") {
		t.Fatal("explicit conclusion conflict not surfaced")
	}
	r.Conclusions[0].Status = "verified"
	if err := curator.Validate(r); err == nil {
		t.Fatal("contradicted conclusion verified")
	}
}
func TestSharedDAGProvenanceScalesWithoutExponentialTraversal(t *testing.T) {
	r := fixture()
	for i := 0; i < 42; i++ {
		s := r.Sources[0]
		s.ID = fmt.Sprintf("dag-%d", i)
		s.URL = "https://example.org/" + s.ID
		s.Status = "rejected"
		s.Original = false
		if i < 2 {
			s.UpstreamIDs = []string{"s1"}
		} else {
			s.UpstreamIDs = []string{fmt.Sprintf("dag-%d", i-1), fmt.Sprintf("dag-%d", i-2)}
		}
		addSource(r, s)
		for _, up := range s.UpstreamIDs {
			r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: s.ID, To: up, Type: "derives_from"})
		}
	}
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.Questions[0].IndependentSources != 1 {
		t.Fatal("shared original inflated corroboration")
	}
}
