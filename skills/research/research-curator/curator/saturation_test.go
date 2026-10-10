package curator_test

import (
	"encoding/json"
	"testing"

	"researchcurator/curator"
)

func saturationReport(t *testing.T) *curator.Report {
	t.Helper()
	report, err := curator.DecodeReport(reportJSON(t))
	if err != nil {
		t.Fatal(err)
	}
	run := &report.Run
	for i := 1; i <= 3; i++ {
		run.Queries = append(run.Queries, curator.Query{ID: "query-search-" + string(rune('0'+i)), QuestionID: "q1", Text: []string{"overview primary sources", "fill the evidence gap", "search for counterevidence"}[i-1], At: "2026-01-01T00:00:00Z"})
	}
	report.Research.Coverage = []curator.ReportCoverage{{QuestionID: "q1", Status: "covered", ClaimIDs: []string{"c1"}, Gaps: []string{}}}
	report.Research.Rounds = []curator.ResearchRound{
		saturationRound("r1", "query-search-1", "discovery", "map the main positions"),
		saturationRound("r2", "query-search-2", "gap_fill", "check a missing angle"),
		saturationRound("r3", "query-search-3", "counterevidence", "search for an opposing result"),
	}
	return report
}

func saturationRound(id, queryID, intent, angle string) curator.ResearchRound {
	return curator.ResearchRound{ID: id, QueryIDs: []string{queryID}, QuestionIDs: []string{"q1"}, SearchIntent: intent, SearchAngle: angle, OriginTypesSearched: []string{"primary literature"}, AssessedSourceIDs: []string{"s1"}, RedundantSourceIDs: []string{}, NewClaimIDs: []string{}, NewOriginIDs: []string{}, NewContradictionRefs: []curator.EvidenceRef{}, NewQuestionIDs: []string{}, MaterialGain: "none", MaterialityReason: "No finding changed the answer, scope, confidence, or open questions.", MaterialityEvidence: []curator.EvidenceRef{}, CoverageSnapshot: []curator.ReportCoverage{{QuestionID: "q1", Status: "covered", ClaimIDs: []string{"c1"}, Gaps: []string{}}}, Gaps: []string{}, Status: "complete"}
}

func TestDecodeReportPreservesRoundCoverageSnapshots(t *testing.T) {
	data, err := json.Marshal(saturationReport(t))
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := curator.DecodeReport(data)
	if err != nil {
		t.Fatalf("valid saturation report rejected: %v", err)
	}
	if got := len(decoded.Research.Rounds[0].CoverageSnapshot); got != 1 {
		t.Fatalf("coverage snapshot lost during decode: %d entries", got)
	}
}

func TestValidateReportRequiresCompletedRoundCoverageSnapshots(t *testing.T) {
	report := saturationReport(t)
	if err := curator.ValidateReport(report); err != nil {
		t.Fatalf("valid report rejected: %v", err)
	}
	report.Research.Rounds[0].CoverageSnapshot = nil
	if err := curator.ValidateReport(report); err == nil {
		t.Fatal("accepted completed round without coverage snapshot")
	}
}

func TestEvaluateSaturationAcceptsThreeCoveredLowGainRounds(t *testing.T) {
	assessment := curator.EvaluateSaturation(saturationReport(t))
	if !assessment.Eligible || assessment.LowGainRounds != 3 {
		t.Fatalf("unexpected assessment: %+v", assessment)
	}
}

func TestEvaluateSaturationAcceptsDocumentedStableDisagreement(t *testing.T) {
	report := saturationReport(t)
	source := report.Run.Sources[0]
	source.ID, source.URL, source.Title, source.Content = "s2", "https://example.org/alternative", "Synthetic alternative source", "An alternative documented position."
	report.Run.Sources = append(report.Run.Sources, source)
	claim := report.Run.Claims[0]
	claim.ID, claim.Text = "c2", "A distinct synthetic claim represents the stable alternative position."
	claim.Evidence = []curator.Evidence{{SourceID: "s2", Quote: "An alternative documented position", Locator: "test alternative paragraph", Relation: "supports", Verification: "verified"}}
	report.Run.Claims = append(report.Run.Claims, claim)
	report.Run.Graph.Nodes = append(report.Run.Graph.Nodes, curator.Node{ID: "s2", Type: "Source", Label: "Alternative source"}, curator.Node{ID: "c2", Type: "Claim", Label: "Alternative claim"})
	report.Run.Graph.Edges = append(report.Run.Graph.Edges, curator.Edge{From: "q1", To: "c2", Type: "addresses"}, curator.Edge{From: "s2", To: "c2", Type: "supports"})
	report.Research.Coverage[0].ClaimIDs = []string{"c1", "c2"}
	report.Run.Queries = append(report.Run.Queries, curator.Query{ID: "query-stable-disagreement", QuestionID: "q1", Text: "record both documented positions", At: "2026-01-01T00:00:00Z"})
	initial := saturationRound("r0", "query-stable-disagreement", "discovery", "record the already assessed alternative")
	initial.MaterialGain = "material"
	initial.MaterialityReason = "The research records an established opposing position; the final coverage preserves both claims."
	initial.NewClaimIDs = []string{"c2"}
	initial.NewOriginIDs = []string{"s2"}
	initial.AssessedSourceIDs = []string{"s2"}
	initial.RedundantSourceIDs = []string{}
	initial.MaterialityEvidence = []curator.EvidenceRef{{ClaimID: "c2", SourceID: "s2", Quote: "An alternative documented position"}}
	initial.CoverageSnapshot[0].ClaimIDs = []string{"c1", "c2"}
	for i := range report.Research.Rounds {
		report.Research.Rounds[i].ID = []string{"r1", "r2", "r3"}[i]
		report.Research.Rounds[i].CoverageSnapshot[0].ClaimIDs = []string{"c1", "c2"}
	}
	report.Research.Rounds = append([]curator.ResearchRound{initial}, report.Research.Rounds...)
	if err := curator.ValidateReport(report); err != nil {
		t.Fatal(err)
	}
	assessment := curator.EvaluateSaturation(report)
	if !assessment.Eligible || assessment.LowGainRounds != 3 {
		t.Fatalf("stable, fully represented disagreement blocked saturation: %+v", assessment)
	}
}

