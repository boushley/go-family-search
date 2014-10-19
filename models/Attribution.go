package models

import (
	"time"
)

type Attribution struct {
	Contributor   ResourceReference `json:"contributor"`
	Modified      time.Time         `json:"modified"`
	ChangeMessage string            `json:"changeMessage"`
	Id            string            `json:"id"`
}
