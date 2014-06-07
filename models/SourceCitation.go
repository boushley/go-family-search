package models

type SourceCitation struct {
	Lang             string            `json:"lang"`
	CitationTemplate ResourceReference `json:"citationTemplate"`
	Fields           []CitationField   `json:"fields"`
	Value            string            `json:"value"`
	Links            []Link            `json:"links"`
	Id               ID                `json:"id"`
}
