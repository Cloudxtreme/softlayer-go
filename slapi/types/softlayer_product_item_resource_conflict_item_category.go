package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Resource_Conflict_Item_Category - <nil>
type SoftLayer_Product_Item_Resource_Conflict_Item_Category struct {
}

// SoftLayer_Product_Item_Resource_Conflict_Item_Category_Extended is SoftLayer_Product_Item_Resource_Conflict_Item_Category with all maskable types.
type SoftLayer_Product_Item_Resource_Conflict_Item_Category_Extended struct {
	SoftLayer_Product_Item_Resource_Conflict_Item_Category

	// Resource - An item category that conflicts with a product item.
	Resource *SoftLayer_Product_Item_Category `json:"resource"`
}

func (softlayer_product_item_resource_conflict_item_category *SoftLayer_Product_Item_Resource_Conflict_Item_Category) String() string {
	return "SoftLayer_Product_Item_Resource_Conflict_Item_Category"
}
