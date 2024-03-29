package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict_Location - <nil>
type SoftLayer_Product_Item_Resource_Conflict_Location struct {

	// PackageId - The unique identifier of the service offering that is associated with the conflict.
	PackageId int `json:"packageId,omitempty"`

	// ResourceTableId - no documentation
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// ItemId - The unique identifier of the item that contains the conflict.
	ItemId int `json:"itemId,omitempty"`

	// Message - no documentation
	Message string `json:"message,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Location `json:"resource,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_item_resource_conflict_location *SoftLayer_Product_Item_Resource_Conflict_Location) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Location"
}
