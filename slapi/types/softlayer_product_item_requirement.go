package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Requirement - The SoftLayer_Product_Item_Requirement data type contains
// information relating to what requirements, if any, exist for an item. The requiredItemId local
// property is the item id that is required.
type SoftLayer_Product_Item_Requirement struct {

	// Id - no documentation
	Id int `json:"id"`

	// ItemId - This is the id of the item affected by the requirement.
	ItemId int `json:"itemId"`

	// Message - This is a custom message to display to the user when this requirement shortfall arises.
	Message string `json:"message"`

	// RequiredItemId - no documentation
	RequiredItemId int `json:"requiredItemId"`
}

// SoftLayer_Product_Item_Requirement_Extended is SoftLayer_Product_Item_Requirement with all maskable types.
type SoftLayer_Product_Item_Requirement_Extended struct {
	SoftLayer_Product_Item_Requirement

	// Item - no documentation
	Item *SoftLayer_Product_Item `json:"item"`

	// Product - no documentation
	Product *SoftLayer_Product_Item `json:"product"`
}

func (softlayer_product_item_requirement *SoftLayer_Product_Item_Requirement) String() string {
	return "SoftLayer_Product_Item_Requirement"
}
