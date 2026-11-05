package curator_test

import (
	"researchcurator/curator"
	"testing"
)

func TestRecorderQueryAddsInspectableGraphProvenance(t *testing.T) {
	r := fixture()
	rec := curator.Recorder{Run: r}
	q := curator.Query{ID: "q-search", QuestionID: "q1", Text: "test search", At: r.Metadata.CreatedAt, Provider: "web search", Tool: "search_tool", RetrievalReference: "retrieval-123", CandidateSourceIDs: []string{"s1"}}
	if err := rec.RecordQuery(q); err != nil {
		t.Fatal(err)
	}
	if len(r.Queries) != 1 || len(r.Graph.Nodes) != 5 || len(r.Graph.Edges) != 5 {
		t.Fatalf("query provenance not added: %+v", r)
	}
	if err := curator.Validate(r); err != nil {
		t.Fatal(err)
	}
	beforeNodes, beforeEdges := len(r.Graph.Nodes), len(r.Graph.Edges)
	q.ID = "bad-query"
	q.Provider = ""
	if err := rec.RecordQuery(q); err == nil {
		t.Fatal("invalid query accepted")
	}
	if len(r.Queries) != 1 || len(r.Graph.Nodes) != beforeNodes || len(r.Graph.Edges) != beforeEdges {
		t.Fatal("failed query mutated ledger")
	}
}
