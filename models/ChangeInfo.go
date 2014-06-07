package models

type ChangeInfo struct {
	ObjectModifier ChangeObjectModifier `json:"objectModifier"`
	Operation      ChangeOperation      `json:"operation"`
	Reason         string               `json:"reason"`
	ObjectType     ChangeObjectType     `json:"objectType"`
	Original       ResourceReference    `json:"original"`
	Parent         ResourceReference    `json:"parent"`
	Removed        ResourceReference    `json:"removed"`
	Resulting      ResourceReference    `json:"resulting"`
}
