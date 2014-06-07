package models

type EventType int

const (
	_ EventType = iota
	EventType_Adoption
	EventType_AdultChristening
	EventType_Annulment
	EventType_Baptism
	EventType_BarMitzvah
	EventType_BatMitzvah
	EventType_Birth
	EventType_Blessing
	EventType_Burial
	EventType_Census
	EventType_Christening
	EventType_Circumcision
	EventType_Confirmation
	EventType_Cremation
	EventType_Death
	EventType_Divorce
	EventType_DivorceFiling
	EventType_Education
	EventType_Engagement
	EventType_Emigration
	EventType_Excommunication
	EventType_FirstCommunion
	EventType_Funeral
	EventType_Immigration
	EventType_LandTransaction
	EventType_Marriage
	EventType_MilitaryAward
	EventType_MilitaryDischarge
	EventType_Mission
	EventType_MoveFrom
	EventType_MoveTo
	EventType_Naturalization
	EventType_Ordination
	EventType_Retirement
)
