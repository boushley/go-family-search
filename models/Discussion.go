package models

import (
	"time"
)

type Discussion struct {
	Title            string            `json:"title"`
	Details          string            `json:"details"`
	Created          time.Time         `json:"created"`
	Contributor      ResourceReference `json:"contributor"`
	Modified         time.Time         `json:"modified"`
	NumberOfComments int               `json:"numberOfComments"`
	Comments         []Comment         `json:"comments"`
	Links            map[string]Link   `json:"links"`
	Id               string            `json:"id"`
}
