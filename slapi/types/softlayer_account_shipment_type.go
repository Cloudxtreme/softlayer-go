package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Type - <nil>
type SoftLayer_Account_Shipment_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_account_shipment_type *SoftLayer_Account_Shipment_Type) String() string {
	return "SoftLayer_Account_Shipment_Type"
}
