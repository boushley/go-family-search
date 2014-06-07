package models

type DocumentType int

const (
	_ DocumentType = iota
	DocumentType_Abstract
	DocumentType_Translation
	DocumentType_Transcription
	DocumentType_Analysis
)
