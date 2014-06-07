package models

type GenderType int

const (
	_ GenderType = iota
	GenderType_Male
	GenderType_Female
	GenderType_Unknown
)
