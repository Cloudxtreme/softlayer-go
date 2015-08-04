package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Attribute - <nil>
type SoftLayer_Product_Package_Attribute struct {

	// Value - <nil>
	Value string `json:"value,omitempty"`
}

func (softlayer_product_package_attribute *SoftLayer_Product_Package_Attribute) String() string {
	return "SoftLayer_Product_Package_Attribute"
}

// SoftLayer_Product_Package_Attribute_Extended is SoftLayer_Product_Package_Attribute with all maskable types.
type SoftLayer_Product_Package_Attribute_Extended struct {
	SoftLayer_Product_Package_Attribute

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Package_Attribute_Type `json:"attributeType,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`
}

func (softlayer_product_package_attribute *SoftLayer_Product_Package_Attribute_Extended) String() string {
	return "SoftLayer_Product_Package_Attribute"
}
