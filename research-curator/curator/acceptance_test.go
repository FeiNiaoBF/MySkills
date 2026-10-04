package curator_test

import (
	"encoding/json"
	"researchcurator/curator"
	"testing"
)

func TestDedupFailurePreservesLedger(t *testing.T) {
	r := fixture()
	s := r.Sources[0]
	s.ID = "s2"
	addSource(r, s)
	r.Claims[0].Evidence[0].SourceID = "s2"
	r.Graph.Edges[0].From = "s2"
	r.Conclusions = []curator.Conclusion{{ID: "end", Text: "test", ClaimIDs: []string{"c1"}, Status: "verified"}}
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "end", Type: "Conclusion", Label: "test"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "c1", To: "end", Type: "supports"})
	if e := curator.Validate(r); e != nil {
		t.Fatal(e)
	}
	before, _ := json.Marshal(r)
	if e := curator.Deduplicate(r); e == nil {
		t.Fatal("dedup accepted unsupported verified conclusion")
	}
	after, _ := json.Marshal(r)
	if string(before) != string(after) {
		t.Fatal("failed dedup mutated caller ledger")
	}
}

func TestMetadataAndRejectionAcceptance(t *testing.T) {
	r := fixture()
	for _, status := range []string{"rejected", "duplicate", "superseded"} {
		s := r.Sources[0]
		s.ID = status
		s.Status = status
		s.Reason = "explicit " + status + " reason"
		if status == "duplicate" {
			s.DuplicateOf = "s1"
			r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: s.ID, To: "s1", Type: "duplicates"})
		}
		addSource(r, s)
	}
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	if r.Metadata.Status != "finalized" || r.Metadata.CompletedAt == "" || len(r.RejectedSources) != 3 {
		t.Fatalf("missing finalization metadata/index: %+v", r)
	}
	for i, id := range []string{"duplicate", "rejected", "superseded"} {
		if r.RejectedSources[i].ID != id || r.RejectedSources[i].Reason != "explicit "+id+" reason" {
			t.Fatal("index loses status/reason or unstable order")
		}
	}
	before, _ := json.Marshal(r)
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	after, _ := json.Marshal(r)
	if string(before) != string(after) {
		t.Fatal("finalize not idempotent")
	}
	r.RejectedSources = r.RejectedSources[:2]
	if e := curator.Validate(r); e == nil {
		t.Fatal("incomplete finalized rejection index accepted")
	}
}

func TestRequiredMetadataAndReasonsSchema(t *testing.T) {
	b, _ := json.Marshal(fixture())
	for _, field := range []string{"metadata", "rejected_sources", "fit_reason", "evidence_reason", "utility_reason"} {
		var value map[string]any
		json.Unmarshal(b, &value)
		if field == "metadata" || field == "rejected_sources" {
			delete(value, field)
		} else {
			delete(value["sources"].([]any)[0].(map[string]any), field)
		}
		data, _ := json.Marshal(value)
		if _, e := curator.Decode(data); e == nil {
			t.Fatalf("missing %s accepted", field)
		}
	}
	for name, mutate := range map[string]func(*curator.Run){
		"invalid creation":             func(r *curator.Run) { r.Metadata.CreatedAt = "yesterday" },
		"finalized without completion": func(r *curator.Run) { r.Metadata.Status = "finalized" },
		"progress with completion":     func(r *curator.Run) { r.Metadata.CompletedAt = "2026-01-01T00:00:00Z" },
		"completion before creation":   func(r *curator.Run) { r.Metadata.Status = "finalized"; r.Metadata.CompletedAt = "2025-01-01T00:00:00Z" },
		"selected rejection": func(r *curator.Run) {
			r.RejectedSources = []curator.RejectedSource{{ID: "s1", Reason: r.Sources[0].Reason}}
		},
		"null tools": func(r *curator.Run) { r.Metadata.Tools = nil },
	} {
		t.Run(name, func(t *testing.T) {
			r := fixture()
			mutate(r)
			if e := curator.Validate(r); e == nil {
				t.Fatal("malformed metadata accepted")
			}
		})
	}
}

func TestRequiredGoalQuestionClaimConnections(t *testing.T) {
	for _, kind := range []string{"contains", "addresses"} {
		r := fixture()
		out := []curator.Edge{}
		for _, edge := range r.Graph.Edges {
			if edge.Type != kind {
				out = append(out, edge)
			}
		}
		r.Graph.Edges = out
		if e := curator.Validate(r); e == nil {
			t.Fatalf("isolated %s accepted", kind)
		}
	}
	r := fixture()
	r.Contract.Questions = append(r.Contract.Questions, curator.Question{ID: "q2", Text: "another question"})
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "q2", Type: "Question", Label: "another"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "goal", To: "q2", Type: "contains"}, curator.Edge{From: "q2", To: "c1", Type: "addresses"})
	if e := curator.Validate(r); e == nil {
		t.Fatal("unrecorded question membership edge accepted")
	}
}

func TestRejectedContradictionCannotBeHidden(t *testing.T) {
	for _, status := range []string{"selected", "rejected", "duplicate", "superseded"} {
		t.Run(status, func(t *testing.T) {
			r := fixture()
			s := r.Sources[0]
			s.ID = "contra"
			s.Status = status
			s.Content = "Contradictory retained quote"
			s.URL = "https://example.org/contra"
			if status == "duplicate" {
				s.DuplicateOf = "s1"
				r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: s.ID, To: "s1", Type: "duplicates"})
			}
			addSource(r, s)
			r.Claims[0].Evidence = append(r.Claims[0].Evidence, curator.Evidence{SourceID: s.ID, Quote: s.Content, Locator: "paragraph", Relation: "contradicts", Verification: "verified"})
			r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: s.ID, To: "c1", Type: "contradicts"})
			if e := curator.Finalize(r); e != nil {
				t.Fatal(e)
			}
			if r.Coverage.Status != "unverified" || len(r.Coverage.Questions[0].Gaps) == 0 {
				t.Fatal("contradiction hidden by status")
			}
			r.Conclusions = []curator.Conclusion{{ID: "end", Text: "unjustified conclusion", ClaimIDs: []string{"c1"}, Status: "verified"}}
			r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "end", Type: "Conclusion", Label: "conclusion"})
			r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "c1", To: "end", Type: "supports"})
			if e := curator.Validate(r); e == nil {
				t.Fatal("verified conclusion despite unresolved contradiction")
			}
		})
	}
}

func TestEventCountsAreScopedAndPaired(t *testing.T) {
	before, after := 2, 1
	r := fixture()
	ev := curator.Event{ID: "count", At: r.Metadata.CreatedAt, Stage: "Deduplicate", Action: "deduplicate", Detail: "synthetic counts", BeforeCount: &before, AfterCount: &after, CountScope: "selected_sources"}
	rec := curator.Recorder{Run: r}
	if e := rec.RecordEvent(ev); e != nil {
		t.Fatal(e)
	}
	ev.ID = "bad"
	ev.AfterCount = nil
	if e := rec.RecordEvent(ev); e == nil {
		t.Fatal("partial event counts accepted")
	}
	ev.AfterCount = &after
	before = -1
	if e := rec.RecordEvent(ev); e == nil {
		t.Fatal("negative event count accepted")
	}
}
