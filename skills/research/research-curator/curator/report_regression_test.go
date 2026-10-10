package curator_test

import (
	"bytes"
	"encoding/json"
	"os"
	"researchcurator/curator"
	"testing"
)

func TestReportPreservesStrictLedgerBoundary(t *testing.T) {
	valid := reportJSON(t)
	for name, data := range map[string][]byte{
		"unknown embedded field": bytes.Replace(valid, []byte(`"run":{`), []byte(`"run":{"unexpected":true,`), 1),
		"missing embedded array": bytes.Replace(valid, []byte(`"queries":[],`), nil, 1),
		"invalid utf8":           bytes.Replace(valid, []byte("general reader"), []byte{'x', 255}, 1),
	} {
		t.Run(name, func(t *testing.T) {
			if bytes.Equal(data, valid) {
				t.Fatal("test mutation did not apply")
			}
			if _, err := curator.DecodeReport(data); err == nil {
				t.Fatal("accepted invalid original ledger bytes")
			}
		})
	}
}

func TestReportRejectsInconsistentRoundBookkeeping(t *testing.T) {
	for name, mutate := range map[string]func(*curator.Report){
		"query for another question": func(r *curator.Report) {
			r.Run.Contract.Questions = append(r.Run.Contract.Questions, curator.Question{ID: "q2", Text: "Other question"})
			r.Run.Graph.Nodes = append(r.Run.Graph.Nodes, curator.Node{ID: "q2", Type: "Question", Label: "Other"})
			r.Run.Graph.Edges = append(r.Run.Graph.Edges, curator.Edge{From: r.Run.ID, To: "q2", Type: "contains"})
			r.Research.Coverage = append(r.Research.Coverage, curator.ReportCoverage{QuestionID: "q2", Status: "uncovered", ClaimIDs: []string{}, Gaps: []string{"not searched"}})
			r.Run.Queries[0].QuestionID = "q2"
		},
		"duplicate redundant": func(r *curator.Report) { r.Research.Rounds[0].RedundantSourceIDs = []string{"s1", "s1"} },
		"duplicate assessed":  func(r *curator.Report) { r.Research.Rounds[0].AssessedSourceIDs = []string{"s1", "s1"} },
		"repeated new claim": func(r *curator.Report) {
			for i := 0; i < 2; i++ {
				r.Research.Rounds[i].NewClaimIDs = []string{"c1"}
				r.Research.Rounds[i].MaterialGain = "minor"
			}
		},
		"new origin already assessed": func(r *curator.Report) {
			r.Research.Rounds[1].NewOriginIDs = []string{"s1"}
			r.Research.Rounds[1].MaterialGain = "minor"
		},
		"new claim not assessed": func(r *curator.Report) {
			r.Research.Rounds[0].NewClaimIDs = []string{"c1"}
			r.Research.Rounds[0].AssessedSourceIDs = []string{}
			r.Research.Rounds[0].RedundantSourceIDs = []string{}
			r.Research.Rounds[0].MaterialGain = "minor"
		},
		"none with new claim": func(r *curator.Report) { r.Research.Rounds[0].NewClaimIDs = []string{"c1"} },
		"reused query":        func(r *curator.Report) { r.Research.Rounds[1].QueryIDs = r.Research.Rounds[0].QueryIDs },
	} {
		t.Run(name, func(t *testing.T) {
			r := saturationReport(t)
			if err := curator.ValidateReport(r); err != nil {
				t.Fatal(err)
			}
			mutate(r)
			if err := curator.ValidateReport(r); err == nil {
				t.Fatal("accepted inconsistent round")
			}
		})
	}
}

func TestReportAcceptsExhaustedResourceBudgetBeforeRoundLimit(t *testing.T) {
	r := saturationReport(t)
	r.Research.Rounds = r.Research.Rounds[:1]
	r.Research.Stop = curator.ResearchStop{Reason: "budget_exhausted", RoundIDs: []string{"r1"}, Rationale: "Declared tool budget ended"}
	b := encodeReport(t, r)
	for _, kind := range []string{"time", "tool_calls", "cost"} {
		t.Run(kind, func(t *testing.T) {
			data := bytes.Replace(b, []byte(`"max_rounds":8`), []byte(`"max_rounds":8,"resource_budget":{"kind":"`+kind+`","unit":"test units","limit":2,"used":2}`), 1)
			if _, err := curator.DecodeReport(data); err != nil {
				t.Fatalf("honest resource stop rejected: %v", err)
			}
			notExhausted := bytes.Replace(data, []byte(`"used":2`), []byte(`"used":1`), 1)
			if _, err := curator.DecodeReport(notExhausted); err == nil {
				t.Fatal("premature budget stop accepted")
			}
		})
	}
}

