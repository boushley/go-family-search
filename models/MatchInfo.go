package models

type MatchInfo struct {
	Status     MatchStatus `json:"status"`
	Collection string      `json:"collection"`
}
