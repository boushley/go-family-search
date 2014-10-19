package models

type Link struct {
	Hreflang string `json:"hreflang"`
	Template string `json:"template"`
	Title    string `json:"title"`
	Allow    string `json:"allow"`
	Accept   string `json:"accept"`
	Rel      string `json:"rel"`
	Type     string `json:"type"`
	Href     string `json:"href"`
}
