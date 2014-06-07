package models

type Merge struct {
	ResourcesToDelete []ResourceReference `json:"resourcesToDelete"`
	ResourcesToCopy   []ResourceReference `json:"resourcesToCopy"`
}
