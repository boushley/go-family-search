package models

type SourceCitation struct {
	Lang             string            `json:"lang"`
	CitationTemplate ResourceReference `json:"citationTemplate"`
	Fields           []CitationField   `json:"fields"`
	Value            string            `json:"value"`
	Links            map[string]Link   `json:"links"`
	Id               string            `json:"id"`
}
