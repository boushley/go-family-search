package models

type SourceReference struct {
	Description string          `json:"description"`
	Attribution Attribution     `json:"attribution"`
	Qualifiers  []Qualifier     `json:"qualifiers"`
	Links       map[string]Link `json:"links"`
	Id          string          `json:"id"`
}
