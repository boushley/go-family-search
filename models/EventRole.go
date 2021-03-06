package models

type EventRole struct {
	Type        EventRoleType     `json:"type"`
	Person      ResourceReference `json:"person"`
	Details     string            `json:"details"`
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
