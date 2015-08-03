package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Item - The SoftLayer_Account_Shipment_Item data type contains information
// relating to a shipment's item. Basic information such as addresses, the shipment courier, and any
// tracking information for as shipment is accessible with this data type.
type SoftLayer_Account_Shipment_Item struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// PackageId - no documentation
	PackageId int `json:"packageId"`

	// Shipment - no documentation
	Shipment *SoftLayer_Account_Shipment `json:"shipment"`

	// ShipmentId - no documentation
	ShipmentId int `json:"shipmentId"`

	// ShipmentItemId - no documentation
	ShipmentItemId int `json:"shipmentItemId"`

	// ShipmentItemType - no documentation
	ShipmentItemType *SoftLayer_Account_Shipment_Item_Type `json:"shipmentItemType"`

	// ShipmentItemTypeId - no documentation
	ShipmentItemTypeId int `json:"shipmentItemTypeId"`
}

func (softlayer_account_shipment_item *SoftLayer_Account_Shipment_Item) String() string {
	return "SoftLayer_Account_Shipment_Item"
}
