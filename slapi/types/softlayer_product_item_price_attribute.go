package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Price_Attribute - <nil>
type SoftLayer_Product_Item_Price_Attribute struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ItemPriceAttributeTypeId - <nil>
	ItemPriceAttributeTypeId int `json:"itemPriceAttributeTypeId,omitempty"`

	// ItemPriceId - <nil>
	ItemPriceId int `json:"itemPriceId,omitempty"`

	// Value - <nil>
	Value string `json:"value,omitempty"`

	// ItemPrice - <nil>
	ItemPrice *SoftLayer_Product_Item_Price `json:"itemPrice,omitempty"`

	// ItemPriceAttributeType - <nil>
	ItemPriceAttributeType *SoftLayer_Product_Item_Price_Attribute_Type `json:"itemPriceAttributeType,omitempty"`
}

func (softlayer_product_item_price_attribute *SoftLayer_Product_Item_Price_Attribute) String() string {
	return "SoftLayer_Product_Item_Price_Attribute"
}
