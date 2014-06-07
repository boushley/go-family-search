package models

type ResourceType int

const (
	_ ResourceType = iota
	ResourceType_Record
	ResourceType_Collection
	ResourceType_DigitalArtifact
	ResourceType_PhysicalArtifact
	ResourceType_Person
)
