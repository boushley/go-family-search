package models

type RecordSet struct {
	Lang     string          `json:"lang"`
	Metadata Gedcomx         `json:"metadata"`
	Records  []Gedcomx       `json:"records"`
	Links    map[string]Link `json:"links"`
	Id       string          `json:"id"`
}
