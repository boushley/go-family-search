package models

type Error struct {
	Code       int    `json:"code"`
	Label      string `json:"label"`
	Message    string `json:"message"`
	Stacktrace string `json:"stacktrace"`
}
