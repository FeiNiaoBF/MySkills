package curator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"strings"
	"unicode/utf8"

	"researchcurator/schema"
)

// Report is the versioned handoff from research to writing and local publication.
type Report struct {
	ReportVersion string         `json:"report_version"`
	Run           Run            `json:"run"`
	Research      ResearchRecord `json:"research"`
	Article       *Article       `json:"article,omitempty"`
}

type ResearchRecord struct {
	Audience             string            `json:"audience"`
	Purpose              string            `json:"purpose"`
	OriginTypesRationale string            `json:"origin_types_rationale"`
	MaxRounds            int               `json:"max_rounds"`
	ResourceBudget       *ResourceBudget   `json:"resource_budget,omitempty"`
	SourceRoles          map[string]string `json:"source_roles,omitempty"`
	LowGainWindow        int               `json:"low_gain_window"`
	Rounds               []ResearchRound   `json:"rounds"`
	Coverage             []ReportCoverage  `json:"coverage"`
	Gaps                 []string          `json:"gaps"`
	Stop                 ResearchStop      `json:"stop"`
}

// ResourceBudget records a limit declared before retrieval and actual usage.
// Units are explicit (for example seconds, calls, or USD); no conversion is inferred.
type ResourceBudget struct {
	Kind  string  `json:"kind"`
	Unit  string  `json:"unit"`
	Limit float64 `json:"limit"`
	Used  float64 `json:"used"`
}

type ResearchRound struct {
	ID                   string            `json:"id"`
	QueryIDs             []string          `json:"query_ids"`
	QuestionIDs          []string          `json:"question_ids"`
	SearchIntent         string            `json:"search_intent"`
	SearchAngle          string            `json:"search_angle"`
	OriginTypesSearched  []string          `json:"origin_types_searched"`
	AssessedSourceIDs    []string          `json:"assessed_source_ids"`
	RedundantSourceIDs   []string          `json:"redundant_source_ids"`
	RedundancyReasons    map[string]string `json:"redundancy_reasons,omitempty"`
	NewClaimIDs          []string          `json:"new_claim_ids"`
	NewOriginIDs         []string          `json:"new_origin_ids"`
	NewContradictionRefs []EvidenceRef     `json:"new_contradiction_refs"`
	NewQuestionIDs       []string          `json:"new_question_ids"`
	MaterialGain         string            `json:"material_gain"`
	MaterialityReason    string            `json:"materiality_reason"`
	MaterialityEvidence  []EvidenceRef     `json:"materiality_evidence"`
	CoverageSnapshot     []ReportCoverage  `json:"coverage_snapshot"`
	Gaps                 []string          `json:"gaps"`
	Status               string            `json:"status"`
}

// EvidenceRef points into the canonical v1 Claim.Evidence record.
type EvidenceRef struct {
	ClaimID  string `json:"claim_id"`
	SourceID string `json:"source_id"`
	Quote    string `json:"quote"`
}

type ReportCoverage struct {
	QuestionID string   `json:"question_id"`
	Status     string   `json:"status"`
	ClaimIDs   []string `json:"claim_ids"`
	Gaps       []string `json:"gaps"`
}

type ResearchStop struct {
	Reason    string   `json:"reason"`
	RoundIDs  []string `json:"round_ids"`
	Rationale string   `json:"rationale"`
}

type SaturationAssessment struct {
	Eligible       bool     `json:"eligible"`
	LowGainRounds  int      `json:"low_gain_rounds"`
	WindowRoundIDs []string `json:"window_round_ids"`
	Reasons        []string `json:"reasons"`
}

type Article struct {
	Title      string           `json:"title"`
	Language   string           `json:"language"`
	OutputForm string           `json:"output_form"`
	Lead       string           `json:"lead"`
	Sections   []ArticleSection `json:"sections"`
}

type ArticleSection struct {
	ID      string         `json:"id"`
	Heading string         `json:"heading"`
	Blocks  []ArticleBlock `json:"blocks"`
}

