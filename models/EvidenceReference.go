package models

import (
	"net/url"
)

type EvidenceReference struct {
	ResourceId  string      `json:"resourceId"`
	Resource    url.URL     `json:"resource"`
	Attribution Attribution `json:"attribution"`
	Links       []Link      `json:"links"`
	Id          ID          `json:"id"`
}
