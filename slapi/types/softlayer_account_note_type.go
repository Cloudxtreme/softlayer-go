package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Note_Type - <nil>
type SoftLayer_Account_Note_Type struct {

	// BrandId - <nil>
	BrandId int `json:"brandId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - <nil>
	Name string `json:"name"`

	// ValueExpression - <nil>
	ValueExpression string `json:"valueExpression"`
}

func (softlayer_account_note_type *SoftLayer_Account_Note_Type) String() string {
	return "SoftLayer_Account_Note_Type"
}
