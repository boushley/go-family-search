package models

type SourceReference struct {
	Description string      `json:"description"`
	Attribution Attribution `json:"attribution"`
	Qualifiers  []Qualifier `json:"qualifiers"`
	Links       []Link      `json:"links"`
	Id          ID          `json:"id"`
}
