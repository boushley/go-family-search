package models

type FieldValueDescriptor struct {
	Optional bool            `json:"optional"`
	Type     FieldValueType  `json:"type"`
	LabelId  string          `json:"labelId"`
	Labels   []TextValue     `json:"labels"`
	Links    map[string]Link `json:"links"`
	Id       string          `json:"id"`
}
