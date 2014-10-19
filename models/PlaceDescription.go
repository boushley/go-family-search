package models

type PlaceDescription struct {
	Type                string                 `json:"type"`
	Names               []TextValue            `json:"names"`
	TemporalDescription Date                   `json:"temporalDescription"`
	Latitude            float64                `json:"latitude"`
	Longitude           float64                `json:"longitude"`
	SpatialDescription  ResourceReference      `json:"spatialDescription"`
	Place               ResourceReference      `json:"place"`
	Jurisdiction        ResourceReference      `json:"jurisdiction"`
	Display             PlaceDisplayProperties `json:"display"`
	Extracted           bool                   `json:"extracted"`
	Evidence            []EvidenceReference    `json:"evidence"`
	Media               []SourceReference      `json:"media"`
	Identifiers         map[string]Identifier  `json:"identifiers"`
	Confidence          ConfidenceLevel        `json:"confidence"`
	SortKey             string                 `json:"sortKey"`
	Lang                string                 `json:"lang"`
	Attribution         Attribution            `json:"attribution"`
	Sources             []SourceReference      `json:"sources"`
	Analysis            ResourceReference      `json:"analysis"`
	Notes               []Note                 `json:"notes"`
	Links               map[string]Link        `json:"links"`
	Id                  string                 `json:"id"`
}
