package models

type OnlineAccount struct {
	AccountName     string            `json:"accountName"`
	ServiceHomepage ResourceReference `json:"serviceHomepage"`
	Id              ID                `json:"id"`
}
