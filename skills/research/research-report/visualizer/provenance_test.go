package visualizer_test

import (
	"researchcurator/visualizer"
	"strings"
	"testing"
)

func TestReportIncludesAdjudicationAndQueryProvenanceViews(t *testing.T) {
	input := []byte(`{"version":"2.0","queries":[{"id":"query-1","provider":"search provider","tool":"web_search","retrieval_reference":"run-42","candidate_source_ids":["src-1"]}],"adjudications":[{"id":"adj-1","target_type":"claim","target_id":"claim-1","claim_ids":["claim-1"],"source_ids":["src-1","src-2"],"status":"resolved","rationale":"Explicit rationale","evidence":[{"source_id":"src-1","quote":"exact evidence","locator":"p. 2","verification":"verified"}],"outcome":"supports_claim","final_effect":"Retain bounded claim"}]}`)
	html, err := visualizer.Render(input)
	if err != nil {
		t.Fatal(err)
	}
	page := string(html)
	for _, want := range []string{"Conflict adjudications", "Query → candidate provenance", "adjudications", "candidate_source_ids", "retrieval_reference", "final_effect"} {
		if !strings.Contains(page, want) {
			t.Errorf("report omits %q", want)
		}
	}
}
