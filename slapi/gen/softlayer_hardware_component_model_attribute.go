package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Attribute - The SoftLayer_Hardware_Component__Model_Attribute
// data type contains general information relating to a single hardware setting or attribute for a
// component model.
type SoftLayer_Hardware_Component_Model_Attribute struct {

	// AttributeTypeId - A hardware component model attribute's associated
	// [[SoftLayer_Hardware_Component_Model_Attribute_Type|type]] Id.
	AttributeTypeId int `json:"attributeTypeId"`

	// HardwareComponent - <nil>
	HardwareComponent *SoftLayer_Hardware_Component_Model `json:"hardwareComponent"`

	// HardwareComponentAttributeType - <nil>
	HardwareComponentAttributeType *SoftLayer_Hardware_Component_Model_Attribute_Type `json:"hardwareComponentAttributeType"`

	// HardwareComponentModelId - A hardware component model attribute's associated
	// [[SoftLayer_Hardware_Component_Model|hardware component model]] Id.
	HardwareComponentModelId int `json:"hardwareComponentModelId"`

	// Value - A hardware component model attribute's value. A value can have many different values
	// depending on the attributes [[SoftLayer_Hardware_Component_Model_Attribute_Type|type]].
	Value string `json:"value"`
}