func TestPublicationRequiresSealedLedger(t *testing.T) {
	b, err := os.ReadFile("../examples/report.json")
	if err != nil {
		t.Fatal(err)
	}
	r, err := curator.DecodeReport(b)
	if err != nil {
		t.Fatal(err)
	}
	if err := curator.ValidatePublication(r); err != nil {
		t.Fatal(err)
	}
	r.Run.Metadata.Status = "in_progress"
	r.Run.Metadata.CompletedAt = ""
	if err := curator.ValidatePublication(r); err == nil {
		t.Fatal("published an unsealed ledger")
	}
}

func TestSaturationRejectsInvalidWindowWithoutPanic(t *testing.T) {
	r := saturationReport(t)
	r.Research.LowGainWindow = -1
	if curator.EvaluateSaturation(r).Eligible {
		t.Fatal("invalid policy accepted")
	}
}

func TestSaturationRequiresAssessedAndCurrentCoverage(t *testing.T) {
	for name, mutate := range map[string]func(*curator.Report){
		"future evidence": func(r *curator.Report) {
			s := r.Run.Sources[0]
			s.ID = "s2"
			s.URL = "https://example.org/unrelated"
			r.Run.Sources = append(r.Run.Sources, s)
			r.Run.Graph.Nodes = append(r.Run.Graph.Nodes, curator.Node{ID: "s2", Type: "Source", Label: "unrelated"})
			for i := range r.Research.Rounds {
				r.Research.Rounds[i].AssessedSourceIDs = []string{"s2"}
				r.Research.Rounds[i].RedundantSourceIDs = []string{}
			}
		},
		"stale final coverage": func(r *curator.Report) {
			r.Research.Rounds[2].CoverageSnapshot[0].Status = "uncovered"
			r.Research.Rounds[2].CoverageSnapshot[0].ClaimIDs = []string{}
			r.Research.Rounds[2].CoverageSnapshot[0].Gaps = []string{"core question unanswered"}
		},
		"blank reason":                 func(r *curator.Report) { r.Research.Rounds[2].MaterialityReason = "  " },
		"blank angle":                  func(r *curator.Report) { r.Research.Rounds[0].SearchAngle = "  " },
		"unexplained redundant source": func(r *curator.Report) { r.Research.Rounds[2].RedundantSourceIDs = []string{"s1"} },
	} {
		t.Run(name, func(t *testing.T) {
			r := saturationReport(t)
			r.Research.Stop = curator.ResearchStop{Reason: "saturated", RoundIDs: []string{"r1", "r2", "r3"}, Rationale: "test"}
			if err := curator.ValidateReport(r); err != nil {
				t.Fatal(err)
			}
			mutate(r)
			if err := curator.ValidateReport(r); err == nil {
				t.Fatal("accepted inconsistent saturation")
			}
		})
	}
}

func TestReportRejectsBlankGapsAndEmptyArticle(t *testing.T) {
	b, err := os.ReadFile("../examples/report.json")
	if err != nil {
		t.Fatal(err)
	}
	for name, mutate := range map[string]func(*curator.Report){
		"empty section": func(r *curator.Report) { r.Article.Sections[0].Blocks = []curator.ArticleBlock{} },
		"blank list": func(r *curator.Report) {
			b := &r.Article.Sections[0].Blocks[0]
			b.Kind = "list"
			b.Text = ""
			b.Items = []string{" "}
		},
		"blank gap": func(r *curator.Report) { r.Research.Gaps = []string{" "} },
		"blank coverage gap": func(r *curator.Report) {
			r.Research.Coverage[0].Status = "partial"
			r.Research.Coverage[0].Gaps = []string{" "}
		},
	} {
		t.Run(name, func(t *testing.T) {
			r, err := curator.DecodeReport(b)
			if err != nil {
				t.Fatal(err)
			}
			mutate(r)
			if err := curator.ValidatePublication(r); err == nil {
				t.Fatal("accepted empty disclosure/content")
			}
		})
	}
}

