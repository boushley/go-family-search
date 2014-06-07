package models

type RelationshipType int

const (
	_ RelationshipType = iota
	RelationshipType_Couple
	RelationshipType_ParentChild
)
