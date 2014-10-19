package models

type NameForm struct {
	Lang     string     `json:"lang"`
	FullText string     `json:"fullText"`
	Parts    []NamePart `json:"parts"`
	Fields   []Field    `json:"fields"`
	Id       string     `json:"id"`
}
