package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price_Attribute - <nil>
type SoftLayer_Product_Item_Price_Attribute struct {

	// Id - <nil>
	Id int `json:"id"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice"`

	// ItemPriceAttributeType - <nil>
	ItemPriceAttributeType *SoftLayer_Product_Item_Price_Attribute_Type `json:"itemPriceAttributeType"`

	// ItemPriceAttributeTypeId - <nil>
	ItemPriceAttributeTypeId int `json:"itemPriceAttributeTypeId"`

	// ItemPriceId - <nil>
	ItemPriceId int `json:"itemPriceId"`

	// Value - <nil>
	Value string `json:"value"`
}

func (softlayer_product_item_price_attribute *SoftLayer_Product_Item_Price_Attribute) String() string {
	return "SoftLayer_Product_Item_Price_Attribute"
}
