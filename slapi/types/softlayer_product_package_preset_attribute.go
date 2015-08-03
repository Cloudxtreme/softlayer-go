package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset_Attribute - Package preset attributes contain supplementary
// information for a package preset.
type SoftLayer_Product_Package_Preset_Attribute struct {

	// AttributeTypeId - The internal identifier of the type of attribute that a pacakge preset attribute
	// belongs to.
	AttributeTypeId int `json:"attributeTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// PresetId - The internal identifier of the package preset an attribute belongs to.
	PresetId int `json:"presetId"`

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_Product_Package_Preset_Attribute_Extended is SoftLayer_Product_Package_Preset_Attribute with all maskable types.
type SoftLayer_Product_Package_Preset_Attribute_Extended struct {
	SoftLayer_Product_Package_Preset_Attribute

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Package_Preset_Attribute_Type `json:"attributeType"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset"`
}

func (softlayer_product_package_preset_attribute *SoftLayer_Product_Package_Preset_Attribute) String() string {
	return "SoftLayer_Product_Package_Preset_Attribute"
}