func TestResourceLimitPrecedesSaturation(t *testing.T) {
	r := saturationReport(t)
	r.Research.Stop = curator.ResearchStop{Reason: "saturated", RoundIDs: []string{"r1", "r2", "r3"}, Rationale: "test"}
	r.Research.ResourceBudget = &curator.ResourceBudget{Kind: "tool_calls", Unit: "calls", Limit: 1, Used: 999}
	for _, reason := range []string{"saturated", "user_stopped", "retrieval_blocked", "in_progress"} {
		r.Research.Stop.Reason = reason
		if err := curator.ValidateReport(r); err == nil {
			t.Fatalf("resource exhaustion reported as %s", reason)
		}
	}
	r.Research.Stop = curator.ResearchStop{Reason: "budget_exhausted", RoundIDs: []string{"r3"}, Rationale: "Record the actual overrun, do not conceal it"}
	if err := curator.ValidateReport(r); err != nil {
		t.Fatalf("cannot disclose actual measured overrun: %v", err)
	}
}

func TestRoundFindingsMustMatchScopeAndHistory(t *testing.T) {
	for name, mutate := range map[string]func(*curator.Report){
		"unrelated new claim": func(r *curator.Report) {
			r.Run.Contract.Questions = append(r.Run.Contract.Questions, curator.Question{ID: "q2", Text: "other"})
			r.Run.Graph.Nodes = append(r.Run.Graph.Nodes, curator.Node{ID: "q2", Type: "Question", Label: "other"})
			r.Run.Graph.Edges = append(r.Run.Graph.Edges, curator.Edge{From: r.Run.ID, To: "q2", Type: "contains"})
			c := r.Run.Claims[0]
			c.ID = "c2"
			c.QuestionIDs = []string{"q2"}
			r.Run.Claims = append(r.Run.Claims, c)
			r.Run.Graph.Nodes = append(r.Run.Graph.Nodes, curator.Node{ID: "c2", Type: "Claim", Label: "other"})
			r.Run.Graph.Edges = append(r.Run.Graph.Edges, curator.Edge{From: "q2", To: "c2", Type: "addresses"}, curator.Edge{From: "s1", To: "c2", Type: "supports"})
			r.Research.Coverage = append(r.Research.Coverage, curator.ReportCoverage{QuestionID: "q2", Status: "uncovered", ClaimIDs: []string{}, Gaps: []string{"not assessed"}})
			r.Research.Rounds[0].NewClaimIDs = []string{"c2"}
			r.Research.Rounds[0].MaterialGain = "minor"
		},
		"claim already in materiality": func(r *curator.Report) {
			r.Research.Rounds[0].CoverageSnapshot[0] = curator.ReportCoverage{QuestionID: "q1", Status: "partial", ClaimIDs: []string{}, Gaps: []string{"pending"}}
			r.Research.Rounds[0].MaterialityEvidence = []curator.EvidenceRef{{ClaimID: "c1", SourceID: "s1", Quote: "Synthetic quotation"}}
			r.Research.Rounds[1].NewClaimIDs = []string{"c1"}
			r.Research.Rounds[1].MaterialGain = "minor"
		},
		"origin without evidence": func(r *curator.Report) {
			s := r.Run.Sources[0]
			s.ID = "s2"
			s.URL = "https://example.org/not-evidence"
			r.Run.Sources = append(r.Run.Sources, s)
			r.Run.Graph.Nodes = append(r.Run.Graph.Nodes, curator.Node{ID: "s2", Type: "Source", Label: "not evidence"})
			r.Research.Rounds[0].AssessedSourceIDs = append(r.Research.Rounds[0].AssessedSourceIDs, "s2")
			r.Research.Rounds[0].NewOriginIDs = []string{"s2"}
			r.Research.Rounds[0].MaterialGain = "minor"
		},
	} {
		t.Run(name, func(t *testing.T) {
			r := saturationReport(t)
			mutate(r)
			if err := curator.ValidateReport(r); err == nil {
				t.Fatal("accepted inconsistent finding")
			}
		})
	}
}

func encodeReport(t *testing.T, r *curator.Report) []byte {
	t.Helper()
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
