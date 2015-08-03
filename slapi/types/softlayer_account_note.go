package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note - <nil>
type SoftLayer_Account_Note struct {

	// Note - <nil>
	Note string `json:"note"`

	// NoteTypeId - <nil>
	NoteTypeId int `json:"noteTypeId"`

	// UserId - <nil>
	UserId int `json:"userId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Id - <nil>
	Id int `json:"id"`
}

// SoftLayer_Account_Note_Extended is SoftLayer_Account_Note with all maskable types.
type SoftLayer_Account_Note_Extended struct {
	SoftLayer_Account_Note

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// NoteHistory - <nil>
	NoteHistory []*SoftLayer_Account_Note_History `json:"noteHistory"`

	// NoteHistoryCount - no documentation
	NoteHistoryCount uint64 `json:"noteHistoryCount"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`

	// NoteType - <nil>
	NoteType *SoftLayer_Account_Note_Type `json:"noteType"`
}

func (softlayer_account_note *SoftLayer_Account_Note) String() string {
	return "SoftLayer_Account_Note"
}
