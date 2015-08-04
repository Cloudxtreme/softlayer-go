package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict_Item - <nil>
type SoftLayer_Product_Item_Resource_Conflict_Item struct {

	// ItemId - The unique identifier of the item that contains the conflict.
	ItemId int `json:"itemId,omitempty"`

	// Message - no documentation
	Message string `json:"message,omitempty"`

	// PackageId - The unique identifier of the service offering that is associated with the conflict.
	PackageId int `json:"packageId,omitempty"`

	// ResourceTableId - no documentation
	ResourceTableId int `json:"resourceTableId,omitempty"`
}

func (softlayer_product_item_resource_conflict_item *SoftLayer_Product_Item_Resource_Conflict_Item) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item"
}

// SoftLayer_Product_Item_Resource_Conflict_Item_Extended is SoftLayer_Product_Item_Resource_Conflict_Item with all maskable types.
type SoftLayer_Product_Item_Resource_Conflict_Item_Extended struct {
	SoftLayer_Product_Item_Resource_Conflict_Item

	// Resource - A product item that conflicts with another product item.
	Resource *SoftLayer_Product_Item `json:"resource,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_item_resource_conflict_item *SoftLayer_Product_Item_Resource_Conflict_Item_Extended) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item"
}