type ArticleBlock struct {
	ID            string     `json:"id"`
	Kind          string     `json:"kind"`
	Role          string     `json:"role"`
	Text          string     `json:"text,omitempty"`
	Items         []string   `json:"items,omitempty"`
	Headers       []string   `json:"headers,omitempty"`
	Rows          [][]string `json:"rows,omitempty"`
	ClaimIDs      []string   `json:"claim_ids"`
	ConclusionIDs []string   `json:"conclusion_ids"`
}

// DecodeReport validates the report schema, the embedded v1 run, and cross-references.
func DecodeReport(data []byte) (*Report, error) {
	if !utf8.Valid(data) || !json.Valid(data) {
		return nil, fmt.Errorf("invalid report JSON")
	}
	scan := json.NewDecoder(bytes.NewReader(data))
	if err := scanValue(scan, 0); err != nil {
		return nil, err
	}
	var raw any
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(&raw); err != nil {
		return nil, err
	}
	if err := d.Decode(new(any)); err != io.EOF {
		return nil, fmt.Errorf("trailing report JSON")
	}
	var schemaDoc map[string]any
	if err := json.Unmarshal(schema.Report, &schemaDoc); err != nil {
		return nil, err
	}
	if err := checkSchema(raw, schemaDoc, "$"); err != nil {
		return nil, err
	}
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(data, &envelope); err != nil {
		return nil, err
	}
	if _, err := Decode(envelope["run"]); err != nil {
		return nil, fmt.Errorf("invalid embedded run: %w", err)
	}
	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, err
	}
	if err := ValidateReport(&report); err != nil {
		return nil, err
	}
	return &report, nil
}

// ValidateReport accepts an in-progress research handoff without an article.
func ValidateReport(report *Report) error {
	if report == nil {
		return fmt.Errorf("nil report")
	}
	if report.ReportVersion != "1.0" {
		return fmt.Errorf("unsupported report version %q", report.ReportVersion)
	}
	if err := Validate(&report.Run); err != nil {
		return fmt.Errorf("invalid embedded run: %w", err)
	}
	r := &report.Research
	if strings.TrimSpace(r.Audience) == "" || strings.TrimSpace(r.Purpose) == "" {
		return fmt.Errorf("research audience and purpose are required")
	}
	if r.MaxRounds < 1 || r.LowGainWindow != 3 || len(r.Rounds) > r.MaxRounds {
		return fmt.Errorf("invalid research round budget or low-gain window")
	}
	if b := r.ResourceBudget; b != nil {
		if (b.Kind != "time" && b.Kind != "tool_calls" && b.Kind != "cost") || strings.TrimSpace(b.Unit) == "" || b.Limit <= 0 || b.Used < 0 || math.IsNaN(b.Limit) || math.IsNaN(b.Used) || math.IsInf(b.Limit, 0) || math.IsInf(b.Used, 0) {
			return fmt.Errorf("invalid resource budget")
		}
	}
	if r.Rounds == nil || r.Coverage == nil || r.Gaps == nil || r.Stop.RoundIDs == nil {
		return fmt.Errorf("research arrays must be non-null")
	}
	if err := validateNonblank(r.Gaps, "research gap"); err != nil {
		return err
	}
	if err := validateResearchReferences(report); err != nil {
		return err
	}
	if r.SourceRoles != nil {
		if err := validateSourceRoles(report); err != nil {
			return err
		}
	}
	if report.Article != nil {
		if err := validateArticle(report); err != nil {
			return err
		}
	}
	return validateStop(report)
}

// ValidatePublication additionally requires a complete article and a terminal stop reason.
func ValidatePublication(report *Report) error {
	if err := ValidateReport(report); err != nil {
		return err
	}
	if report.Article == nil {
		return fmt.Errorf("article is required for publication")
	}
	if report.Run.Metadata.Status != "finalized" {
		return fmt.Errorf("publication requires a finalized evidence ledger")
	}
	if report.Research.Stop.Reason == "in_progress" {
		return fmt.Errorf("in-progress research cannot be published as final")
	}
	return nil
}

