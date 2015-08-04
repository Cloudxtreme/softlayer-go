package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Shipment_Item - The SoftLayer_Account_Shipment_Item data type contains information
// relating to a shipment's item. Basic information such as addresses, the shipment courier, and any
// tracking information for as shipment is accessible with this data type.
type SoftLayer_Account_Shipment_Item struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// PackageId - no documentation
	PackageId int `json:"packageId,omitempty"`

	// ShipmentItemId - no documentation
	ShipmentItemId int `json:"shipmentItemId,omitempty"`

	// ShipmentId - no documentation
	ShipmentId int `json:"shipmentId,omitempty"`

	// ShipmentItemTypeId - no documentation
	ShipmentItemTypeId int `json:"shipmentItemTypeId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_account_shipment_item *SoftLayer_Account_Shipment_Item) String() string {
	return "SoftLayer_Account_Shipment_Item"
}

// SoftLayer_Account_Shipment_Item_Extended is SoftLayer_Account_Shipment_Item with all maskable types.
type SoftLayer_Account_Shipment_Item_Extended struct {
	SoftLayer_Account_Shipment_Item

	// Shipment - no documentation
	Shipment *SoftLayer_Account_Shipment `json:"shipment,omitempty"`

	// ShipmentItemType - no documentation
	ShipmentItemType *SoftLayer_Account_Shipment_Item_Type `json:"shipmentItemType,omitempty"`
}

func (softlayer_account_shipment_item *SoftLayer_Account_Shipment_Item_Extended) String() string {
	return "SoftLayer_Account_Shipment_Item"
}
