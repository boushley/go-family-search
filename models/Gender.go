package models

type Gender struct {
	Type        GenderType        `json:"type"`
	Fields      []Field           `json:"fields"`
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
