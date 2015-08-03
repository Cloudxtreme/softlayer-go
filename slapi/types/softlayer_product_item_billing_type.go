package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Billing_Type - The SoftLayer_Product_Item_Billing_Type data type models
// special billing types for non-monthly billed items in the SoftLayer product catalog.
type SoftLayer_Product_Item_Billing_Type struct {

	// Name - A keyword describing a SoftLayer product item billing type.
	Name string `json:"name"`
}

func (softlayer_product_item_billing_type *SoftLayer_Product_Item_Billing_Type) String() string {
	return "SoftLayer_Product_Item_Billing_Type"
}
