package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Attribute - The SoftLayer_Hardware_Attribute type contains general information
// for a hardware attribute. Hardware attributes can be assigned to specific hardware objects to
// describe relatively arbitrary information.
type SoftLayer_Hardware_Attribute struct {

	// HardwareAttributeTypeId - The unique identifier of a hardware attribute's type.
	HardwareAttributeTypeId int `json:"hardwareAttributeTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_Hardware_Attribute_Extended is SoftLayer_Hardware_Attribute with all maskable types.
type SoftLayer_Hardware_Attribute_Extended struct {
	SoftLayer_Hardware_Attribute

	// HardwareAttributeType - The type of hardware attribute that this represents.
	HardwareAttributeType *SoftLayer_Hardware_Attribute_Type `json:"hardwareAttributeType"`
}

func (softlayer_hardware_attribute *SoftLayer_Hardware_Attribute) String() string {
	return "SoftLayer_Hardware_Attribute"
}
