package curator_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"researchcurator/curator"
	"strings"
	"testing"
)

// This fixture is synthetic test data, not a research example or retrieval receipt.
func fixture() *curator.Run {
	return &curator.Run{Version: "2.0", ID: "goal", Metadata: curator.Metadata{CreatedAt: "2026-01-01T00:00:00Z", CompletedAt: "", Status: "in_progress", Tools: []string{}, SearchProvenance: "synthetic fixture; no searches"}, RejectedSources: []curator.RejectedSource{}, Contract: curator.Contract{Question: "Synthetic test question", Types: []string{"documentation"}, Excludes: []string{}, Freshness: "any", Preferences: []string{}, Output: curator.OutputRequirement{TargetSources: 1}, Coverage: curator.CoverageRequirement{MinIndependentOrigins: 1}, Depth: "brief", Questions: []curator.Question{{ID: "q1", Text: "Synthetic coverage question"}}}, Sources: []curator.Source{{ID: "s1", URL: "https://example.org/doc", Title: "Synthetic source", Type: "documentation", Original: true, RetrievedAt: "2026-01-01T00:00:00Z", Content: "Synthetic quotation for testing only.", UpstreamIDs: []string{}, Freshness: "test", Fit: 1, Evidence: 1, Utility: 1, FitReason: "synthetic fit", EvidenceReason: "synthetic quote", UtilityReason: "synthetic utility", Status: "selected", Reason: "synthetic fixture", Verification: "verified"}}, Claims: []curator.Claim{{ID: "c1", Text: "Synthetic test claim", QuestionIDs: []string{"q1"}, Evidence: []curator.Evidence{{SourceID: "s1", Quote: "Synthetic quotation", Locator: "test paragraph 1", Relation: "supports", Verification: "verified"}}}}, Conclusions: []curator.Conclusion{}, Graph: curator.Graph{Nodes: []curator.Node{{ID: "goal", Type: "Goal", Label: "Test"}, {ID: "q1", Type: "Question", Label: "Test"}, {ID: "s1", Type: "Source", Label: "Test"}, {ID: "c1", Type: "Claim", Label: "Test"}}, Edges: []curator.Edge{{From: "s1", To: "c1", Type: "supports"}, {From: "goal", To: "q1", Type: "contains"}, {From: "q1", To: "c1", Type: "addresses"}}}, Queries: []curator.Query{}, Events: []curator.Event{}, Decisions: []curator.Decision{}, Adjudications: []curator.ConflictAdjudication{}, Stages: []curator.Stage{}, Coverage: curator.Coverage{Status: "unverified", Questions: []curator.QuestionCoverage{}, Warnings: []string{}}}
}
func TestDecodeRejectsMissingRequiredData(t *testing.T) {
	if _, err := curator.Decode([]byte(`{"version":"2.0","id":"empty"}`)); err == nil {
		t.Fatal("missing data accepted")
	}
}
func TestFinalizeAndRecorderRoundTrip(t *testing.T) {
	r := fixture()
	if err := curator.Finalize(r); err != nil {
		t.Fatal(err)
	}
	if r.Coverage.Status != "verified" || r.Coverage.Questions[0].IndependentSources != 1 {
		t.Fatalf("wrong coverage: %+v", r.Coverage)
	}
	if len(r.Stages) != 0 {
		t.Fatal("invented stages")
	}
	rec := curator.Recorder{Run: r}
	if err := rec.RecordEvent(curator.Event{ID: "ev1", At: "2026-01-01T00:00:00Z", Stage: "Verify", Action: "inspect", Detail: "synthetic test"}); err != nil {
		t.Fatal(err)
	}
	p := filepath.Join(t.TempDir(), "run.json")
	if err := rec.Save(p); err != nil {
		t.Fatal(err)
	}
	b, e := os.ReadFile(p)
	if e != nil {
		t.Fatal(e)
	}
	got, e := curator.Decode(b)
	if e != nil {
		t.Fatal(e)
	}
	if len(got.Events) != 1 || got.Events[0].ID != "ev1" {
		t.Fatal("event lost")
	}
	if err := rec.RecordEvent(got.Events[0]); err == nil {
		t.Fatal("duplicate event accepted")
	}
	if len(r.Events) != 1 {
		t.Fatal("failed record mutated run")
	}
}
func TestShapeAndGraphRejectMalformedData(t *testing.T) {
	for name, mutate := range map[string]func(*curator.Run){"null array": func(r *curator.Run) { r.Events = nil }, "missing target": func(r *curator.Run) { r.Graph.Edges[0].To = "absent" }, "reversed support": func(r *curator.Run) { r.Graph.Edges[0].From = "c1"; r.Graph.Edges[0].To = "s1" }, "duplicate ID": func(r *curator.Run) { r.Sources[0].ID = "q1" }, "missing quote": func(r *curator.Run) { r.Claims[0].Evidence[0].Quote = "fabricated" }, "unsafe URL": func(r *curator.Run) { r.Sources[0].URL = "javascript:alert(1)" }, "unsupported verified stage": func(r *curator.Run) {
		r.Stages = []curator.Stage{{Name: "Discovery", Status: "verified", EvidenceIDs: []string{}, Reason: "invented"}}
	}, "unsupported coverage": func(r *curator.Run) {
		r.Sources[0].Verification = "unverified"
		r.Claims[0].Evidence[0].Verification = "unverified"
		r.Coverage.Status = "verified"
	}} {
		t.Run(name, func(t *testing.T) {
			r := fixture()
			mutate(r)
			if err := curator.Validate(r); err == nil {
				t.Fatal("malformed run accepted")
			}
		})
	}
	b, _ := json.Marshal(fixture())
	b = []byte(strings.Replace(string(b), `"id":"goal"`, `"id":"goal","surprise":true`, 1))
	if _, e := curator.Decode(b); e == nil {
		t.Fatal("unknown property accepted")
	}
	if e := curator.Finalize(&curator.Run{Version: "2.0"}); e == nil {
		t.Fatal("missing required data finalized")
	}
}
func addSource(r *curator.Run, s curator.Source) {
	r.Sources = append(r.Sources, s)
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: s.ID, Type: "Source", Label: s.Title})
}
func TestExactDedupAndStableRanking(t *testing.T) {
	r := fixture()
	s := r.Sources[0]
	s.ID = "s2"
	s.URL += "?utm_source=test#anchor"
	addSource(r, s)
	s.ID = "s3"
	s.URL = "https://example.org/another"
	addSource(r, s)
	if e := curator.Deduplicate(r); e != nil {
		t.Fatal(e)
	}
	for _, s := range r.Sources[1:] {
		if s.Status != "duplicate" || s.DuplicateOf != "s1" {
			t.Fatalf("not deduped: %+v", s)
		}
	}
	if len(curator.Rank(r)) != 1 {
		t.Fatal("duplicates ranked")
	}
	if e := curator.Deduplicate(r); e != nil {
		t.Fatal("dedup not idempotent", e)
	}
	r = fixture()
	s = r.Sources[0]
	s.ID = "s0"
	s.URL = "https://example.org/different"
	s.Content = "Different exact text"
	addSource(r, s)
	ranked := curator.Rank(r)
	if ranked[0].ID != "s0" {
		t.Fatal("ties not deterministic")
	}
	if curator.Score(s) != 1 {
		t.Fatal("score changed")
	}
}
func TestSharedOriginalIsNotIndependentCorroboration(t *testing.T) {
	r := fixture()
	r.Contract.Coverage.MinIndependentOrigins = 2
	s := r.Sources[0]
	s.ID = "s2"
	s.URL = "https://example.org/syndicated"
	s.Original = false
	s.UpstreamIDs = []string{"s1"}
	addSource(r, s)
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "s2", To: "s1", Type: "derives_from"}, curator.Edge{From: "s2", To: "c1", Type: "supports"})
	ev := r.Claims[0].Evidence[0]
	ev.SourceID = "s2"
	r.Claims[0].Evidence = append(r.Claims[0].Evidence, ev)
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	if r.Coverage.Status != "unverified" || r.Coverage.Questions[0].IndependentSources != 1 {
		t.Fatal("shared upstream counted independently", r.Coverage)
	}
	r.Sources[0].Original = false
	r.Sources[0].UpstreamIDs = []string{"s2"}
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "s1", To: "s2", Type: "derives_from"})
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	if r.Coverage.Questions[0].IndependentSources != 0 || !strings.Contains(strings.Join(r.Coverage.Warnings, " "), "cycle") {
		t.Fatal("cycle hidden", r.Coverage)
	}
}

