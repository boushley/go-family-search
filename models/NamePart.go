package models

type NamePart struct {
	Value      string       `json:"value"`
	Type       NamePartType `json:"type"`
	Fields     []Field      `json:"fields"`
	Qualifiers []Qualifier  `json:"qualifiers"`
	Id         ID           `json:"id"`
}
