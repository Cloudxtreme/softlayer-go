package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Item_Category_ZeroFee_Count - The
// SoftLayer_Container_Product_Item_Category_ZeroFee_Count data type represents a count of zero fee
// billing/invoice items.
type SoftLayer_Container_Product_Item_Category_ZeroFee_Count struct {

	// CategoryName - no documentation
	CategoryName string `json:"categoryName,omitempty"`

	// Count - no documentation
	Count int `json:"count,omitempty"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode,omitempty"`

	// CategoryId - no documentation
	CategoryId int `json:"categoryId,omitempty"`
}

func (softlayer_container_product_item_category_zerofee_count *SoftLayer_Container_Product_Item_Category_ZeroFee_Count) String() string {
	return "SoftLayer_Container_Product_Item_Category_ZeroFee_Count"
}
