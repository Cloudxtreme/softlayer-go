package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Item_Type - <nil>
type SoftLayer_Account_Shipment_Item_Type struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_account_shipment_item_type *SoftLayer_Account_Shipment_Item_Type) String() string {
	return "SoftLayer_Account_Shipment_Item_Type"
}
