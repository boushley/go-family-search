package models

type MergeConflict struct {
	SurvivorResource  ResourceReference `json:"survivorResource"`
	DuplicateResource ResourceReference `json:"duplicateResource"`
}
