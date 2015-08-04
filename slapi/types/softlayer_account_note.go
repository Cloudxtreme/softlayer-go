package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note - <nil>
type SoftLayer_Account_Note struct {

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Note - <nil>
	Note string `json:"note,omitempty"`

	// NoteTypeId - <nil>
	NoteTypeId int `json:"noteTypeId,omitempty"`

	// UserId - <nil>
	UserId int `json:"userId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// NoteType - <nil>
	NoteType *SoftLayer_Account_Note_Type `json:"noteType,omitempty"`

	// NoteHistoryCount - no documentation
	NoteHistoryCount uint64 `json:"noteHistoryCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// NoteHistory - <nil>
	NoteHistory []*SoftLayer_Account_Note_History `json:"noteHistory,omitempty"`
}

func (softlayer_account_note *SoftLayer_Account_Note) String() string {
	return "SoftLayer_Account_Note"
}
