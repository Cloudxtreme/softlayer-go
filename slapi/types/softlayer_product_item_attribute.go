package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Attribute - <nil>
type SoftLayer_Product_Item_Attribute struct {

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Item_Attribute_Type `json:"attributeType"`

	// AttributeTypeKeyName - <nil>
	AttributeTypeKeyName string `json:"attributeTypeKeyName"`

	// Id - <nil>
	Id int `json:"id"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ItemAttributeTypeId - <nil>
	ItemAttributeTypeId int `json:"itemAttributeTypeId"`

	// ItemId - <nil>
	ItemId int `json:"itemId"`

	// Value - <nil>
	Value string `json:"value"`
}

func (softlayer_product_item_attribute *SoftLayer_Product_Item_Attribute) String() string {
	return "SoftLayer_Product_Item_Attribute"
}
