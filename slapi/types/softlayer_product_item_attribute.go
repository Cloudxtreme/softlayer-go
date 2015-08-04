package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Attribute - <nil>
type SoftLayer_Product_Item_Attribute struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ItemAttributeTypeId - <nil>
	ItemAttributeTypeId int `json:"itemAttributeTypeId,omitempty"`

	// ItemId - <nil>
	ItemId int `json:"itemId,omitempty"`

	// Value - <nil>
	Value string `json:"value,omitempty"`

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Item_Attribute_Type `json:"attributeType,omitempty"`

	// AttributeTypeKeyName - <nil>
	AttributeTypeKeyName string `json:"attributeTypeKeyName,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`
}

func (softlayer_product_item_attribute *SoftLayer_Product_Item_Attribute) String() string {
	return "SoftLayer_Product_Item_Attribute"
}
