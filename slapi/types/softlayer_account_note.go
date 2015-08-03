package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note - <nil>
type SoftLayer_Account_Note struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Note - <nil>
	Note string `json:"note"`

	// NoteHistory - <nil>
	NoteHistory []*SoftLayer_Account_Note_History `json:"noteHistory"`

	// NoteHistoryCount - no documentation
	NoteHistoryCount uint64 `json:"noteHistoryCount"`

	// NoteType - <nil>
	NoteType *SoftLayer_Account_Note_Type `json:"noteType"`

	// NoteTypeId - <nil>
	NoteTypeId int `json:"noteTypeId"`

	// UserId - <nil>
	UserId int `json:"userId"`
}

func (softlayer_account_note *SoftLayer_Account_Note) String() string {
	return "SoftLayer_Account_Note"
}
