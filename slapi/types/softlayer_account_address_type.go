package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Address_Type - <nil>
type SoftLayer_Account_Address_Type struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_account_address_type *SoftLayer_Account_Address_Type) String() string {
	return "SoftLayer_Account_Address_Type"
}
