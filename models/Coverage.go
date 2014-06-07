package models

type Coverage struct {
	RecordType RecordType     `json:"recordType"`
	Spatial    PlaceReference `json:"spatial"`
	Temporal   Date           `json:"temporal"`
	Links      []Link         `json:"links"`
	Id         ID             `json:"id"`
}