func TestEvaluateSaturationResetsOnMaterialCounterexampleOrEvidenceUpgrade(t *testing.T) {
	report := saturationReport(t)
	last := &report.Research.Rounds[2]
	last.MaterialGain = "material"
	last.MaterialityReason = "A new result changes the conclusion's scope."
	last.MaterialityEvidence = []curator.EvidenceRef{{ClaimID: "c1", SourceID: "s1", Quote: "Synthetic quotation"}}
	counter := report.Run.Sources[0]
	counter.ID = "s2"
	counter.URL = "https://example.org/counter"
	counter.Title = "Synthetic counter-source"
	counter.Content = "Counter quotation for testing."
	report.Run.Sources = append(report.Run.Sources, counter)
	report.Run.Claims[0].Evidence = append(report.Run.Claims[0].Evidence, curator.Evidence{SourceID: "s2", Quote: "Counter quotation", Locator: "test paragraph 2", Relation: "contradicts", Verification: "verified"})
	report.Run.Graph.Nodes = append(report.Run.Graph.Nodes, curator.Node{ID: "s2", Type: "Source", Label: "Counter source"})
	report.Run.Graph.Edges = append(report.Run.Graph.Edges, curator.Edge{From: "s2", To: "c1", Type: "contradicts"})
	last.NewContradictionRefs = []curator.EvidenceRef{{ClaimID: "c1", SourceID: "s2", Quote: "Counter quotation"}}
	last.AssessedSourceIDs = append(last.AssessedSourceIDs, "s2")
	last.CoverageSnapshot[0].Status = "partial"
	last.CoverageSnapshot[0].Gaps = []string{"new counterexample needs assessment"}
	report.Research.Coverage[0].Status = "partial"
	report.Research.Coverage[0].Gaps = []string{"new counterexample needs assessment"}
	if err := curator.ValidateReport(report); err != nil {
		t.Fatal(err)
	}
	assessment := curator.EvaluateSaturation(report)
	if assessment.Eligible {
		t.Fatalf("material counterevidence did not reset saturation: %+v", assessment)
	}
}

func TestEvaluateSaturationResetsOnMaterialEvidenceUpgrade(t *testing.T) {
	report := saturationReport(t)
	last := &report.Research.Rounds[2]
	last.MaterialGain = "material"
	last.MaterialityReason = "A newly assessed passage materially increases confidence in an existing claim."
	last.MaterialityEvidence = []curator.EvidenceRef{{ClaimID: "c1", SourceID: "s1", Quote: "Synthetic quotation"}}
	if assessment := curator.EvaluateSaturation(report); assessment.Eligible {
		t.Fatalf("material evidence upgrade did not reset saturation: %+v", assessment)
	}
}

func TestEvaluateSaturationExcludesIncompleteEmptyAndFailedRounds(t *testing.T) {
	for name, mutate := range map[string]func(*curator.Report){
		"partial":                func(r *curator.Report) { r.Research.Rounds[2].Status = "partial" },
		"failed":                 func(r *curator.Report) { r.Research.Rounds[2].Status = "failed" },
		"empty assessed sources": func(r *curator.Report) { r.Research.Rounds[2].AssessedSourceIDs = []string{} },
	} {
		t.Run(name, func(t *testing.T) {
			report := saturationReport(t)
			mutate(report)
			if assessment := curator.EvaluateSaturation(report); assessment.Eligible {
				t.Fatalf("ineligible round counted: %+v", assessment)
			}
		})
	}
}

func TestValidateSaturatedStopMustReferenceExactWindow(t *testing.T) {
	report := saturationReport(t)
	report.Research.Stop = curator.ResearchStop{Reason: "saturated", RoundIDs: []string{"r1", "r2", "r3"}, Rationale: "Three low-gain rounds followed coverage and counterevidence checks."}
	if err := curator.Finalize(&report.Run); err != nil {
		t.Fatal(err)
	}
	report.Article = &curator.Article{Title: "Synthetic report", Language: "en", OutputForm: "article", Lead: "A supported synthetic finding.", Sections: []curator.ArticleSection{{ID: "sec1", Heading: "Finding", Blocks: []curator.ArticleBlock{{ID: "b1", Kind: "paragraph", Role: "evidence", Text: "The source supports this synthetic finding.", ClaimIDs: []string{"c1"}, ConclusionIDs: []string{}}}}}}
	if err := curator.ValidatePublication(report); err != nil {
		t.Fatalf("rejected valid saturated publication: %v", err)
	}
	report.Research.Stop.RoundIDs = []string{"r1", "r2", "r2"}
	if err := curator.ValidatePublication(report); err == nil {
		t.Fatal("accepted stop that cites the wrong round window")
	}
}

func TestEvaluateSaturationRequiresCounterevidenceSearch(t *testing.T) {
	report := saturationReport(t)
	report.Research.Rounds[2].SearchIntent = "gap_fill"
	if assessment := curator.EvaluateSaturation(report); assessment.Eligible {
		t.Fatalf("saturated without counterevidence search: %+v", assessment)
	}
}
