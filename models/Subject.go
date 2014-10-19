package models

type Subject struct {
	Extracted   bool                  `json:"extracted"`
	Evidence    []EvidenceReference   `json:"evidence"`
	Media       []SourceReference     `json:"media"`
	Identifiers map[string]Identifier `json:"identifiers"`
	Confidence  ConfidenceLevel       `json:"confidence"`
	SortKey     string                `json:"sortKey"`
	Lang        string                `json:"lang"`
	Attribution Attribution           `json:"attribution"`
	Sources     []SourceReference     `json:"sources"`
	Analysis    ResourceReference     `json:"analysis"`
	Notes       []Note                `json:"notes"`
	Links       map[string]Link       `json:"links"`
	Id          string                `json:"id"`
}