func validateResearchReferences(report *Report) error {
	run := &report.Run
	questions, queries, sources, claims := map[string]bool{}, map[string]Query{}, map[string]bool{}, map[string]Claim{}
	for _, q := range run.Contract.Questions {
		questions[q.ID] = true
	}
	for _, q := range run.Queries {
		queries[q.ID] = q
	}
	for _, s := range run.Sources {
		sources[s.ID] = true
	}
	for _, c := range run.Claims {
		claims[c.ID] = c
	}
	verifiedSources := sourceMap(run.Sources)
	rounds := map[string]ResearchRound{}
	seenQueries, seenClaims, seenOrigins, seenQuestions := map[string]bool{}, map[string]bool{}, map[string]bool{}, map[string]bool{}
	seenContradictions := map[EvidenceRef]bool{}
	assessedSources := map[string]Source{}
	for _, round := range report.Research.Rounds {
		if round.ID == "" || rounds[round.ID].ID != "" {
			return fmt.Errorf("empty or duplicate research round ID %q", round.ID)
		}
		rounds[round.ID] = round
		if round.SearchIntent != "discovery" && round.SearchIntent != "gap_fill" && round.SearchIntent != "counterevidence" {
			return fmt.Errorf("invalid search intent in round %s", round.ID)
		}
		if strings.TrimSpace(round.SearchAngle) == "" || strings.TrimSpace(round.MaterialityReason) == "" {
			return fmt.Errorf("round %s requires search angle and materiality reason", round.ID)
		}
		if round.Status == "complete" && (len(round.QueryIDs) == 0 || len(round.QuestionIDs) == 0 || len(round.OriginTypesSearched) == 0) {
			return fmt.Errorf("completed round %s must record a query", round.ID)
		}
		if err := validateUniqueStrings(round.OriginTypesSearched, "round origin type"); err != nil {
			return err
		}
		if round.MaterialGain != "none" && round.MaterialGain != "minor" && round.MaterialGain != "material" {
			return fmt.Errorf("invalid material gain in round %s", round.ID)
		}
		if round.Status != "complete" && round.Status != "partial" && round.Status != "failed" {
			return fmt.Errorf("invalid status in round %s", round.ID)
		}
		if round.QueryIDs == nil || round.QuestionIDs == nil || round.OriginTypesSearched == nil || round.AssessedSourceIDs == nil || round.RedundantSourceIDs == nil || round.NewClaimIDs == nil || round.NewOriginIDs == nil || round.NewContradictionRefs == nil || round.NewQuestionIDs == nil || round.MaterialityEvidence == nil || round.CoverageSnapshot == nil || round.Gaps == nil {
			return fmt.Errorf("round %s arrays must be non-null", round.ID)
		}
		if err := validateNonblank(round.Gaps, "round gap"); err != nil {
			return err
		}
		for label, ids := range map[string][]string{"query": round.QueryIDs, "question": round.QuestionIDs, "assessed source": round.AssessedSourceIDs, "redundant source": round.RedundantSourceIDs, "new claim": round.NewClaimIDs, "new origin": round.NewOriginIDs, "new question": round.NewQuestionIDs} {
			if err := validateUniqueStrings(ids, "round "+label); err != nil {
				return err
			}
		}
		if round.MaterialGain == "none" && len(round.NewClaimIDs)+len(round.NewOriginIDs)+len(round.NewContradictionRefs)+len(round.NewQuestionIDs) != 0 {
			return fmt.Errorf("round %s records new findings but no information gain", round.ID)
		}
		for _, id := range round.QueryIDs {
			query, ok := queries[id]
			if !ok || !contains(round.QuestionIDs, query.QuestionID) || seenQueries[id] {
				return fmt.Errorf("round %s references missing, unrelated, or reused query %s", round.ID, id)
			}
			seenQueries[id] = true
		}
		for _, id := range round.QuestionIDs {
			if !questions[id] {
				return fmt.Errorf("round %s references missing question %s", round.ID, id)
			}
		}
		for _, id := range round.AssessedSourceIDs {
			if !sources[id] {
				return fmt.Errorf("round %s references missing assessed source %s", round.ID, id)
			}
		}
		assessed := stringSet(round.AssessedSourceIDs)
		for _, id := range round.AssessedSourceIDs {
			assessedSources[id] = verifiedSources[id]
		}
		for id, reason := range round.RedundancyReasons {
			if !contains(round.RedundantSourceIDs, id) || strings.TrimSpace(reason) == "" {
				return fmt.Errorf("round %s has invalid redundancy reason for %s", round.ID, id)
			}
		}
		for _, id := range round.RedundantSourceIDs {
			if !assessed[id] {
				return fmt.Errorf("round %s redundant source %s was not assessed", round.ID, id)
			}
			source := verifiedSources[id]
			if strings.TrimSpace(round.RedundancyReasons[id]) == "" && (strings.TrimSpace(source.Reason) == "" || (source.DuplicateOf == "" && len(source.UpstreamIDs) == 0)) {
				return fmt.Errorf("round %s redundant source %s needs a round reason or explicit ledger duplication provenance", round.ID, id)
			}
		}
		for _, id := range round.NewClaimIDs {
			claim, ok := claims[id]
			if !ok || seenClaims[id] || !claimAddressesRound(claim, round) {
				return fmt.Errorf("round %s references missing or previously known new claim %s", round.ID, id)
			}
			found := false
			for _, evidence := range claim.Evidence {
				if assessed[evidence.SourceID] {
					found = true
				}
			}
			if !found {
				return fmt.Errorf("round %s new claim %s has no assessed evidence", round.ID, id)
			}
			seenClaims[id] = true
		}
		for _, id := range round.NewOriginIDs {
			providesEvidence := false
			for _, claim := range claims {
				for _, evidence := range claim.Evidence {
					if evidence.SourceID == id {
						providesEvidence = true
					}
				}
			}
			if !providesEvidence {
				return fmt.Errorf("round %s new origin %s provides no recorded evidence", round.ID, id)
			}
			source, ok := sourceByID(run.Sources, id)
			if !ok || !assessed[id] || seenOrigins[id] || !source.Original || len(source.UpstreamIDs) > 0 || source.DuplicateOf != "" {
				return fmt.Errorf("round %s references a non-original evidence origin %s", round.ID, id)
			}
		}
		for _, id := range round.NewQuestionIDs {
			if !questions[id] || !contains(round.QuestionIDs, id) || seenQuestions[id] {
				return fmt.Errorf("round %s references missing new question %s", round.ID, id)
			}
		}
		for _, ref := range round.NewContradictionRefs {
			if seenContradictions[ref] || !claimAddressesRound(claims[ref.ClaimID], round) || !assessed[ref.SourceID] || !hasEvidenceRelation(claims[ref.ClaimID], ref.SourceID, ref.Quote, "contradicts") {
				return fmt.Errorf("round %s references missing, repeated, or unassessed contradictory evidence", round.ID)
			}
			seenContradictions[ref] = true
		}
		for _, ref := range round.MaterialityEvidence {
			if !claimAddressesRound(claims[ref.ClaimID], round) || !assessed[ref.SourceID] || !hasEvidence(claims[ref.ClaimID], ref.SourceID, ref.Quote) {
				return fmt.Errorf("round %s references missing materiality evidence", round.ID)
			}
		}
		if round.MaterialGain == "material" && len(round.MaterialityEvidence) == 0 {
			return fmt.Errorf("material round %s requires evidence references", round.ID)
		}
		if err := validateRoundCoverage(round, claims, assessedSources, questions); err != nil {
			return err
		}
		for _, ref := range append(append([]EvidenceRef{}, round.MaterialityEvidence...), round.NewContradictionRefs...) {
			seenClaims[ref.ClaimID] = true
			if hasEvidenceRelation(claims[ref.ClaimID], ref.SourceID, ref.Quote, "contradicts") {
				seenContradictions[ref] = true
			}
		}
		for _, id := range round.AssessedSourceIDs {
			seenOrigins[id] = true
		}
		for _, id := range round.QuestionIDs {
			seenQuestions[id] = true
		}
		for _, item := range round.CoverageSnapshot {
			for _, id := range item.ClaimIDs {
				seenClaims[id] = true
			}
		}
	}
	coverage := map[string]bool{}
	for _, item := range report.Research.Coverage {
		if !questions[item.QuestionID] || coverage[item.QuestionID] {
			return fmt.Errorf("invalid or duplicate coverage question %s", item.QuestionID)
		}
		coverage[item.QuestionID] = true
		if err := validateNonblank(item.Gaps, "coverage gap"); err != nil {
			return err
		}
		if err := validateUniqueStrings(item.ClaimIDs, "coverage claim"); err != nil {
			return err
		}
		if item.Status != "covered" && item.Status != "partial" && item.Status != "uncovered" {
			return fmt.Errorf("invalid coverage status for %s", item.QuestionID)
		}
		if item.ClaimIDs == nil || item.Gaps == nil {
			return fmt.Errorf("coverage arrays for %s must be non-null", item.QuestionID)
		}
		for _, id := range item.ClaimIDs {
			claim, ok := claims[id]
			if !ok || !contains(claim.QuestionIDs, item.QuestionID) {
				return fmt.Errorf("coverage for %s references unrelated claim %s", item.QuestionID, id)
			}
		}
		if item.Status == "covered" && (len(item.ClaimIDs) == 0 || len(item.Gaps) != 0) {
			return fmt.Errorf("covered question %s must have claims and no gaps", item.QuestionID)
		}
		if item.Status == "covered" {
			verified := false
			for _, id := range item.ClaimIDs {
				if verifiedClaim(claims[id], verifiedSources) {
					verified = true
				}
			}
			if !verified {
				return fmt.Errorf("covered question %s lacks verified supporting evidence", item.QuestionID)
			}
		}
		if item.Status != "covered" && len(item.Gaps) == 0 {
			return fmt.Errorf("incomplete question %s requires a gap", item.QuestionID)
		}
	}
	if len(coverage) != len(questions) {
		return fmt.Errorf("research coverage must include every contract question")
	}
	stopRounds := stringSet(report.Research.Stop.RoundIDs)
	if len(stopRounds) != len(report.Research.Stop.RoundIDs) {
		return fmt.Errorf("stop decision repeats a round ID")
	}
	for id := range stopRounds {
		if rounds[id].ID == "" {
			return fmt.Errorf("stop decision references missing round %s", id)
		}
	}
	return nil
}

