package models

type Field struct {
	Type   FieldType       `json:"type"`
	Values []FieldValue    `json:"values"`
	Links  map[string]Link `json:"links"`
	Id     string          `json:"id"`
}
