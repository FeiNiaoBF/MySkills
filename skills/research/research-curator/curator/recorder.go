package curator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Recorder appends validated records and atomically persists a ledger. It is not concurrency safe.
type Recorder struct{ Run *Run }

func (rec *Recorder) RecordEvent(ev Event) error {
	if rec.Run == nil {
		return fmt.Errorf("nil run")
	}
	trial := *rec.Run
	trial.Events = append(append([]Event{}, rec.Run.Events...), ev)
	if e := Validate(&trial); e != nil {
		return e
	}
	rec.Run.Events = trial.Events
	return nil
}
func (rec *Recorder) RecordQuery(q Query) error {
	if rec.Run == nil {
		return fmt.Errorf("nil run")
	}
	trial := *rec.Run
	trial.Queries = append(append([]Query{}, rec.Run.Queries...), q)
	trial.Graph.Nodes = append(append([]Node{}, rec.Run.Graph.Nodes...), Node{ID: q.ID, Type: "Query", Label: q.Text})
	trial.Graph.Edges = append(append([]Edge{}, rec.Run.Graph.Edges...), Edge{From: q.QuestionID, To: q.ID, Type: "searched_by"})
	for _, sid := range q.CandidateSourceIDs {
		trial.Graph.Edges = append(trial.Graph.Edges, Edge{From: q.ID, To: sid, Type: "candidate"})
	}
	if e := Validate(&trial); e != nil {
		return e
	}
	rec.Run.Queries = trial.Queries
	rec.Run.Graph = trial.Graph
	return nil
}
func (rec *Recorder) RecordSource(source Source) error {
	if rec.Run == nil {
		return fmt.Errorf("nil run")
	}
	trial := *rec.Run
	trial.Sources = append(append([]Source{}, rec.Run.Sources...), source)
	trial.Graph.Nodes = append(append([]Node{}, rec.Run.Graph.Nodes...), Node{ID: source.ID, Type: "Source", Label: source.Title})
	if e := Validate(&trial); e != nil {
		return e
	}
	rec.Run.Sources = trial.Sources
	rec.Run.Graph.Nodes = trial.Graph.Nodes
	return nil
}

func (rec *Recorder) RecordAdjudication(a ConflictAdjudication) error {
	if rec.Run == nil {
		return fmt.Errorf("nil run")
	}
	trial := *rec.Run
	trial.Adjudications = append(append([]ConflictAdjudication{}, rec.Run.Adjudications...), a)
	if e := Validate(&trial); e != nil {
		return e
	}
	rec.Run.Adjudications = trial.Adjudications
	return nil
}
func (rec *Recorder) RecordDecision(d Decision) error {
	if rec.Run == nil {
		return fmt.Errorf("nil run")
	}
	trial := *rec.Run
	trial.Decisions = append(append([]Decision{}, rec.Run.Decisions...), d)
	if e := Validate(&trial); e != nil {
		return e
	}
	rec.Run.Decisions = trial.Decisions
	return nil
}
func (rec *Recorder) Save(path string) error {
	if e := Validate(rec.Run); e != nil {
		return e
	}
	data, e := json.MarshalIndent(rec.Run, "", "  ")
	if e != nil {
		return e
	}
	return WriteFileAtomic(path, append(data, '\n'))
}

// WriteFileAtomic writes mode 0600 into the destination directory then renames.
func WriteFileAtomic(path string, data []byte) error {
	f, e := os.CreateTemp(filepath.Dir(path), ".research-curator-*")
	if e != nil {
		return e
	}
	name := f.Name()
	defer os.Remove(name)
	if _, e = f.Write(data); e != nil {
		f.Close()
		return e
	}
	if e = f.Sync(); e != nil {
		f.Close()
		return e
	}
	if e = f.Close(); e != nil {
		return e
	}
	return os.Rename(name, path)
}
