package models

type Document struct {
	TextType    string            `json:"textType"`
	Extracted   bool              `json:"extracted"`
	Type        DocumentType      `json:"type"`
	Text        string            `json:"text"`
	Confidence  ConfidenceLevel   `json:"confidence"`
	Lang        string            `json:"lang"`
	Attribution Attribution       `json:"attribution"`
	Sources     []SourceReference `json:"sources"`
	Analysis    ResourceReference `json:"analysis"`
	Notes       []Note            `json:"notes"`
	Links       []Link            `json:"links"`
	Id          ID                `json:"id"`
}
