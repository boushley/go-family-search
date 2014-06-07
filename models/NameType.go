package models

type NameType int

const (
	_ NameType = iota
	NameType_BirthName
	NameType_DeathName
	NameType_MarriedName
	NameType_AlsoKnownAs
	NameType_Nickname
	NameType_AdoptiveName
	NameType_FormalName
	NameType_ReligiousName
)
