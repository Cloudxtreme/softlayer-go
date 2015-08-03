package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Type - The SoftLayer_Hardware_Component_Type data type provides details
// on the type of component requested
type SoftLayer_Hardware_Component_Type struct {

	// Type - no documentation
	Type string `json:"type"`

	// TypeParentId - no documentation
	TypeParentId int `json:"typeParentId"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`
}

// SoftLayer_Hardware_Component_Type_Extended is SoftLayer_Hardware_Component_Type with all maskable types.
type SoftLayer_Hardware_Component_Type_Extended struct {
	SoftLayer_Hardware_Component_Type

	// HardwareGenericComponentModelCount - A count of the generic component model description for this
	// component type object.
	HardwareGenericComponentModelCount uint64 `json:"hardwareGenericComponentModelCount"`

	// HardwareGenericComponentModels - The generic component model description for this component type
	// object.
	HardwareGenericComponentModels []*SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModels"`

	// TypeParent - The parent generic component model object for this generic component model object.
	TypeParent *SoftLayer_Hardware_Component_Type `json:"typeParent"`
}

func (softlayer_hardware_component_type *SoftLayer_Hardware_Component_Type) String() string {
	return "SoftLayer_Hardware_Component_Type"
}