func validateRoundCoverage(round ResearchRound, claims map[string]Claim, sources map[string]Source, questions map[string]bool) error {
	inRound := stringSet(round.QuestionIDs)
	seen := map[string]bool{}
	for _, item := range round.CoverageSnapshot {
		if !questions[item.QuestionID] || !inRound[item.QuestionID] || seen[item.QuestionID] {
			return fmt.Errorf("round %s has invalid coverage question %s", round.ID, item.QuestionID)
		}
		seen[item.QuestionID] = true
		if err := validateNonblank(item.Gaps, "snapshot gap"); err != nil {
			return err
		}
		if err := validateUniqueStrings(item.ClaimIDs, "snapshot claim"); err != nil {
			return err
		}
		if item.Status != "covered" && item.Status != "partial" && item.Status != "uncovered" {
			return fmt.Errorf("round %s has invalid coverage status for %s", round.ID, item.QuestionID)
		}
		if item.ClaimIDs == nil || item.Gaps == nil {
			return fmt.Errorf("round %s coverage arrays must be non-null", round.ID)
		}
		verified := false
		for _, id := range item.ClaimIDs {
			claim, ok := claims[id]
			if !ok || !contains(claim.QuestionIDs, item.QuestionID) {
				return fmt.Errorf("round %s coverage references unrelated claim %s", round.ID, id)
			}
			// A historical snapshot is not today's coverage verdict. Later
			// contradictions or source replacement must not rewrite its history.
			for _, evidence := range claim.Evidence {
				if evidence.Relation == "supports" && evidence.Verification == "verified" && sources[evidence.SourceID].Verification == "verified" {
					verified = true
				}
			}
		}
		if item.Status == "covered" && (!verified || len(item.Gaps) != 0) {
			return fmt.Errorf("round %s marks unsupported or gapped question %s covered", round.ID, item.QuestionID)
		}
		if item.Status != "covered" && len(item.Gaps) == 0 {
			return fmt.Errorf("round %s incomplete question %s requires a gap", round.ID, item.QuestionID)
		}
	}
	if round.Status == "complete" && len(seen) != len(inRound) {
		return fmt.Errorf("completed round %s must record coverage for every addressed question", round.ID)
	}
	return nil
}

