package models

type FieldDescriptor struct {
	OriginalLabel string                 `json:"originalLabel"`
	Description   []TextValue            `json:"description"`
	Values        []FieldValueDescriptor `json:"values"`
	Links         []Link                 `json:"links"`
	Id            ID                     `json:"id"`
}
