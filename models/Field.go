package models

type Field struct {
	Type   FieldType    `json:"type"`
	Values []FieldValue `json:"values"`
	Links  []Link       `json:"links"`
	Id     ID           `json:"id"`
}
