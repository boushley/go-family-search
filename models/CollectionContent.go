package models

type CollectionContent struct {
	Completeness float32         `json:"completeness"`
	Count        int             `json:"count"`
	ResourceType string          `json:"resourceType"`
	Links        map[string]Link `json:"links"`
	Id           string          `json:"id"`
}
