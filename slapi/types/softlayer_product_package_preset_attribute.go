package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset_Attribute - Package preset attributes contain supplementary
// information for a package preset.
type SoftLayer_Product_Package_Preset_Attribute struct {

	// AttributeType - <nil>
	AttributeType *SoftLayer_Product_Package_Preset_Attribute_Type `json:"attributeType"`

	// AttributeTypeId - The internal identifier of the type of attribute that a pacakge preset attribute
	// belongs to.
	AttributeTypeId int `json:"attributeTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// Preset - <nil>
	Preset *SoftLayer_Product_Package_Preset `json:"preset"`

	// PresetId - The internal identifier of the package preset an attribute belongs to.
	PresetId int `json:"presetId"`

	// Value - no documentation
	Value string `json:"value"`
}
