package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"researchcurator/curator"
	"strings"
	"testing"
)

// Synthetic CLI fixture: no retrieval or research is claimed.
const synthetic = `{"version":"2.0","id":"g","metadata":{"created_at":"2026-01-01T00:00:00Z","completed_at":"","status":"in_progress","tools":[],"search_provenance":"synthetic; no retrieval"},"rejected_sources":[],"contract":{"question":"Synthetic CLI test","types":["docs"],"excludes":[],"freshness":"any","preferences":[],"output":{"target_sources":1},"coverage":{"min_independent_origins":1},"depth":"brief","questions":[{"id":"q","text":"test"}]},"sources":[{"id":"s","url":"https://example.org/","title":"Synthetic","type":"docs","original":true,"retrieved_at":"2026-01-01T00:00:00Z","content":"Synthetic quote </script><img src=x onerror=alert(1)>","upstream_ids":[],"freshness":"any","fit":1,"evidence":1,"utility":1,"fit_reason":"test","evidence_reason":"test","utility_reason":"test","status":"selected","reason":"test","verification":"verified"}],"claims":[{"id":"c","text":"Synthetic claim","question_ids":["q"],"evidence":[{"source_id":"s","quote":"Synthetic quote","locator":"test paragraph","relation":"supports","verification":"verified"}]}],"conclusions":[],"graph":{"nodes":[{"id":"g","type":"Goal","label":"test"},{"id":"q","type":"Question","label":"test"},{"id":"s","type":"Source","label":"test"},{"id":"c","type":"Claim","label":"test"}],"edges":[{"from":"s","to":"c","type":"supports"},{"from":"g","to":"q","type":"contains"},{"from":"q","to":"c","type":"addresses"}]},"queries":[],"events":[],"decisions":[],"adjudications":[],"stages":[],"coverage":{"status":"unverified","questions":[],"warnings":[]}}`

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
	if e := execute([]string{"finalize"}, strings.NewReader(`{"version":"2.0"}`), &out, &out); e == nil {
		t.Fatal("missing fields accepted")
	}
	if out.Len() != 0 {
		t.Fatal("invalid run produced output")
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
