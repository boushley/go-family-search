package models

import (
	"net/url"
)

type ResourceReference struct {
	ResourceId string  `json:"resourceId"`
	Resource   url.URL `json:"resource"`
}
