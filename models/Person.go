package models

type Person struct {
	Principal   bool                `json:"principal"`
	Private     bool                `json:"private"`
	Living      bool                `json:"living"`
	Gender      Gender              `json:"gender"`
	Names       []Name              `json:"names"`
	Facts       []Fact              `json:"facts"`
	Fields      []Field             `json:"fields"`
	Display     DisplayProperties   `json:"display"`
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
