package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Attribute - The SoftLayer_Hardware_Component__Model_Attribute
// data type contains general information relating to a single hardware setting or attribute for a
// component model.
type SoftLayer_Hardware_Component_Model_Attribute struct {

	// HardwareComponentModelId - A hardware component model attribute's associated
	// [[SoftLayer_Hardware_Component_Model|hardware component model]] Id.
	HardwareComponentModelId int `json:"hardwareComponentModelId,omitempty"`

	// Value - A hardware component model attribute's value. A value can have many different values
	// depending on the attributes [[SoftLayer_Hardware_Component_Model_Attribute_Type|type]].
	Value string `json:"value,omitempty"`

	// AttributeTypeId - A hardware component model attribute's associated
	// [[SoftLayer_Hardware_Component_Model_Attribute_Type|type]] Id.
	AttributeTypeId int `json:"attributeTypeId,omitempty"`

	// HardwareComponentAttributeType - <nil>
	HardwareComponentAttributeType *SoftLayer_Hardware_Component_Model_Attribute_Type `json:"hardwareComponentAttributeType,omitempty"`

	// HardwareComponent - <nil>
	HardwareComponent *SoftLayer_Hardware_Component_Model `json:"hardwareComponent,omitempty"`
}

func (softlayer_hardware_component_model_attribute *SoftLayer_Hardware_Component_Model_Attribute) String() string {
	return "SoftLayer_Hardware_Component_Model_Attribute"
}
