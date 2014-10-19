package models

type Agent struct {
	Accounts    []OnlineAccount       `json:"accounts"`
	Addresses   []Address             `json:"addresses"`
	Emails      []ResourceReference   `json:"emails"`
	Homepage    ResourceReference     `json:"homepage"`
	Identifiers map[string]Identifier `json:"identifiers"`
	Names       []TextValue           `json:"names"`
	Openid      ResourceReference     `json:"openid"`
	Phones      []ResourceReference   `json:"phones"`
	Links       map[string]Link       `json:"links"`
	Id          string                `json:"id"`
}