func validateStop(report *Report) error {
	stop := report.Research.Stop
	if b := report.Research.ResourceBudget; b != nil && b.Used >= b.Limit && stop.Reason != "budget_exhausted" {
		return fmt.Errorf("declared resource budget ended; stop reason must be budget_exhausted")
	}
	if strings.TrimSpace(stop.Rationale) == "" {
		return fmt.Errorf("stop rationale is required")
	}
	switch stop.Reason {
	case "in_progress":
		if len(stop.RoundIDs) != 0 {
			return fmt.Errorf("in-progress run cannot cite terminal stop rounds")
		}
	case "saturated":
		assessment := EvaluateSaturation(report)
		if !assessment.Eligible {
			return fmt.Errorf("saturation not established: %s", strings.Join(assessment.Reasons, "; "))
		}
		if !sameStrings(stop.RoundIDs, assessment.WindowRoundIDs) {
			return fmt.Errorf("saturated stop must cite the assessed low-gain window")
		}
	case "budget_exhausted":
		b := report.Research.ResourceBudget
		if len(report.Research.Rounds) < report.Research.MaxRounds && (b == nil || b.Used < b.Limit) {
			return fmt.Errorf("budget stop before a declared budget was exhausted")
		}
		if len(report.Research.Rounds) == 0 {
			if len(stop.RoundIDs) != 0 {
				return fmt.Errorf("budget stop before first round cannot cite rounds")
			}
			return nil
		}
		last := report.Research.Rounds[len(report.Research.Rounds)-1].ID
		if len(stop.RoundIDs) != 1 || stop.RoundIDs[0] != last {
			return fmt.Errorf("budget stop must cite the final attempted round")
		}
	case "retrieval_blocked", "user_stopped":
		if len(report.Research.Rounds) > 0 {
			last := report.Research.Rounds[len(report.Research.Rounds)-1]
			if !contains(stop.RoundIDs, last.ID) {
				return fmt.Errorf("terminal stop must reference the final research round")
			}
			if stop.Reason == "retrieval_blocked" && last.Status == "complete" {
				return fmt.Errorf("retrieval_blocked stop must reference a partial or failed final round")
			}
		}
	default:
		return fmt.Errorf("invalid research stop reason %q", stop.Reason)
	}
	return nil
}

