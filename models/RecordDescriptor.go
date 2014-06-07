package models

type RecordDescriptor struct {
	Lang   string            `json:"lang"`
	Fields []FieldDescriptor `json:"fields"`
	Links  []Link            `json:"links"`
	Id     ID                `json:"id"`
}
