package models

type Collection struct {
	Lang        string              `json:"lang"`
	Title       string              `json:"title"`
	Size        int                 `json:"size"`
	Content     []CollectionContent `json:"content"`
	Attribution Attribution         `json:"attribution"`
	Links       map[string]Link     `json:"links"`
	Id          string              `json:"id"`
}
