package models

type ConfidenceLevel int

const (
	_ ConfidenceLevel = iota
	ConfidenceLevel_High
	ConfidenceLevel_Medium
	ConfidenceLevel_Low
)
