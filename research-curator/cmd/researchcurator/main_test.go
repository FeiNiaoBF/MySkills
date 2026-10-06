package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"researchcurator/curator"
	"strings"
	"testing"
	"time"
)

// Synthetic CLI fixture: no retrieval or research is claimed.
const synthetic = `{"version":"1.0","id":"g","metadata":{"created_at":"2026-01-01T00:00:00Z","completed_at":"","status":"in_progress","tools":[],"search_provenance":"synthetic; no retrieval"},"rejected_sources":[],"contract":{"question":"Synthetic CLI test","types":["docs"],"excludes":[],"freshness":"any","preferences":[],"quantity":1,"depth":"brief","questions":[{"id":"q","text":"test"}]},"sources":[{"id":"s","url":"https://example.org/","title":"Synthetic","type":"docs","original":true,"retrieved_at":"2026-01-01T00:00:00Z","content":"Synthetic quote </script><img src=x onerror=alert(1)>","upstream_ids":[],"freshness":"any","fit":1,"evidence":1,"utility":1,"fit_reason":"test","evidence_reason":"test","utility_reason":"test","status":"selected","reason":"test","verification":"verified"}],"claims":[{"id":"c","text":"Synthetic claim","question_ids":["q"],"evidence":[{"source_id":"s","quote":"Synthetic quote","locator":"test paragraph","relation":"supports","verification":"verified"}]}],"conclusions":[],"graph":{"nodes":[{"id":"g","type":"Goal","label":"test"},{"id":"q","type":"Question","label":"test"},{"id":"s","type":"Source","label":"test"},{"id":"c","type":"Claim","label":"test"}],"edges":[{"from":"s","to":"c","type":"supports"},{"from":"g","to":"q","type":"contains"},{"from":"q","to":"c","type":"addresses"}]},"queries":[],"events":[],"decisions":[],"stages":[],"coverage":{"status":"unverified","questions":[],"warnings":[]}}`

