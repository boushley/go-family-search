package models

type NamePartType int

const (
	_ NamePartType = iota
	NamePartType_Prefix
	NamePartType_Suffix
	NamePartType_Given
	NamePartType_Surname
)
