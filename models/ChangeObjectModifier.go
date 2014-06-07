package models

type ChangeObjectModifier int

const (
	_ ChangeObjectModifier = iota
	ChangeObjectModifier_Person
	ChangeObjectModifier_Couple
	ChangeObjectModifier_ChildAndParentsRelationship
)
