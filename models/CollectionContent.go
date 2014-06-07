package models

type CollectionContent struct {
	Completeness float32 `json:"completeness"`
	Count        int     `json:"count"`
	ResourceType string  `json:"resourceType"`
	Links        []Link  `json:"links"`
	Id           ID      `json:"id"`
}
