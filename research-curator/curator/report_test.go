package curator_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"researchcurator/curator"
)

func reportJSON(t *testing.T) []byte {
	t.Helper()
	run, err := json.Marshal(fixture())
	if err != nil {
		t.Fatal(err)
	}
	return []byte(`{"report_version":"1.0","run":` + string(run) + `,"research":{"audience":"general reader","purpose":"explain the test question","origin_types_rationale":"The source type is appropriate for this synthetic test question.","max_rounds":8,"low_gain_window":3,"rounds":[],"coverage":[{"question_id":"q1","status":"uncovered","claim_ids":[],"gaps":["research not started"]}],"gaps":["research not started"],"stop":{"reason":"in_progress","round_ids":[],"rationale":"research not started"}}}`)
}

func TestDecodeReportRequiresKnownVersionAndStrictEnvelope(t *testing.T) {
	valid := reportJSON(t)
	if _, err := curator.DecodeReport(valid); err != nil {
		t.Fatalf("valid research handoff rejected: %v", err)
	}
	for name, input := range map[string][]byte{
		"unknown field":   append(append([]byte(nil), valid[:len(valid)-1]...), []byte(`,"unexpected":true}`)...),
		"unknown version": []byte(`{"report_version":"99.0"}`),
		"missing run":     []byte(`{"report_version":"1.0","research":{},"article":null}`),
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := curator.DecodeReport(input); err == nil {
				t.Fatalf("accepted invalid report: %s", input)
			}
		})
	}
}

func TestDecodeReportRejectsDuplicateProperties(t *testing.T) {
	valid := reportJSON(t)
	duplicate := bytes.Replace(valid, []byte(`"report_version":"1.0",`), []byte(`"report_version":"1.0","report_version":"1.0",`), 1)
	if _, err := curator.DecodeReport(duplicate); err == nil {
		t.Fatal("accepted duplicate report property")
	}
}

func TestValidateReportRejectsDanglingEvidenceAndArticleReferences(t *testing.T) {
	report, err := curator.DecodeReport(reportJSON(t))
	if err != nil {
		t.Fatal(err)
	}
	report.Research.Rounds = []curator.ResearchRound{{ID: "r1", QueryIDs: []string{"missing-query"}, QuestionIDs: []string{"q1"}, SearchIntent: "discovery", SearchAngle: "overview", AssessedSourceIDs: []string{"s1"}, RedundantSourceIDs: []string{}, NewClaimIDs: []string{"c1"}, NewOriginIDs: []string{"s1"}, NewContradictionRefs: []curator.EvidenceRef{}, NewQuestionIDs: []string{}, MaterialGain: "none", MaterialityReason: "no material change", MaterialityEvidence: []curator.EvidenceRef{}, Status: "complete"}}
	if err := curator.ValidateReport(report); err == nil {
		t.Fatal("accepted round with missing query")
	}

	report.Research.Rounds = []curator.ResearchRound{}
	report.Article = &curator.Article{Title: "Synthetic article", Language: "en", OutputForm: "article", Lead: "Synthetic lead.", Sections: []curator.ArticleSection{{ID: "sec1", Heading: "Evidence", Blocks: []curator.ArticleBlock{{ID: "b1", Kind: "paragraph", Role: "evidence", Text: "A supported statement.", ClaimIDs: []string{"missing-claim"}, ConclusionIDs: []string{}}}}}}
	if err := curator.ValidateReport(report); err == nil {
		t.Fatal("accepted article with missing claim reference")
	}
	report.Article.Sections[0].Blocks[0] = curator.ArticleBlock{ID: "b1", Kind: "list", Role: "context", Text: "wrong shape", ClaimIDs: []string{}, ConclusionIDs: []string{}}
	if err := curator.ValidateReport(report); err == nil {
		t.Fatal("accepted article block with fields that do not match its kind")
	}
}

