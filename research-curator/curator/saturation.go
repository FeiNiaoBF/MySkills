package curator

import (
	"fmt"
	"strings"
)

// EvaluateSaturation assesses the recorded low-gain window. It does not infer semantic novelty.
func EvaluateSaturation(report *Report) SaturationAssessment {
	assessment := SaturationAssessment{Reasons: []string{}}
	if report == nil {
		assessment.Reasons = append(assessment.Reasons, "report is missing")
		return assessment
	}
	research := report.Research
	if b := research.ResourceBudget; b != nil && b.Used >= b.Limit {
		assessment.Reasons = append(assessment.Reasons, "declared resource budget ended; report budget_exhausted, not saturated")
	}
	if research.LowGainWindow != 3 {
		assessment.Reasons = append(assessment.Reasons, "low-gain window must be three rounds")
		return assessment
	}
	if research.MaxRounds < 1 || len(research.Rounds) > research.MaxRounds {
		assessment.Reasons = append(assessment.Reasons, "round budget is invalid or exceeded")
	}
	if len(report.Run.Contract.Questions) == 0 {
		assessment.Reasons = append(assessment.Reasons, "no coverage questions are defined")
	}
	coverage := map[string]ReportCoverage{}
	for _, item := range research.Coverage {
		coverage[item.QuestionID] = item
	}
	sources := map[string]Source{}
	for _, source := range report.Run.Sources {
		sources[source.ID] = source
	}
	claims := map[string]Claim{}
	for _, claim := range report.Run.Claims {
		claims[claim.ID] = claim
	}
	latestCoverage := map[string]ReportCoverage{}
	assessedSources := map[string]Source{}
	for _, round := range research.Rounds {
		for _, id := range round.AssessedSourceIDs {
			assessedSources[id] = sources[id]
		}
		for _, item := range round.CoverageSnapshot {
			latestCoverage[item.QuestionID] = item
		}
	}
	for _, question := range report.Run.Contract.Questions {
		item, ok := coverage[question.ID]
		if !ok || item.Status != "covered" || latestCoverage[question.ID].Status != "covered" {
			assessment.Reasons = append(assessment.Reasons, "core question is not covered: "+question.ID)
			continue
		}
		verified := false
		for _, id := range item.ClaimIDs {
			if claim, ok := claims[id]; ok && verifiedClaim(claim, assessedSources) && verifiedClaim(claim, sources) && contains(claim.QuestionIDs, question.ID) {
				verified = true
				break
			}
		}
		if !verified {
			assessment.Reasons = append(assessment.Reasons, "core question lacks a verified supported claim: "+question.ID)
		}
	}
	if len(research.Rounds) < research.LowGainWindow {
		assessment.Reasons = append(assessment.Reasons, fmt.Sprintf("need %d completed low-gain rounds", research.LowGainWindow))
		return assessment
	}
	window := research.Rounds[len(research.Rounds)-research.LowGainWindow:]
	assessment.WindowRoundIDs = make([]string, 0, len(window))
	angles, queryTexts, originTypes := map[string]bool{}, map[string]bool{}, map[string]bool{}
	counterevidence := false
	queryByID := map[string]Query{}
	for _, query := range report.Run.Queries {
		queryByID[query.ID] = query
	}
	for _, round := range window {
		assessment.WindowRoundIDs = append(assessment.WindowRoundIDs, round.ID)
		if round.Status != "complete" || len(round.AssessedSourceIDs) == 0 || len(round.QueryIDs) == 0 {
			assessment.Reasons = append(assessment.Reasons, "low-gain window contains an incomplete, failed, or empty round: "+round.ID)
			continue
		}
		if round.MaterialGain == "material" {
			assessment.Reasons = append(assessment.Reasons, "material information gain resets the window: "+round.ID)
			continue
		}
		if round.MaterialGain != "none" && round.MaterialGain != "minor" {
			assessment.Reasons = append(assessment.Reasons, "round has no valid materiality assessment: "+round.ID)
			continue
		}
		if strings.TrimSpace(round.MaterialityReason) == "" || normalizeSearchText(round.SearchAngle) == "" {
			assessment.Reasons = append(assessment.Reasons, "round has blank angle or assessment reason: "+round.ID)
			continue
		}
		assessment.LowGainRounds++
		angles[normalizeSearchText(round.SearchAngle)] = true
		if round.SearchIntent == "counterevidence" {
			counterevidence = true
		}
		for _, originType := range round.OriginTypesSearched {
			if value := normalizeSearchText(originType); value != "" {
				originTypes[value] = true
			}
		}
		for _, queryID := range round.QueryIDs {
			if query, ok := queryByID[queryID]; ok {
				if value := normalizeSearchText(query.Text); value != "" {
					queryTexts[value] = true
				}
			}
		}
	}
	if assessment.LowGainRounds != research.LowGainWindow {
		assessment.Reasons = append(assessment.Reasons, "not all rounds in the trailing window are low-gain")
	}
	if len(angles) < 2 || len(queryTexts) < 2 {
		assessment.Reasons = append(assessment.Reasons, "search formulations or angles lack sufficient variation")
	}
	if !counterevidence {
		assessment.Reasons = append(assessment.Reasons, "trailing window has no deliberate counterevidence search")
	}
	if len(originTypes) < 2 && strings.TrimSpace(research.OriginTypesRationale) == "" {
		assessment.Reasons = append(assessment.Reasons, "source-origin type coverage lacks variety or an explanation")
	}
	assessment.Eligible = len(assessment.Reasons) == 0
	return assessment
}

func normalizeSearchText(value string) string {
	return strings.Join(strings.Fields(strings.ToLower(value)), " ")
}
