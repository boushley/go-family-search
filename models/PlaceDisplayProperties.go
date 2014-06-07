package models

type PlaceDisplayProperties struct {
	FullName string `json:"fullName"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Id       ID     `json:"id"`
}
