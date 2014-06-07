package models

type Address struct {
	City            string `json:"city"`
	Country         string `json:"country"`
	PostalCode      string `json:"postalCode"`
	StateOrProvince string `json:"stateOrProvince"`
	Street          string `json:"street"`
	Street2         string `json:"street2"`
	Street3         string `json:"street3"`
	Street4         string `json:"street4"`
	Street5         string `json:"street5"`
	Street6         string `json:"street6"`
	Value           string `json:"value"`
	Id              ID     `json:"id"`
}
