package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Generic_Attribute - The
// SoftLayer_Hardware_Component_Model_Generic_Attribute data type contains information relating to a
// single SoftLayer generic component model. Generic component model attributes can hold any
// information to describe functionality of the model. For Example: The number of cores that a
// processor has.
type SoftLayer_Hardware_Component_Model_Generic_Attribute struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`
}

func (softlayer_hardware_component_model_generic_attribute *SoftLayer_Hardware_Component_Model_Generic_Attribute) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic_Attribute"
}

// SoftLayer_Hardware_Component_Model_Generic_Attribute_Extended is SoftLayer_Hardware_Component_Model_Generic_Attribute with all maskable types.
type SoftLayer_Hardware_Component_Model_Generic_Attribute_Extended struct {
	SoftLayer_Hardware_Component_Model_Generic_Attribute

	// HardwareGenericComponentModel - no documentation
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel,omitempty"`
}

func (softlayer_hardware_component_model_generic_attribute *SoftLayer_Hardware_Component_Model_Generic_Attribute_Extended) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic_Attribute"
}
