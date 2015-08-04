package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note_History - <nil>
type SoftLayer_Account_Note_History struct {

	// UserId - <nil>
	UserId int `json:"userId,omitempty"`

	// AccountNoteId - <nil>
	AccountNoteId int `json:"accountNoteId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Note - <nil>
	Note string `json:"note,omitempty"`
}

func (softlayer_account_note_history *SoftLayer_Account_Note_History) String() string {
	return "SoftLayer_Account_Note_History"
}

// SoftLayer_Account_Note_History_Extended is SoftLayer_Account_Note_History with all maskable types.
type SoftLayer_Account_Note_History_Extended struct {
	SoftLayer_Account_Note_History

	// AccountNote - <nil>
	AccountNote *SoftLayer_Account_Note `json:"accountNote,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`
}

func (softlayer_account_note_history *SoftLayer_Account_Note_History_Extended) String() string {
	return "SoftLayer_Account_Note_History"
}
