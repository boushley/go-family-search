package models

type Agent struct {
	Accounts    []OnlineAccount     `json:"accounts"`
	Addresses   []Address           `json:"addresses"`
	Emails      []ResourceReference `json:"emails"`
	Homepage    ResourceReference   `json:"homepage"`
	Identifiers []Identifier        `json:"identifiers"`
	Names       []TextValue         `json:"names"`
	Openid      ResourceReference   `json:"openid"`
	Phones      []ResourceReference `json:"phones"`
	Links       []Link              `json:"links"`
	Id          ID                  `json:"id"`
}
