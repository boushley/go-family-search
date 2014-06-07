package models

type Subject struct {
	Extracted   bool                `json:"extracted"`
	Evidence    []EvidenceReference `json:"evidence"`
	Media       []SourceReference   `json:"media"`
	Identifiers []Identifier        `json:"identifiers"`
	Confidence  ConfidenceLevel     `json:"confidence"`
	Lang        string              `json:"lang"`
	Attribution Attribution         `json:"attribution"`
	Sources     []SourceReference   `json:"sources"`
	Analysis    ResourceReference   `json:"analysis"`
	Notes       []Note              `json:"notes"`
	Links       []Link              `json:"links"`
	Id          ID                  `json:"id"`
}
