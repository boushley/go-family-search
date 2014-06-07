package models

type RecordType int

const (
	_ RecordType = iota
	RecordType_Admission
	RecordType_Adoption
	RecordType_Affidavit
	RecordType_Application
	RecordType_Arrival
	RecordType_Bank
	RecordType_Baptism
	RecordType_Birth
	RecordType_Burial
	RecordType_Business
	RecordType_Cemetery
	RecordType_Census
	RecordType_Christening
	RecordType_Confirmation
	RecordType_Correspondence
	RecordType_Death
	RecordType_Departure
	RecordType_Divorce
	RecordType_Duplicate
	RecordType_Draft
	RecordType_Estate
	RecordType_Index
	RecordType_IntendedMarriage
	RecordType_Land
	RecordType_Legal
	RecordType_Marriage
	RecordType_MarriageAffidavit
	RecordType_MarriageAmendment
	RecordType_MarriageBanns
	RecordType_MarriageConsent
	RecordType_MarriageDuplicate
	RecordType_MarriageLicense
	RecordType_MarriageReturns
	RecordType_Membership
	RecordType_Migration
	RecordType_Military
	RecordType_Naturalization
	RecordType_Passenger
	RecordType_Pension
	RecordType_Probate
	RecordType_RelatedDocument
	RecordType_ReligiousCreeds
	RecordType_Roll
	RecordType_Tax
	RecordType_Vital
)
