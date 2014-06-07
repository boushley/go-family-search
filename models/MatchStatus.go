package models

type MatchStatus int

const (
	_ MatchStatus = iota
	MatchStatus_Pending
	MatchStatus_Accepted
	MatchStatus_Rejected
)
