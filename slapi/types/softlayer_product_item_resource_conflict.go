package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict - <nil>
type SoftLayer_Product_Item_Resource_Conflict struct {

	// ItemId - The unique identifier of the item that contains the conflict.
	ItemId int `json:"itemId,omitempty"`

	// Message - no documentation
	Message string `json:"message,omitempty"`

	// PackageId - The unique identifier of the service offering that is associated with the conflict.
	PackageId int `json:"packageId,omitempty"`

	// ResourceTableId - no documentation
	ResourceTableId int `json:"resourceTableId,omitempty"`
}

func (softlayer_product_item_resource_conflict *SoftLayer_Product_Item_Resource_Conflict) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict"
}

// SoftLayer_Product_Item_Resource_Conflict_Extended is SoftLayer_Product_Item_Resource_Conflict with all maskable types.
type SoftLayer_Product_Item_Resource_Conflict_Extended struct {
	SoftLayer_Product_Item_Resource_Conflict

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`
}

func (softlayer_product_item_resource_conflict *SoftLayer_Product_Item_Resource_Conflict_Extended) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict"
}
