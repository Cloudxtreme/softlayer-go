package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Item_Category_ZeroFee_Count - The
// SoftLayer_Container_Product_Item_Category_ZeroFee_Count data type represents a count of zero fee
// billing/invoice items.
type SoftLayer_Container_Product_Item_Category_ZeroFee_Count struct {

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode"`

	// CategoryId - no documentation
	CategoryId int `json:"categoryId"`

	// CategoryName - no documentation
	CategoryName string `json:"categoryName"`

	// Count - no documentation
	Count int `json:"count"`
}

func (softlayer_container_product_item_category_zerofee_count *SoftLayer_Container_Product_Item_Category_ZeroFee_Count) String() string {
	return "SoftLayer_Container_Product_Item_Category_ZeroFee_Count"
}
