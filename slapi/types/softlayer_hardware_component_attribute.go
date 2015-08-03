package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Attribute - The SoftLayer_Hardware_Component_Attribute data type
// contains general information relating to a single hardware setting or attribute for a component
// model. For Example: A controller may be setup for many different configurations. A controller with a
// configuration of RAID-1 will have a single attribute for this setting.
type SoftLayer_Hardware_Component_Attribute struct {

	// HardwareComponentAttributeTypeId - A hardware component attribute's associated
	// [[SoftLayer_Hardware_Component_Attribute_Type|type]] Id.
	HardwareComponentAttributeTypeId int `json:"hardwareComponentAttributeTypeId"`

	// HardwareComponentId - A hardware component attribute's associated
	// [[SoftLayer_Hardware_Component|hardware component]] Id.
	HardwareComponentId int `json:"hardwareComponentId"`

	// Value - A hardware component attribute's value. A value can have many different values depending on
	// the attributes [[SoftLayer_Hardware_Component_Attribute_Type|type]].
	Value string `json:"value"`
}

func (softlayer_hardware_component_attribute *SoftLayer_Hardware_Component_Attribute) String() string {
	return "SoftLayer_Hardware_Component_Attribute"
}

// SoftLayer_Hardware_Component_Attribute_Extended is SoftLayer_Hardware_Component_Attribute with all maskable types.
type SoftLayer_Hardware_Component_Attribute_Extended struct {
	SoftLayer_Hardware_Component_Attribute

	// HardwareComponent - A hardware component attribute's associated
	// [[SoftLayer_Hardware_Component|Hardware Component]].
	HardwareComponent *SoftLayer_Hardware_Component `json:"hardwareComponent"`

	// HardwareComponentAttributeType - A hardware component attribute's associated
	// [[SoftLayer_Hardware_Component_Attribute_Type|type]].
	HardwareComponentAttributeType *SoftLayer_Hardware_Component_Attribute_Type `json:"hardwareComponentAttributeType"`
}

func (softlayer_hardware_component_attribute *SoftLayer_Hardware_Component_Attribute_Extended) String() string {
	return "SoftLayer_Hardware_Component_Attribute"
}
