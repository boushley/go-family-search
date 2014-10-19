package models

type Relationship struct {
	Type        RelationshipType      `json:"type"`
	Person1     ResourceReference     `json:"person1"`
	Person2     ResourceReference     `json:"person2"`
	Facts       []Fact                `json:"facts"`
	Fields      []Field               `json:"fields"`
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