func TestCLIImportFinalizeRankRender(t *testing.T) {
	for _, cmd := range []string{"validate", "import", "dedup", "rank", "finalize", "render"} {
		t.Run(cmd, func(t *testing.T) {
			var out bytes.Buffer
			args := []string{cmd}
			if cmd == "finalize" {
				args = append(args, "-report", filepath.Join(t.TempDir(), "report.html"))
			}
			if e := execute(args, strings.NewReader(synthetic), &out, &out); e != nil {
				t.Fatal(e)
			}
			if cmd == "render" {
				if !strings.Contains(out.String(), "<!doctype html>") && !strings.Contains(out.String(), "<!DOCTYPE html>") {
					t.Fatal("no HTML")
				}
				if strings.Contains(out.String(), "Synthetic quote </script><img") {
					t.Fatal("source injection unescaped")
				}
			} else {
				r, e := curator.Decode(out.Bytes())
				if e != nil {
					t.Fatal(e)
				}
				if cmd == "finalize" && r.Coverage.Status != "verified" {
					t.Fatal("coverage not computed")
				}
			}
		})
	}
	var out bytes.Buffer
	if e := execute([]string{"finalize"}, strings.NewReader(`{"version":"1.0"}`), &out, &out); e == nil {
		t.Fatal("missing fields accepted")
	}
	if out.Len() != 0 {
		t.Fatal("invalid run produced output")
	}
}
func syntheticReport(t *testing.T) []byte {
	t.Helper()
	run, err := curator.Decode([]byte(synthetic))
	if err != nil {
		t.Fatal(err)
	}
	if err := curator.Finalize(run); err != nil {
		t.Fatal(err)
	}
	coverage := []curator.ReportCoverage{{QuestionID: "q", Status: "uncovered", ClaimIDs: []string{}, Gaps: []string{"synthetic test handoff"}}}
	report := curator.Report{ReportVersion: "1.0", Run: *run, Research: curator.ResearchRecord{Audience: "test reader", Purpose: "exercise CLI publication", OriginTypesRationale: "synthetic fixture only", MaxRounds: 8, LowGainWindow: 3, Rounds: []curator.ResearchRound{}, Coverage: coverage, Gaps: []string{"synthetic test handoff"}, Stop: curator.ResearchStop{Reason: "user_stopped", RoundIDs: []string{}, Rationale: "synthetic CLI test"}}, Article: &curator.Article{Title: "Synthetic CLI report", Language: "en", OutputForm: "article", Lead: "No research claim is made.", Sections: []curator.ArticleSection{{ID: "sec", Heading: "Test", Blocks: []curator.ArticleBlock{{ID: "block", Kind: "paragraph", Role: "context", Text: "Synthetic fixture.", ClaimIDs: []string{}, ConclusionIDs: []string{}}}}}}}
	data, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestCLIValidateReportAcceptsResearchHandoff(t *testing.T) {
	var report curator.Report
	if err := json.Unmarshal(syntheticReport(t), &report); err != nil {
		t.Fatal(err)
	}
	report.Article = nil
	report.Research.Stop = curator.ResearchStop{Reason: "in_progress", RoundIDs: []string{}, Rationale: "research handoff"}
	data, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := execute([]string{"validate-report"}, strings.NewReader(string(data)), &out, &out); err != nil {
		t.Fatal(err)
	}
	if out.String() != "valid report envelope 1.0\n" {
		t.Fatalf("unexpected validation output: %s", out.String())
	}
}

func TestCLIPublishAcceptsReportEnvelope(t *testing.T) {
	destination := filepath.Join(t.TempDir(), "synthetic-topic")
	var out bytes.Buffer
	if err := execute([]string{"publish", "-out", destination}, strings.NewReader(string(syntheticReport(t))), &out, &out); err != nil {
		t.Fatal(err)
	}
	if out.Len() != 0 {
		t.Fatal("publish wrote an unexpected stdout result")
	}
	page, err := os.ReadFile(filepath.Join(destination, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(page), "Synthetic CLI report") {
		t.Fatal("published report missing article")
	}
}

func TestCLIPublishRejectsInvalidEnvelopeWithoutOutput(t *testing.T) {
	destination := filepath.Join(t.TempDir(), "invalid-topic")
	var out bytes.Buffer
	if err := execute([]string{"publish", "-out", destination}, strings.NewReader(`{"report_version":"1.0"}`), &out, &out); err == nil {
		t.Fatal("accepted invalid report")
	}
	if _, err := os.Stat(destination); !os.IsNotExist(err) {
		t.Fatal("invalid report created output")
	}
}

func TestCLIRecordRoundClearsStaleStopAndUnsealsLedger(t *testing.T) {
	d := t.TempDir()
	input, output, record := filepath.Join(d, "report.json"), filepath.Join(d, "updated.json"), filepath.Join(d, "round.json")
	var prior curator.Report
	if err := json.Unmarshal(syntheticReport(t), &prior); err != nil {
		t.Fatal(err)
	}
	prior.Article = nil
	initial, err := json.Marshal(prior)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(input, initial, 0600); err != nil {
		t.Fatal(err)
	}
	round := curator.ResearchRound{ID: "r1", QueryIDs: []string{}, QuestionIDs: []string{}, SearchIntent: "discovery", SearchAngle: "failed first attempt", OriginTypesSearched: []string{}, AssessedSourceIDs: []string{}, RedundantSourceIDs: []string{}, NewClaimIDs: []string{}, NewOriginIDs: []string{}, NewContradictionRefs: []curator.EvidenceRef{}, NewQuestionIDs: []string{}, MaterialGain: "none", MaterialityReason: "retrieval returned no inspectable material", MaterialityEvidence: []curator.EvidenceRef{}, CoverageSnapshot: []curator.ReportCoverage{}, Gaps: []string{"retrieval unavailable"}, Status: "failed"}
	data, err := json.Marshal(round)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(record, data, 0600); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := execute([]string{"record-round", "-in", input, "-record", record, "-out", output}, strings.NewReader(""), &out, &out); err != nil {
		t.Fatal(err)
	}
	updated, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	var report curator.Report
	if err := json.Unmarshal(updated, &report); err != nil {
		t.Fatal(err)
	}
	if report.Research.Stop.Reason != "in_progress" || len(report.Research.Stop.RoundIDs) != 0 || report.Run.Metadata.Status != "in_progress" || report.Run.Metadata.CompletedAt != "" {
		t.Fatalf("new round did not invalidate old stop: stop=%+v metadata=%+v", report.Research.Stop, report.Run.Metadata)
	}
}

func TestCLIRecordCapturesQueryAndRetrievalTimes(t *testing.T) {
	d := t.TempDir()
	input := filepath.Join(d, "run.json")
	if err := os.WriteFile(input, []byte(synthetic), 0600); err != nil {
		t.Fatal(err)
	}
	for _, item := range []struct{ kind, name, data string }{
		{"query", "query.json", `{"id":"q1","question_id":"q","text":"synthetic query","at":"2000-01-01T00:00:00Z"}`},
		{"source", "source.json", `{"id":"s2","url":"https://example.org/lead","title":"Synthetic lead","type":"docs","original":true,"retrieved_at":"2000-01-01T00:00:00Z","content":"Unverified synthetic lead.","upstream_ids":[],"freshness":"any","fit":0.5,"evidence":0,"utility":0.2,"fit_reason":"synthetic","evidence_reason":"not inspected","utility_reason":"synthetic","status":"rejected","reason":"not inspected as evidence","verification":"unverified"}`},
	} {
		record := filepath.Join(d, item.name)
		output := input
		if err := os.WriteFile(record, []byte(item.data), 0600); err != nil {
			t.Fatal(err)
		}
		started := time.Now().UTC()
		var out bytes.Buffer
		if err := execute([]string{"record", "-kind", item.kind, "-record", record, "-in", input, "-out", output}, strings.NewReader(""), &out, &out); err != nil {
			t.Fatalf("record %s: %v", item.kind, err)
		}
		data, err := os.ReadFile(output)
		if err != nil {
			t.Fatal(err)
		}
		var run curator.Run
		if err := json.Unmarshal(data, &run); err != nil {
			t.Fatal(err)
		}
		var stamp string
		if item.kind == "query" {
			stamp = run.Queries[len(run.Queries)-1].At
		} else {
			stamp = run.Sources[len(run.Sources)-1].RetrievedAt
			foundNode := false
			for _, node := range run.Graph.Nodes {
				if node.ID == "s2" && node.Type == "Source" {
					foundNode = true
				}
			}
			if !foundNode {
				t.Fatal("recorded source is missing its required graph node")
			}
		}
		at, err := time.Parse(time.RFC3339, stamp)
		if err != nil || at.Before(started.Add(-time.Second)) || at.After(time.Now().UTC().Add(time.Second)) {
			t.Fatalf("%s timestamp was not captured at record time: %q (%v)", item.kind, stamp, err)
		}
	}
}

func TestCLIRecordAndFileOutput(t *testing.T) {
	d := t.TempDir()
	input := filepath.Join(d, "run.json")
	record := filepath.Join(d, "event.json")
	output := filepath.Join(d, "output.json")
	if e := os.WriteFile(input, []byte(synthetic), 0600); e != nil {
		t.Fatal(e)
	}
	if e := os.WriteFile(record, []byte(`{"id":"ev","at":"2026-01-01T00:00:00Z","stage":"Verify","action":"inspect","detail":"synthetic test"}`), 0600); e != nil {
		t.Fatal(e)
	}
	var out bytes.Buffer
	if e := execute([]string{"record", "-kind", "event", "-record", record, "-in", input, "-out", output}, strings.NewReader(""), &out, &out); e != nil {
		t.Fatal(e)
	}
	data, e := os.ReadFile(output)
	if e != nil {
		t.Fatal(e)
	}
	var r curator.Run
	if e = json.Unmarshal(data, &r); e != nil {
		t.Fatal(e)
	}
	if len(r.Events) != 1 || r.Events[0].ID != "ev" {
		t.Fatal("record not persisted")
	}
}
