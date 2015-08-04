package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Preset_Attribute_Type - SoftLayer_Product_Package_Preset_Attribute_Type
// models the type of attribute that can be assigned to a package preset.
type SoftLayer_Product_Package_Preset_Attribute_Type struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - A brief description of a package preset attribute type.
	Description string `json:"description,omitempty"`

	// Id - A package preset attribute type's internal identifier.
	Id int `json:"id,omitempty"`

	// KeyName - A package preset attribute type's key name. This is typically a shorter version of an
	// attribute type's name.
	KeyName string `json:"keyName,omitempty"`
}

func (softlayer_product_package_preset_attribute_type *SoftLayer_Product_Package_Preset_Attribute_Type) String() string {
	return "SoftLayer_Product_Package_Preset_Attribute_Type"
}
