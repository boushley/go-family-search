package models

type FieldValue struct {
	Resource    string            `json:"resource"`
	Datatype    string            `json:"datatype"`
	Type        FieldValueType    `json:"type"`
	LabelId     string            `json:"labelId"`
	Text        string            `json:"text"`
	Confidence  ConfidenceLevel   `json:"confidence"`
	SortKey     string            `json:"sortKey"`
	Lang        string            `json:"lang"`
	Attribution Attribution       `json:"attribution"`
	Sources     []SourceReference `json:"sources"`
	Analysis    ResourceReference `json:"analysis"`
	Notes       []Note            `json:"notes"`
	Links       map[string]Link   `json:"links"`
	Id          string            `json:"id"`
}
