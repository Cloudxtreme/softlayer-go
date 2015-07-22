package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Type - <nil>
type SoftLayer_Account_Shipment_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Description - <nil>
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_account_shipment_type *SoftLayer_Account_Shipment_Type) GetObject() (*SoftLayer_Account_Shipment_Type, error) {
	var returnValue *SoftLayer_Account_Shipment_Type
	return returnValue, nil
}
