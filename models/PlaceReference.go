package models

type PlaceReference struct {
	Description string      `json:"description"`
	Original    string      `json:"original"`
	Normalized  []TextValue `json:"normalized"`
	Fields      []Field     `json:"fields"`
	Id          string      `json:"id"`
}
