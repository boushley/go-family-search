package models

type HypermediaEnabledData struct {
	Links map[string]Link `json:"links"`
	Id    string          `json:"id"`
}
