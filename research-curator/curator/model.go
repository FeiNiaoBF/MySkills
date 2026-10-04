package curator

// Run is the versioned JSON ledger. Arrays must be present, even when empty.
type Run struct {
	Version         string           `json:"version"`
	ID              string           `json:"id"`
	Metadata        Metadata         `json:"metadata"`
	RejectedSources []RejectedSource `json:"rejected_sources"`
	Contract        Contract         `json:"contract"`
	Sources         []Source         `json:"sources"`
	Claims          []Claim          `json:"claims"`
	Conclusions     []Conclusion     `json:"conclusions"`
	Graph           Graph            `json:"graph"`
	Queries         []Query          `json:"queries"`
	Events          []Event          `json:"events"`
	Decisions       []Decision       `json:"decisions"`
	Stages          []Stage          `json:"stages"`
	Coverage        Coverage         `json:"coverage"`
}

// Metadata describes execution, not evidence truth. completed_at is empty until finalized.
type Metadata struct {
	CreatedAt        string   `json:"created_at"`
	CompletedAt      string   `json:"completed_at"`
	Status           string   `json:"status"`
	Tools            []string `json:"tools"`
	SearchProvenance string   `json:"search_provenance"`
}

// RejectedSource indexes all nonselected candidates without discarding their full records.
type RejectedSource struct {
	ID     string `json:"id"`
	Reason string `json:"reason"`
}
type Contract struct {
	Question    string     `json:"question"`
	Types       []string   `json:"types"`
	Excludes    []string   `json:"excludes"`
	Freshness   string     `json:"freshness"`
	Preferences []string   `json:"preferences"`
	Quantity    int        `json:"quantity"`
	Depth       string     `json:"depth"`
	Questions   []Question `json:"questions"`
}
type Question struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type Source struct {
	ID             string   `json:"id"`
	URL            string   `json:"url"`
	Title          string   `json:"title"`
	Type           string   `json:"type"`
	Original       bool     `json:"original"`
	PublishedAt    string   `json:"published_at,omitempty"`
	RetrievedAt    string   `json:"retrieved_at"`
	Content        string   `json:"content"`
	ContentHash    string   `json:"content_hash,omitempty"`
	UpstreamIDs    []string `json:"upstream_ids"`
	Freshness      string   `json:"freshness"`
	Fit            float64  `json:"fit"`
	Evidence       float64  `json:"evidence"`
	Utility        float64  `json:"utility"`
	FitReason      string   `json:"fit_reason"`
	EvidenceReason string   `json:"evidence_reason"`
	UtilityReason  string   `json:"utility_reason"`
	Status         string   `json:"status"`
	Reason         string   `json:"reason"`
	DuplicateOf    string   `json:"duplicate_of,omitempty"`
	Verification   string   `json:"verification"`
}
type Claim struct {
	ID          string     `json:"id"`
	Text        string     `json:"text"`
	QuestionIDs []string   `json:"question_ids"`
	Evidence    []Evidence `json:"evidence"`
}
type Evidence struct {
	SourceID     string `json:"source_id"`
	Quote        string `json:"quote"`
	Locator      string `json:"locator"`
	Relation     string `json:"relation"`
	Verification string `json:"verification"`
}
type Conclusion struct {
	ID       string   `json:"id"`
	Text     string   `json:"text"`
	ClaimIDs []string `json:"claim_ids"`
	Status   string   `json:"status"`
}
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
type Node struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label string `json:"label"`
}
type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
}
type Query struct {
	ID         string `json:"id"`
	QuestionID string `json:"question_id"`
	Text       string `json:"text"`
	At         string `json:"at"`
}
type Event struct {
	ID          string `json:"id"`
	At          string `json:"at"`
	Stage       string `json:"stage"`
	Action      string `json:"action"`
	Detail      string `json:"detail"`
	BeforeCount *int   `json:"before_count,omitempty"`
	AfterCount  *int   `json:"after_count,omitempty"`
	CountScope  string `json:"count_scope,omitempty"`
}
type Decision struct {
	ID       string `json:"id"`
	At       string `json:"at"`
	SourceID string `json:"source_id"`
	Action   string `json:"action"`
	Reason   string `json:"reason"`
}
type Stage struct {
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	EvidenceIDs []string `json:"evidence_ids"`
	Reason      string   `json:"reason"`
}
type Coverage struct {
	Status    string             `json:"status"`
	Questions []QuestionCoverage `json:"questions"`
	Warnings  []string           `json:"warnings"`
}
type QuestionCoverage struct {
	QuestionID         string   `json:"question_id"`
	ClaimIDs           []string `json:"claim_ids"`
	IndependentSources int      `json:"independent_sources"`
	Gaps               []string `json:"gaps"`
}

var StageNames = []string{"Requirement First", "Discovery", "Normalize", "Deduplicate", "Qualify", "Verify", "Claim Extraction", "Evidence Graph", "Coverage Check", "Rank", "Synthesis"}
