package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// CreateObject - <nil>
func (softlayer_account_note *SoftLayer_Account_Note) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Account_Note) (*SoftLayer_Account_Note, error) {
	var returnValue *SoftLayer_Account_Note
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_account_note *SoftLayer_Account_Note) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_account_note *SoftLayer_Account_Note) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Account_Note) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_note *SoftLayer_Account_Note) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Account_Note, error) {
	var returnValue *SoftLayer_Account_Note
	return returnValue, nil
}
