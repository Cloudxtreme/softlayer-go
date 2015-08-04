package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Attribute_Type - The SoftLayer_Hardware_Component_Attribute_Type data
// type contains general information for the type of an attribute for a hardware component.
type SoftLayer_Hardware_Component_Attribute_Type struct {

	// Description - The description for the date that a hardware component attribute type's
	// [[SoftLayer_Hardware_Component_Attribute|Attribute]] contains.
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_hardware_component_attribute_type *SoftLayer_Hardware_Component_Attribute_Type) String() string {
	return "SoftLayer_Hardware_Component_Attribute_Type"
}
