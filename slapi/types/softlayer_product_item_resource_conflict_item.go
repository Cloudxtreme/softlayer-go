package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict_Item - <nil>
type SoftLayer_Product_Item_Resource_Conflict_Item struct {
}

func (softlayer_product_item_resource_conflict_item *SoftLayer_Product_Item_Resource_Conflict_Item) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item"
}

// SoftLayer_Product_Item_Resource_Conflict_Item_Extended is SoftLayer_Product_Item_Resource_Conflict_Item with all maskable types.
type SoftLayer_Product_Item_Resource_Conflict_Item_Extended struct {
	SoftLayer_Product_Item_Resource_Conflict_Item

	// Resource - A product item that conflicts with another product item.
	Resource *SoftLayer_Product_Item `json:"resource"`
}

func (softlayer_product_item_resource_conflict_item *SoftLayer_Product_Item_Resource_Conflict_Item_Extended) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item"
}
