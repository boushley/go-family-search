package models

type ArtifactType int

const (
	_ ArtifactType = iota
	ArtifactType_Audio
	ArtifactType_Document
	ArtifactType_Image
	ArtifactType_Portrait
	ArtifactType_Story
	ArtifactType_Video
)
