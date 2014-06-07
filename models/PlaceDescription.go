package models

type PlaceDescription struct {
	Type                string                 `json:"type"`
	Names               []TextValue            `json:"names"`
	TemporalDescription Date                   `json:"temporalDescription"`
	Latitude            float64                `json:"latitude"`
	Longitude           float64                `json:"longitude"`
	SpatialDescription  ResourceReference      `json:"spatialDescription"`
	Jurisdiction        ResourceReference      `json:"jurisdiction"`
	Display             PlaceDisplayProperties `json:"display"`
	Extracted           bool                   `json:"extracted"`
	Evidence            []EvidenceReference    `json:"evidence"`
	Media               []SourceReference      `json:"media"`
	Identifiers         []Identifier           `json:"identifiers"`
	Confidence          ConfidenceLevel        `json:"confidence"`
	Lang                string                 `json:"lang"`
	Attribution         Attribution            `json:"attribution"`
	Sources             []SourceReference      `json:"sources"`
	Analysis            ResourceReference      `json:"analysis"`
	Notes               []Note                 `json:"notes"`
	Links               []Link                 `json:"links"`
	Id                  ID                     `json:"id"`
}
