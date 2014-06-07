package models

type Fact struct {
	Primary     bool              `json:"primary"`
	Type        FactType          `json:"type"`
	Date        Date              `json:"date"`
	Place       PlaceReference    `json:"place"`
	Value       string            `json:"value"`
	Qualifiers  []Qualifier       `json:"qualifiers"`
	Fields      []Field           `json:"fields"`
	Confidence  ConfidenceLevel   `json:"confidence"`
	Lang        string            `json:"lang"`
	Attribution Attribution       `json:"attribution"`
	Sources     []SourceReference `json:"sources"`
	Analysis    ResourceReference `json:"analysis"`
	Notes       []Note            `json:"notes"`
	Links       []Link            `json:"links"`
	Id          ID                `json:"id"`
}