func validateArticle(report *Report) error {
	a := report.Article
	if strings.TrimSpace(a.Title) == "" || strings.TrimSpace(a.Language) == "" || strings.TrimSpace(a.OutputForm) == "" || strings.TrimSpace(a.Lead) == "" || a.Sections == nil || len(a.Sections) == 0 {
		return fmt.Errorf("article requires title, language, output form, lead, and sections")
	}
	claims, conclusions := map[string]bool{}, map[string]bool{}
	for _, c := range report.Run.Claims {
		claims[c.ID] = true
	}
	for _, c := range report.Run.Conclusions {
		conclusions[c.ID] = true
	}
	sections, blocks := map[string]bool{}, map[string]bool{}
	for _, section := range a.Sections {
		if section.ID == "" || sections[section.ID] || strings.TrimSpace(section.Heading) == "" || len(section.Blocks) == 0 {
			return fmt.Errorf("invalid article section %q", section.ID)
		}
		sections[section.ID] = true
		for _, block := range section.Blocks {
			if block.ID == "" || blocks[block.ID] {
				return fmt.Errorf("empty or duplicate article block ID %q", block.ID)
			}
			blocks[block.ID] = true
			if block.Role != "evidence" && block.Role != "synthesis" && block.Role != "context" {
				return fmt.Errorf("invalid article block role %s", block.Role)
			}
			switch block.Kind {
			case "paragraph":
				if strings.TrimSpace(block.Text) == "" || block.Items != nil || block.Headers != nil || block.Rows != nil {
					return fmt.Errorf("invalid paragraph block %s", block.ID)
				}
			case "list":
				if err := validateNonblank(block.Items, "list item"); err != nil {
					return err
				}
				if len(block.Items) == 0 || block.Text != "" || block.Headers != nil || block.Rows != nil {
					return fmt.Errorf("invalid list block %s", block.ID)
				}
			case "table":
				if err := validateNonblank(block.Headers, "table header"); err != nil {
					return err
				}
				if len(block.Headers) == 0 || len(block.Rows) == 0 || block.Text != "" || block.Items != nil {
					return fmt.Errorf("invalid table block %s", block.ID)
				}
				for _, row := range block.Rows {
					if strings.TrimSpace(strings.Join(row, "")) == "" {
						return fmt.Errorf("empty table row in block %s", block.ID)
					}
					if len(row) != len(block.Headers) {
						return fmt.Errorf("table row width mismatch in block %s", block.ID)
					}
				}
			default:
				return fmt.Errorf("invalid article block kind %s", block.Kind)
			}
			if block.ClaimIDs == nil || block.ConclusionIDs == nil {
				return fmt.Errorf("article block references must be non-null")
			}
			if block.Role == "evidence" && len(block.ClaimIDs) == 0 {
				return fmt.Errorf("evidence block %s requires claim references", block.ID)
			}
			if block.Role == "synthesis" && len(block.ClaimIDs)+len(block.ConclusionIDs) == 0 {
				return fmt.Errorf("synthesis block %s requires evidence references", block.ID)
			}
			for _, id := range block.ClaimIDs {
				if !claims[id] {
					return fmt.Errorf("article block %s references missing claim %s", block.ID, id)
				}
			}
			for _, id := range block.ConclusionIDs {
				if !conclusions[id] {
					return fmt.Errorf("article block %s references missing conclusion %s", block.ID, id)
				}
			}
		}
	}
	return nil
}