func TestValidatePublicationAcceptsEachTerminalStopReason(t *testing.T) {
	for _, reason := range []string{"user_stopped", "retrieval_blocked"} {
		t.Run(reason, func(t *testing.T) {
			report, err := curator.DecodeReport(reportJSON(t))
			if err != nil {
				t.Fatal(err)
			}
			report.Article = &curator.Article{Title: "Qualified report", Language: "en", OutputForm: "article", Lead: "A limited synthetic report.", Sections: []curator.ArticleSection{{ID: "sec", Heading: "Limits", Blocks: []curator.ArticleBlock{{ID: "b", Kind: "paragraph", Role: "context", Text: "Research was limited.", ClaimIDs: []string{}, ConclusionIDs: []string{}}}}}}
			report.Research.Stop = curator.ResearchStop{Reason: reason, RoundIDs: []string{}, Rationale: "synthetic test stop"}
			if err := curator.Finalize(&report.Run); err != nil {
				t.Fatal(err)
			}
			if err := curator.ValidatePublication(report); err != nil {
				t.Fatalf("rejected %s stop: %v", reason, err)
			}
		})
	}
	t.Run("budget_exhausted", func(t *testing.T) {
		report, err := curator.DecodeReport(reportJSON(t))
		if err != nil {
			t.Fatal(err)
		}
		report.Article = &curator.Article{Title: "Budget-limited report", Language: "en", OutputForm: "article", Lead: "A limited synthetic report.", Sections: []curator.ArticleSection{{ID: "sec", Heading: "Limits", Blocks: []curator.ArticleBlock{{ID: "b", Kind: "paragraph", Role: "context", Text: "Research stopped at its budget.", ClaimIDs: []string{}, ConclusionIDs: []string{}}}}}}
		report.Research.MaxRounds = 1
		report.Run.Queries = append(report.Run.Queries, curator.Query{ID: "budget-query", QuestionID: "q1", Text: "synthetic budget query", At: "2026-01-01T00:00:00Z"})
		report.Research.Rounds = []curator.ResearchRound{{ID: "r1", QueryIDs: []string{"budget-query"}, QuestionIDs: []string{"q1"}, SearchIntent: "discovery", SearchAngle: "synthetic test", OriginTypesSearched: []string{"docs"}, AssessedSourceIDs: []string{}, RedundantSourceIDs: []string{}, NewClaimIDs: []string{}, NewOriginIDs: []string{}, NewContradictionRefs: []curator.EvidenceRef{}, NewQuestionIDs: []string{}, MaterialGain: "none", MaterialityReason: "the failed query produced no evidence", MaterialityEvidence: []curator.EvidenceRef{}, CoverageSnapshot: []curator.ReportCoverage{}, Gaps: []string{"retrieval failed"}, Status: "failed"}}
		report.Research.Stop = curator.ResearchStop{Reason: "budget_exhausted", RoundIDs: []string{"r1"}, Rationale: "the one-round test budget was exhausted"}
		if err := curator.Finalize(&report.Run); err != nil {
			t.Fatal(err)
		}
		if err := curator.ValidatePublication(report); err != nil {
			t.Fatalf("rejected budget stop: %v", err)
		}
	})
}

func TestSourceRolesMatchRecordedEvidence(t *testing.T) {
	report, err := curator.DecodeReport(reportJSON(t))
	if err != nil {
		t.Fatal(err)
	}
	report.Research.SourceRoles = map[string]string{"s1": "candidate_lead"}
	if err := curator.ValidateReport(report); err == nil {
		t.Fatal("accepted selected claim evidence as a candidate lead")
	}
	report.Research.SourceRoles["s1"] = "evidence"
	if err := curator.ValidateReport(report); err != nil {
		t.Fatalf("rejected verified claim evidence: %v", err)
	}
}

func TestRecordResearchRoundInvalidatesOldStop(t *testing.T) {
	report, err := curator.DecodeReport(reportJSON(t))
	if err != nil {
		t.Fatal(err)
	}
	report.Research.Stop = curator.ResearchStop{Reason: "user_stopped", RoundIDs: []string{}, Rationale: "user requested stop"}
	round := curator.ResearchRound{ID: "r1", QueryIDs: []string{}, QuestionIDs: []string{}, SearchIntent: "discovery", SearchAngle: "blocked first attempt", OriginTypesSearched: []string{}, AssessedSourceIDs: []string{}, RedundantSourceIDs: []string{}, NewClaimIDs: []string{}, NewOriginIDs: []string{}, NewContradictionRefs: []curator.EvidenceRef{}, NewQuestionIDs: []string{}, MaterialGain: "none", MaterialityReason: "retrieval did not return inspectable material", MaterialityEvidence: []curator.EvidenceRef{}, CoverageSnapshot: []curator.ReportCoverage{}, Gaps: []string{"retrieval unavailable"}, Status: "failed"}
	if err := curator.RecordResearchRound(report, round); err != nil {
		t.Fatalf("record round: %v", err)
	}
	if report.Research.Stop.Reason != "in_progress" || len(report.Research.Stop.RoundIDs) != 0 {
		t.Fatalf("old stop survived new round: %+v", report.Research.Stop)
	}
	report.Research.Stop = curator.ResearchStop{Reason: "retrieval_blocked", RoundIDs: []string{}, Rationale: "stale stop"}
	if err := curator.ValidateReport(report); err == nil {
		t.Fatal("accepted terminal stop that ignored the final round")
	}
}

func TestValidatePublicationRequiresArticleAndTerminalStop(t *testing.T) {
	report, err := curator.DecodeReport(reportJSON(t))
	if err != nil {
		t.Fatal(err)
	}
	if err := curator.ValidateReport(report); err != nil {
		t.Fatalf("rejected pre-writing handoff: %v", err)
	}
	if err := curator.ValidatePublication(report); err == nil {
		t.Fatal("published without article")
	}
	report.Article = &curator.Article{Title: "Synthetic article", Language: "en", OutputForm: "article", Lead: "Synthetic lead.", Sections: []curator.ArticleSection{{ID: "sec1", Heading: "Evidence", Blocks: []curator.ArticleBlock{{ID: "b1", Kind: "paragraph", Role: "evidence", Text: "A supported statement.", ClaimIDs: []string{"c1"}, ConclusionIDs: []string{}}}}}}
	if err := curator.ValidatePublication(report); err == nil {
		t.Fatal("published while research is in progress")
	}
	report.Research.Stop = curator.ResearchStop{Reason: "user_stopped", RoundIDs: []string{}, Rationale: "synthetic test stop"}
	if err := curator.Finalize(&report.Run); err != nil {
		t.Fatal(err)
	}
	if err := curator.ValidatePublication(report); err != nil {
		t.Fatalf("rejected qualified limited report: %v", err)
	}
}
