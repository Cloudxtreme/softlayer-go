package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Tax_Category - The SoftLayer_Product_Item_Tax_Category data type contains the
// tax categories that are associated with products.
type SoftLayer_Product_Item_Tax_Category struct {

	// StatusFlag - no documentation
	StatusFlag int `json:"statusFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ItemCount - no documentation
	ItemCount uint64 `json:"itemCount,omitempty"`

	// Items - <nil>
	Items []*SoftLayer_Product_Item `json:"items,omitempty"`
}

func (softlayer_product_item_tax_category *SoftLayer_Product_Item_Tax_Category) String() string {
	return "SoftLayer_Product_Item_Tax_Category"
}
