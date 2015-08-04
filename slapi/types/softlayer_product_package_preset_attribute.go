package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset_Attribute - Package preset attributes contain supplementary
// information for a package preset.
type SoftLayer_Product_Package_Preset_Attribute struct {

	// AttributeTypeId - The internal identifier of the type of attribute that a pacakge preset attribute
	// belongs to.
	AttributeTypeId int `json:"attributeTypeId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// PresetId - The internal identifier of the package preset an attribute belongs to.
	PresetId int `json:"presetId,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Package_Preset_Attribute_Type `json:"attributeType,omitempty"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset,omitempty"`
}

func (softlayer_product_package_preset_attribute *SoftLayer_Product_Package_Preset_Attribute) String() string {
	return "SoftLayer_Product_Package_Preset_Attribute"
}
