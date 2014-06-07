package models

type HypermediaEnabledData struct {
	Links []Link `json:"links"`
	Id    ID     `json:"id"`
}
