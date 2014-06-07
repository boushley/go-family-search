package models

type IdentifierType int

const (
	_ IdentifierType = iota
	IdentifierType_Primary
	IdentifierType_Evidence
	IdentifierType_Deprecated
	IdentifierType_Persistent
)
