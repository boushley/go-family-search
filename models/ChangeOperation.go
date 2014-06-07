package models

type ChangeOperation int

const (
	_ ChangeOperation = iota
	ChangeOperation_Create
	ChangeOperation_Read
	ChangeOperation_Update
	ChangeOperation_Delete
	ChangeOperation_Merge
	ChangeOperation_Unmerge
	ChangeOperation_Restore
)
