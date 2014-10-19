package models

type Link struct {
	Template string `json:"template"`
	Allow    string `json:"allow"`
	Count    int    `json:"count"`
	Accept   string `json:"accept"`
	Type     string `json:"type"`
	Hreflang string `json:"hreflang"`
	Title    string `json:"title"`
	Results  int    `json:"results"`
	Rel      string `json:"rel"`
	Offset   int    `json:"offset"`
	Href     string `json:"href"`
}
