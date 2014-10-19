package models

import "time"

type SourceDescription struct {
	About        string            `json:"about"`
	MediaType    string            `json:"mediaType"`
	ResourceType ResourceType      `json:"resourceType"`
	Citations    []SourceCitation  `json:"citations"`
	Mediator     ResourceReference `json:"mediator"`
	Sources      []SourceReference `json:"sources"`
	Analysis     ResourceReference `json:"analysis"`
	ComponentOf  SourceReference   `json:"componentOf"`
	Titles       []TextValue       `json:"titles"`
	TitleLabel   TextValue         `json:"titleLabel"`
	Notes        []Note            `json:"notes"`
	Attribution  Attribution       `json:"attribution"`
	SortKey      string            `json:"sortKey"`
	Description  []TextValue       `json:"description"`
	Identifiers  []Identifier      `json:"identifiers"`
	Created      time.Time         `json:"created"`
	Modified     time.Time         `json:"modified"`
	Coverage     []Coverage        `json:"coverage"`
	Rights       []string          `json:"rights"`
	Fields       []Field           `json:"fields"`
	Repository   ResourceReference `json:"repository"`
	Descriptor   ResourceReference `json:"descriptor"`
	Links        []Link            `json:"links"`
	Id           ID                `json:"id"`
}
