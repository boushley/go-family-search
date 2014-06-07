package models

type Name struct {
	Type        NameType          `json:"type"`
	Preferred   bool              `json:"preferred"`
	Date        Date              `json:"date"`
	NameForms   []NameForm        `json:"nameForms"`
	Confidence  ConfidenceLevel   `json:"confidence"`
	Lang        string            `json:"lang"`
	Attribution Attribution       `json:"attribution"`
	Sources     []SourceReference `json:"sources"`
	Analysis    ResourceReference `json:"analysis"`
	Notes       []Note            `json:"notes"`
	Links       []Link            `json:"links"`
	Id          ID                `json:"id"`
}
