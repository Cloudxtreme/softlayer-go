package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note - <nil>
type SoftLayer_Account_Note struct {

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// UserId - <nil>
	UserId int `json:"userId,omitempty"`

	// Note - <nil>
	Note string `json:"note,omitempty"`

	// NoteTypeId - <nil>
	NoteTypeId int `json:"noteTypeId,omitempty"`
}

func (softlayer_account_note *SoftLayer_Account_Note) String() string {
	return "SoftLayer_Account_Note"
}

// SoftLayer_Account_Note_Extended is SoftLayer_Account_Note with all maskable types.
type SoftLayer_Account_Note_Extended struct {
	SoftLayer_Account_Note

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// NoteHistory - <nil>
	NoteHistory []*SoftLayer_Account_Note_History `json:"noteHistory,omitempty"`

	// NoteType - <nil>
	NoteType *SoftLayer_Account_Note_Type `json:"noteType,omitempty"`

	// NoteHistoryCount - no documentation
	NoteHistoryCount uint64 `json:"noteHistoryCount,omitempty"`
}

func (softlayer_account_note *SoftLayer_Account_Note_Extended) String() string {
	return "SoftLayer_Account_Note"
}
