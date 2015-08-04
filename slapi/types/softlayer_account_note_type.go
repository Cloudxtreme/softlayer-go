package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note_Type - <nil>
type SoftLayer_Account_Note_Type struct {

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// ValueExpression - <nil>
	ValueExpression string `json:"valueExpression,omitempty"`

	// BrandId - <nil>
	BrandId int `json:"brandId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_account_note_type *SoftLayer_Account_Note_Type) String() string {
	return "SoftLayer_Account_Note_Type"
}
