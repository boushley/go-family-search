package models

type FieldValueDescriptor struct {
	Optional bool           `json:"optional"`
	Type     FieldValueType `json:"type"`
	LabelId  string         `json:"labelId"`
	Labels   []TextValue    `json:"labels"`
	Links    []Link         `json:"links"`
	Id       ID             `json:"id"`
}
