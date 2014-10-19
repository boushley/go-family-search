package models

import (
	"time"
)

type SourceDescription struct {
	About        string                `json:"about"`
	Lang         string                `json:"lang"`
	MediaType    string                `json:"mediaType"`
	SortKey      string                `json:"sortKey"`
	ResourceType ResourceType          `json:"resourceType"`
	Citations    []SourceCitation      `json:"citations"`
	Mediator     ResourceReference     `json:"mediator"`
	Sources      []SourceReference     `json:"sources"`
	Analysis     ResourceReference     `json:"analysis"`
	ComponentOf  SourceReference       `json:"componentOf"`
	Titles       []TextValue           `json:"titles"`
	TitleLabel   TextValue             `json:"titleLabel"`
	Notes        []Note                `json:"notes"`
	Attribution  Attribution           `json:"attribution"`
	Descriptions []TextValue           `json:"descriptions"`
	Identifiers  map[string]Identifier `json:"identifiers"`
	Created      time.Time             `json:"created"`
	Modified     time.Time             `json:"modified"`
	Coverage     []Coverage            `json:"coverage"`
	Rights       []string              `json:"rights"`
	Fields       []Field               `json:"fields"`
	Repository   ResourceReference     `json:"repository"`
	Descriptor   ResourceReference     `json:"descriptor"`
	Links        map[string]Link       `json:"links"`
	Id           string                `json:"id"`
}
