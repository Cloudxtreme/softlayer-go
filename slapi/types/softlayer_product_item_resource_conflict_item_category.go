package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict_Item_Category - <nil>
type SoftLayer_Product_Item_Resource_Conflict_Item_Category struct {

	// PackageId - The unique identifier of the service offering that is associated with the conflict.
	PackageId int `json:"packageId,omitempty"`

	// ResourceTableId - no documentation
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// ItemId - The unique identifier of the item that contains the conflict.
	ItemId int `json:"itemId,omitempty"`

	// Message - no documentation
	Message string `json:"message,omitempty"`
}

func (softlayer_product_item_resource_conflict_item_category *SoftLayer_Product_Item_Resource_Conflict_Item_Category) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item_Category"
}

// SoftLayer_Product_Item_Resource_Conflict_Item_Category_Extended is SoftLayer_Product_Item_Resource_Conflict_Item_Category with all maskable types.
type SoftLayer_Product_Item_Resource_Conflict_Item_Category_Extended struct {
	SoftLayer_Product_Item_Resource_Conflict_Item_Category

	// Resource - An item category that conflicts with a product item.
	Resource *SoftLayer_Product_Item_Category `json:"resource,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_item_resource_conflict_item_category *SoftLayer_Product_Item_Resource_Conflict_Item_Category_Extended) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item_Category"
}
