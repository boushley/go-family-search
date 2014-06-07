package models

import (
	"time"
)

type FeatureSet struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Enabled        bool      `json:"enabled"`
	ActivationDate time.Time `json:"activationDate"`
}
