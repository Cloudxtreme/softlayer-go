package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Status - <nil>
type SoftLayer_Account_Shipment_Status struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_account_shipment_status *SoftLayer_Account_Shipment_Status) String() string {
	return "SoftLayer_Account_Shipment_Status"
}
