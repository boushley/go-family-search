package models

type DiscussionReference struct {
	ResourceId  string      `json:"resourceId"`
	Resource    string      `json:"resource"`
	Attribution Attribution `json:"attribution"`
	Links       []Link      `json:"links"`
	Id          ID          `json:"id"`
}
