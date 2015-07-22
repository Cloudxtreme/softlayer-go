package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Generic_Attribute - The
// SoftLayer_Hardware_Component_Model_Generic_Attribute data type contains information relating to a
// single SoftLayer generic component model. Generic component model attributes can hold any
// information to describe functionality of the model. For Example: The number of cores that a
// processor has.
type SoftLayer_Hardware_Component_Model_Generic_Attribute struct {

	// HardwareGenericComponentModel - no documentation
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel"`

	// Value - no documentation
	Value string `json:"value"`
}