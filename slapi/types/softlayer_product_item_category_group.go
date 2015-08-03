package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category_Group - The SoftLayer_Product_Item_Category_Group data type contains
// general category group information.
type SoftLayer_Product_Item_Category_Group struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - The friendly, descriptive name of the category group as seen on the order forms and on
	// invoices.
	Name string `json:"name"`
}

func (softlayer_product_item_category_group *SoftLayer_Product_Item_Category_Group) String() string {
	return "SoftLayer_Product_Item_Category_Group"
}
