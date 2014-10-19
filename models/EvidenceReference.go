package models

type EvidenceReference struct {
	ResourceId  string          `json:"resourceId"`
	Resource    string          `json:"resource"`
	Attribution Attribution     `json:"attribution"`
	Links       map[string]Link `json:"links"`
	Id          string          `json:"id"`
}
