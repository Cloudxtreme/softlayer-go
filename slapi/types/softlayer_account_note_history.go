package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note_History - <nil>
type SoftLayer_Account_Note_History struct {

	// AccountNote - <nil>
	AccountNote *SoftLayer_Account_Note `json:"accountNote"`

	// AccountNoteId - <nil>
	AccountNoteId int `json:"accountNoteId"`

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

	// UserId - <nil>
	UserId int `json:"userId"`
}
