package models

import (
	"time"
)

type Comment struct {
	Text        string            `json:"text"`
	Created     time.Time         `json:"created"`
	Contributor ResourceReference `json:"contributor"`
	Links       []Link            `json:"links"`
	Id          ID                `json:"id"`
}
