package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Item_Type - <nil>
type SoftLayer_Account_Shipment_Item_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_account_shipment_item_type *SoftLayer_Account_Shipment_Item_Type) String() string {
	return "SoftLayer_Account_Shipment_Item_Type"
}
