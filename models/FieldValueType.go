package models

type FieldValueType int

const (
	_ FieldValueType = iota
	FieldValueType_Original
	FieldValueType_Interpreted
)
