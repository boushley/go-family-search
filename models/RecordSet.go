package models

type RecordSet struct {
	Lang     string    `json:"lang"`
	Metadata Gedcomx   `json:"metadata"`
	Records  []Gedcomx `json:"records"`
	Links    []Link    `json:"links"`
	Id       ID        `json:"id"`
}
