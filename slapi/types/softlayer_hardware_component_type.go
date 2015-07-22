package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Type - The SoftLayer_Hardware_Component_Type data type provides details
// on the type of component requested
type SoftLayer_Hardware_Component_Type struct {

	// HardwareGenericComponentModelCount - A count of the generic component model description for this
	// component type object.
	HardwareGenericComponentModelCount uint64 `json:"hardwareGenericComponentModelCount"`

	// HardwareGenericComponentModels - The generic component model description for this component type
	// object.
	HardwareGenericComponentModels []*SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModels"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Type - no documentation
	Type string `json:"type"`

	// TypeParent - The parent generic component model object for this generic component model object.
	TypeParent *SoftLayer_Hardware_Component_Type `json:"typeParent"`

	// TypeParentId - no documentation
	TypeParentId int `json:"typeParentId"`
}
