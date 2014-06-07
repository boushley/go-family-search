package models

type MergeAnalysis struct {
	SurvivorResources    []ResourceReference `json:"survivorResources"`
	DuplicateResources   []ResourceReference `json:"duplicateResources"`
	ConflictingResources []MergeConflict     `json:"conflictingResources"`
	Survivor             ResourceReference   `json:"survivor"`
	Duplicate            ResourceReference   `json:"duplicate"`
}
