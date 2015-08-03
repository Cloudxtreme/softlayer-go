package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price_Attribute - <nil>
type SoftLayer_Product_Item_Price_Attribute struct {

	// Id - <nil>
	Id int `json:"id"`

	// ItemPriceAttributeTypeId - <nil>
	ItemPriceAttributeTypeId int `json:"itemPriceAttributeTypeId"`

	// ItemPriceId - <nil>
	ItemPriceId int `json:"itemPriceId"`

	// Value - <nil>
	Value string `json:"value"`
}

// SoftLayer_Product_Item_Price_Attribute_Extended is SoftLayer_Product_Item_Price_Attribute with all maskable types.
type SoftLayer_Product_Item_Price_Attribute_Extended struct {
	SoftLayer_Product_Item_Price_Attribute

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ItemPriceAttributeType - <nil>
	ItemPriceAttributeType *SoftLayer_Product_Item_Price_Attribute_Type `json:"itemPriceAttributeType"`
}

func (softlayer_product_item_price_attribute *SoftLayer_Product_Item_Price_Attribute) String() string {
	return "SoftLayer_Product_Item_Price_Attribute"
}
