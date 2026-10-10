package curator

import (
	"fmt"
	"time"
)

// RecordResearchRound appends a validated round and invalidates any prior stop decision.
// It does not mutate report when the new handoff is invalid.
func RecordResearchRound(report *Report, round ResearchRound) error {
	if report == nil {
		return fmt.Errorf("nil report")
	}
	if report.Article != nil {
		return fmt.Errorf("cannot append research after writing has started")
	}
	trial := *report
	trial.Research = report.Research
	trial.Research.Rounds = append(append([]ResearchRound{}, report.Research.Rounds...), round)
	trial.Research.Stop = ResearchStop{
		Reason:    "in_progress",
		RoundIDs:  []string{},
		Rationale: "A new research round invalidated the previous stop decision.",
	}
	trial.Run.Metadata.Status = "in_progress"
	trial.Run.Metadata.CompletedAt = ""
	if err := ValidateReport(&trial); err != nil {
		return err
	}
	*report = trial
	return nil
}

// validateSourceRoles keeps the reader-facing evidence categories tied to ledger facts.
func validateSourceRoles(report *Report) error {
	roles := report.Research.SourceRoles
	if len(roles) != len(report.Run.Sources) {
		return fmt.Errorf("source_roles must classify every source exactly once")
	}
	verifiedEvidence := map[string]bool{}
	anyEvidence := map[string]bool{}
	for _, claim := range report.Run.Claims {
		for _, evidence := range claim.Evidence {
			anyEvidence[evidence.SourceID] = true
			if evidence.Verification == "verified" {
				if source, ok := sourceByID(report.Run.Sources, evidence.SourceID); ok && source.Verification == "verified" {
					verifiedEvidence[evidence.SourceID] = true
				}
			}
		}
	}
	for _, source := range report.Run.Sources {
		role, ok := roles[source.ID]
		if !ok {
			return fmt.Errorf("source_roles omits %s", source.ID)
		}
		switch role {
		case "evidence":
			if source.Status != "selected" || source.Verification != "verified" || !verifiedEvidence[source.ID] {
				return fmt.Errorf("source %s is not verified claim evidence", source.ID)
			}
		case "candidate_lead":
			if source.Status == "selected" || verifiedEvidence[source.ID] {
				return fmt.Errorf("candidate lead %s cannot be selected or verified claim evidence", source.ID)
			}
		case "context_source":
			if source.Status != "selected" || source.Verification != "verified" || anyEvidence[source.ID] {
				return fmt.Errorf("context source %s must be selected, verified, and uncited as claim evidence", source.ID)
			}
		default:
			return fmt.Errorf("source %s has unknown role %q", source.ID, role)
		}
	}
	for id := range roles {
		if _, ok := sourceByID(report.Run.Sources, id); !ok {
			return fmt.Errorf("source_roles references missing source %s", id)
		}
	}
	return nil
}

// RecordQuery stamps the query when it is recorded, so generated report time cannot
// be mistaken for retrieval time. Call immediately before issuing the search.
func (rec *Recorder) RecordQueryAtCapture(q Query) error {
	q.At = time.Now().UTC().Format(time.RFC3339)
	return rec.RecordQuery(q)
}

// RecordSourceAtCapture stamps source retrieval when its inspected material is recorded.
func (rec *Recorder) RecordSourceAtCapture(source Source) error {
	source.RetrievedAt = time.Now().UTC().Format(time.RFC3339)
	return rec.RecordSource(source)
}
