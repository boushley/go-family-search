package models

type FieldDescriptor struct {
	OriginalLabel string                 `json:"originalLabel"`
	Descriptions  []TextValue            `json:"descriptions"`
	Values        []FieldValueDescriptor `json:"values"`
	Links         map[string]Link        `json:"links"`
	Id            string                 `json:"id"`
}
