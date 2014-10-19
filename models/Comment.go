package models

import (
	"time"
)

type Comment struct {
	Text        string            `json:"text"`
	Created     time.Time         `json:"created"`
	Contributor ResourceReference `json:"contributor"`
	Links       map[string]Link   `json:"links"`
	Id          string            `json:"id"`
}
