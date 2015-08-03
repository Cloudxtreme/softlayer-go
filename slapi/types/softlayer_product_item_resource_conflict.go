package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict - <nil>
type SoftLayer_Product_Item_Resource_Conflict struct {

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ItemId - The unique identifier of the item that contains the conflict.
	ItemId int `json:"itemId"`

	// Message - no documentation
	Message string `json:"message"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`

	// PackageId - The unique identifier of the service offering that is associated with the conflict.
	PackageId int `json:"packageId"`

	// ResourceTableId - no documentation
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_product_item_resource_conflict *SoftLayer_Product_Item_Resource_Conflict) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict"
}
