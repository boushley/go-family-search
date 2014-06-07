package models

type Gedcomx struct {
	Lang               string              `json:"lang"`
	Description        string              `json:"description"`
	Profile            string              `json:"profile"`
	Attribution        Attribution         `json:"attribution"`
	Persons            []Person            `json:"persons"`
	Relationships      []Relationship      `json:"relationships"`
	SourceDescriptions []SourceDescription `json:"sourceDescriptions"`
	Agents             []Agent             `json:"agents"`
	Events             []Event             `json:"events"`
	Places             []PlaceDescription  `json:"places"`
	Documents          []Document          `json:"documents"`
	Collections        []Collection        `json:"collections"`
	Fields             []Field             `json:"fields"`
	RecordDescriptors  []RecordDescriptor  `json:"recordDescriptors"`
	Links              []Link              `json:"links"`
	Id                 ID                  `json:"id"`
}
