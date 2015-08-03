package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Tax_Category - The SoftLayer_Product_Item_Tax_Category data type contains the
// tax categories that are associated with products.
type SoftLayer_Product_Item_Tax_Category struct {

	// StatusFlag - no documentation
	StatusFlag int `json:"statusFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

// SoftLayer_Product_Item_Tax_Category_Extended is SoftLayer_Product_Item_Tax_Category with all maskable types.
type SoftLayer_Product_Item_Tax_Category_Extended struct {
	SoftLayer_Product_Item_Tax_Category

	// ItemCount - no documentation
	ItemCount uint64 `json:"itemCount"`

	// Items - <nil>
	Items []*SoftLayer_Product_Item `json:"items"`
}

func (softlayer_product_item_tax_category *SoftLayer_Product_Item_Tax_Category) String() string {
	return "SoftLayer_Product_Item_Tax_Category"
}