func claimAddressesRound(claim Claim, round ResearchRound) bool {
	for _, id := range claim.QuestionIDs {
		if contains(round.QuestionIDs, id) {
			return true
		}
	}
	return false
}

func validateNonblank(values []string, label string) error {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("blank %s", label)
		}
	}
	return nil
}

func hasEvidence(claim Claim, sourceID, quote string) bool {
	return hasEvidenceRelation(claim, sourceID, quote, "")
}

func hasEvidenceRelation(claim Claim, sourceID, quote, relation string) bool {
	if sourceID == "" || quote == "" {
		return false
	}
	for _, evidence := range claim.Evidence {
		if evidence.SourceID == sourceID && evidence.Quote == quote && (relation == "" || evidence.Relation == relation) {
			return true
		}
	}
	return false
}

func validateUniqueStrings(values []string, label string) error {
	seen := map[string]bool{}
	for _, value := range values {
		if strings.TrimSpace(value) == "" || seen[value] {
			return fmt.Errorf("empty or duplicate %s %q", label, value)
		}
		seen[value] = true
	}
	return nil
}

func sourceByID(sources []Source, id string) (Source, bool) {
	for _, source := range sources {
		if source.ID == id {
			return source, true
		}
	}
	return Source{}, false
}

func sourceMap(sources []Source) map[string]Source {
	out := make(map[string]Source, len(sources))
	for _, source := range sources {
		out[source.ID] = source
	}
	return out
}

func sameStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func stringSet(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, value := range values {
		out[value] = true
	}
	return out
}
