package models

type Date struct {
	Original   string      `json:"original"`
	Formal     string      `json:"formal"`
	Normalized []TextValue `json:"normalized"`
	Fields     []Field     `json:"fields"`
	Id         ID          `json:"id"`
}
