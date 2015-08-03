package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category_Order_Option_Type - The
// SoftLayer_Product_Item_Category_Order_Option_Type data type contains options that can be applied to
// orders for prices.
type SoftLayer_Product_Item_Category_Order_Option_Type struct {

	// Id - no documentation
	Id int `json:"id"`

	// Keyname - A simple description for an item category order type.
	Keyname string `json:"keyname"`

	// Name - no documentation
	Name string `json:"name"`

	// Value - no documentation
	Value string `json:"value"`

	// Description - no documentation
	Description string `json:"description"`
}

func (softlayer_product_item_category_order_option_type *SoftLayer_Product_Item_Category_Order_Option_Type) String() string {
	return "SoftLayer_Product_Item_Category_Order_Option_Type"
}