func TestJSONAmbiguityAndCoverageTampering(t *testing.T) {
	b, _ := json.Marshal(fixture())
	for _, bad := range []string{strings.Replace(string(b), `"version":"2.0"`, `"version":"2.0","version":"2.0"`, 1), string(b) + ` {}`, strings.Replace(string(b), `"target_sources":1`, `"target_sources":1.5`, 1), strings.Replace(string(b), `"fit":1`, `"fit":2`, 1)} {
		if _, e := curator.Decode([]byte(bad)); e == nil {
			t.Fatal("ambiguous or invalid JSON accepted")
		}
	}
	r := fixture()
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	r.Coverage.Questions[0].ClaimIDs = []string{}
	if e := curator.Validate(r); e == nil {
		t.Fatal("forged coverage accepted")
	}
}
func TestContradictionKeepsCoverageUnverified(t *testing.T) {
	r := fixture()
	s := r.Sources[0]
	s.ID = "s2"
	s.URL = "https://example.org/contradiction"
	s.Content = "Synthetic contradictory quotation"
	addSource(r, s)
	r.Claims[0].Evidence = append(r.Claims[0].Evidence, curator.Evidence{SourceID: "s2", Quote: "Synthetic contradictory", Locator: "test paragraph", Relation: "contradicts", Verification: "verified"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "s2", To: "c1", Type: "contradicts"})
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	if r.Coverage.Status != "unverified" || len(r.Coverage.Questions[0].Gaps) == 0 {
		t.Fatal("contradiction hidden")
	}
}
func TestClaimProvenanceCycleIsFlagged(t *testing.T) {
	r := fixture()
	c := r.Claims[0]
	c.ID = "c2"
	r.Claims = append(r.Claims, c)
	r.Graph.Nodes = append(r.Graph.Nodes, curator.Node{ID: "c2", Type: "Claim", Label: "test"})
	r.Graph.Edges = append(r.Graph.Edges, curator.Edge{From: "q1", To: "c2", Type: "addresses"}, curator.Edge{From: "s1", To: "c2", Type: "supports"}, curator.Edge{From: "c1", To: "c2", Type: "derives_from"}, curator.Edge{From: "c2", To: "c1", Type: "derives_from"})
	if e := curator.Finalize(r); e != nil {
		t.Fatal(e)
	}
	if r.Coverage.Status != "unverified" || !strings.Contains(strings.Join(r.Coverage.Warnings, " "), "cycle") {
		t.Fatal("claim cycle hidden")
	}
}

func TestCanonicalizationPreservesMeaningfulQuery(t *testing.T) {
	got, e := curator.CanonicalURL("https://EXAMPLE.org:443/doc?version=2&utm_campaign=test#top")
	if e != nil {
		t.Fatal(e)
	}
	if got != "https://example.org/doc?version=2" {
		t.Fatal(got)
	}
	r := fixture()
	r.Sources[0].Content = ""
	r.Claims = []curator.Claim{}
	r.Graph.Nodes = r.Graph.Nodes[:3]
	r.Graph.Edges = []curator.Edge{{From: "goal", To: "q1", Type: "contains"}}
	s := r.Sources[0]
	s.ID = "s2"
	s.URL = "https://example.org/other"
	addSource(r, s)
	if e := curator.Deduplicate(r); e != nil {
		t.Fatal(e)
	}
	if r.Sources[1].Status != "selected" {
		t.Fatal("empty content deduplicated")
	}
}
