package models

import (
	"net/url"
)

type DiscussionReference struct {
	ResourceId  string      `json:"resourceId"`
	Resource    url.URL     `json:"resource"`
	Attribution Attribution `json:"attribution"`
	Links       []Link      `json:"links"`
	Id          ID          `json:"id"`
}
