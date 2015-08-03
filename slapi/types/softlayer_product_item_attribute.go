package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Attribute - <nil>
type SoftLayer_Product_Item_Attribute struct {

	// ItemId - <nil>
	ItemId int `json:"itemId"`

	// Value - <nil>
	Value string `json:"value"`

	// Id - <nil>
	Id int `json:"id"`

	// ItemAttributeTypeId - <nil>
	ItemAttributeTypeId int `json:"itemAttributeTypeId"`
}

// SoftLayer_Product_Item_Attribute_Extended is SoftLayer_Product_Item_Attribute with all maskable types.
type SoftLayer_Product_Item_Attribute_Extended struct {
	SoftLayer_Product_Item_Attribute

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Item_Attribute_Type `json:"attributeType"`

	// AttributeTypeKeyName - <nil>
	AttributeTypeKeyName string `json:"attributeTypeKeyName"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`
}

func (softlayer_product_item_attribute *SoftLayer_Product_Item_Attribute) String() string {
	return "SoftLayer_Product_Item_Attribute"
}
