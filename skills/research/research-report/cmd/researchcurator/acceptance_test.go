package main

import (
	"bytes"
	"os"
	"path/filepath"
	"researchcurator/curator"
	"researchcurator/visualizer"
	"strings"
	"testing"
)

func TestFinalizeAutomaticallyReportsExactSerializedLedger(t *testing.T) {
	for _, stdout := range []bool{false, true} {
		t.Run(map[bool]string{false: "adjacent default", true: "explicit stdout report"}[stdout], func(t *testing.T) {
			d := t.TempDir()
			output := filepath.Join(d, "finalized.json")
			report := filepath.Join(d, "report.html")
			args := []string{"finalize", "-out", output}
			if stdout {
				args = []string{"finalize", "-report", report}
			}
			var out bytes.Buffer
			if e := execute(args, strings.NewReader(synthetic), &out, &out); e != nil {
				t.Fatal(e)
			}
			data := out.Bytes()
			if !stdout {
				var e error
				data, e = os.ReadFile(output)
				if e != nil {
					t.Fatal(e)
				}
			}
			r, e := curator.Decode(data)
			if e != nil {
				t.Fatal(e)
			}
			if r.Metadata.Status != "finalized" || r.Metadata.CompletedAt == "" || r.Coverage.Status != "verified" {
				t.Fatal("not sealed")
			}
			expected, e := visualizer.Render(data)
			if e != nil {
				t.Fatal(e)
			}
			actual, e := os.ReadFile(report)
			if e != nil {
				t.Fatal(e)
			}
			if !bytes.Equal(actual, expected) {
				t.Fatal("report does not use identical finalized JSON")
			}
		})
	}
}

func TestFinalizeRejectsUnsafeOutputAndInvalidInput(t *testing.T) {
	d := t.TempDir()
	for _, args := range [][]string{{"finalize"}, {"finalize", "-report", "-"}, {"finalize", "-out", filepath.Join(d, "same"), "-report", filepath.Join(d, "same")}, {"validate", "-report", filepath.Join(d, "report.html")}} {
		var out bytes.Buffer
		if e := execute(args, strings.NewReader(synthetic), &out, &out); e == nil {
			t.Fatal("invalid report arguments accepted", args)
		}
		if out.Len() != 0 {
			t.Fatal("failed command emitted ledger")
		}
	}
	report := filepath.Join(d, "report.html")
	var out bytes.Buffer
	if e := execute([]string{"finalize", "-report", report}, strings.NewReader(`{"version":"1.0"}`), &out, &out); e == nil {
		t.Fatal("invalid input finalized")
	}
	if _, e := os.Stat(report); !os.IsNotExist(e) {
		t.Fatal("report created for invalid input")
	}
	// Reporting failure must not publish a JSON ledger or stdout.
	output := filepath.Join(d, "run.json")
	if e := execute([]string{"finalize", "-out", output, "-report", filepath.Join(d, "missing", "report.html")}, strings.NewReader(synthetic), &out, &out); e == nil {
		t.Fatal("report IO failure ignored")
	}
	if _, e := os.Stat(output); !os.IsNotExist(e) {
		t.Fatal("JSON written despite report failure")
	}
	if out.Len() != 0 {
		t.Fatal("failed finalization emitted stdout")
	}
}
