package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Address_Type - <nil>
type SoftLayer_Account_Address_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_account_address_type *SoftLayer_Account_Address_Type) GetObject() (*SoftLayer_Account_Address_Type, error) {
	var returnValue *SoftLayer_Account_Address_Type
	return returnValue, nil
}
