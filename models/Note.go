package models

type Note struct {
	Lang        string      `json:"lang"`
	Subject     string      `json:"subject"`
	Text        string      `json:"text"`
	Attribution Attribution `json:"attribution"`
	Links       []Link      `json:"links"`
	Id          ID          `json:"id"`
}
