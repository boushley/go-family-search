package models

type ArtifactMetadata struct {
	ArtifactType ArtifactType `json:"artifactType"`
	Filename     string       `json:"filename"`
	Height       int          `json:"height"`
	Width        int          `json:"width"`
}
