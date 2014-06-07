package models

type EventRoleType int

const (
	_ EventRoleType = iota
	EventRoleType_Principal
	EventRoleType_Participant
	EventRoleType_Official
	EventRoleType_Witness
)
