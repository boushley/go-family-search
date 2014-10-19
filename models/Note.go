package models

type Note struct {
	Lang        string          `json:"lang"`
	Subject     string          `json:"subject"`
	Text        string          `json:"text"`
	Attribution Attribution     `json:"attribution"`
	Links       map[string]Link `json:"links"`
	Id          string          `json:"id"`
}
