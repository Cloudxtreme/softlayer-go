package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Attribute_Type - The
// SoftLayer_Hardware_Component_Model_Attribute_Type data type contains general information for the
// type of an attribute for a hardware component model.
type SoftLayer_Hardware_Component_Model_Attribute_Type struct {

	// Description - The description for the data that a hardware component model type's
	// [[SoftLayer_Hardware_Component_Model_Attribute|Attribute]] contains.
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - A hardware component model attribute type's unique name.
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_hardware_component_model_attribute_type *SoftLayer_Hardware_Component_Model_Attribute_Type) String() string {
	return "SoftLayer_Hardware_Component_Model_Attribute_Type"
}

// SoftLayer_Hardware_Component_Model_Attribute_Type_Extended is SoftLayer_Hardware_Component_Model_Attribute_Type with all maskable types.
type SoftLayer_Hardware_Component_Model_Attribute_Type_Extended struct {
	SoftLayer_Hardware_Component_Model_Attribute_Type

	// ValidComponentTypes - <nil>
	ValidComponentTypes []*SoftLayer_Hardware_Component_Type `json:"validComponentTypes,omitempty"`

	// ValidComponentTypeCount - no documentation
	ValidComponentTypeCount uint64 `json:"validComponentTypeCount,omitempty"`
}

func (softlayer_hardware_component_model_attribute_type *SoftLayer_Hardware_Component_Model_Attribute_Type_Extended) String() string {
	return "SoftLayer_Hardware_Component_Model_Attribute_Type"
}
