package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Attribute - The SoftLayer_Hardware_Attribute type contains general information
// for a hardware attribute. Hardware attributes can be assigned to specific hardware objects to
// describe relatively arbitrary information.
type SoftLayer_Hardware_Attribute struct {

	// HardwareAttributeType - The type of hardware attribute that this represents.
	HardwareAttributeType *SoftLayer_Hardware_Attribute_Type `json:"hardwareAttributeType"`

	// HardwareAttributeTypeId - The unique identifier of a hardware attribute's type.
	HardwareAttributeTypeId int `json:"hardwareAttributeTypeId"`

	// Id - no documentation
	Id int `json:"id"`

	// Value - no documentation
	Value string `json:"value"`
}