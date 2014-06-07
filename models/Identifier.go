package models

type Identifier struct {
	Type  IdentifierType `json:"type"`
	Value string         `json:"value"`
}
