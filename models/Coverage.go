package models

type Coverage struct {
	RecordType RecordType      `json:"recordType"`
	Spatial    PlaceReference  `json:"spatial"`
	Temporal   Date            `json:"temporal"`
	Links      map[string]Link `json:"links"`
	Id         string          `json:"id"`
}
