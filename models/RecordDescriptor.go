package models

type RecordDescriptor struct {
	Lang   string            `json:"lang"`
	Fields []FieldDescriptor `json:"fields"`
	Links  map[string]Link   `json:"links"`
	Id     string            `json:"id"`
}
