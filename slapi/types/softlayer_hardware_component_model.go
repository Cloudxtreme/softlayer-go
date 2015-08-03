package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model - The SoftLayer_Hardware_Component_Model data type contains
// general information relating to a single SoftLayer component model. A component model represents a
// vendor specific representation of a hardware component. Every piece of hardware on a server will
// have a specific hardware component model.
type SoftLayer_Hardware_Component_Model struct {

	// Description - A colon delimited list of hardware component model attributes.
	Description string `json:"description"`

	// Name - no documentation
	Name string `json:"name"`

	// Version - The model number or model description of a hardware component model.
	Version string `json:"version"`

	// Manufacturer - no documentation
	Manufacturer string `json:"manufacturer"`

	// ArchitectureTypeId - <nil>
	ArchitectureTypeId string `json:"architectureTypeId"`

	// Id - A hardware component model's internal identifier number.
	Id int `json:"id"`

	// LongDescription - <nil>
	LongDescription string `json:"longDescription"`

	// Capacity - A component model's capacity. The capacity of a component model depends on the model
	// itself. For Example: Hard drives have a capacity that reflects the amount of data that hard drive
	// can store.
	Capacity float64 `json:"capacity"`

	// HardwareGenericComponentModelId - The internal identifier of the generic component model for a
	// component model.
	HardwareGenericComponentModelId int `json:"hardwareGenericComponentModelId"`
}

func (softlayer_hardware_component_model *SoftLayer_Hardware_Component_Model) String() string {
	return "SoftLayer_Hardware_Component_Model"
}

// SoftLayer_Hardware_Component_Model_Extended is SoftLayer_Hardware_Component_Model with all maskable types.
type SoftLayer_Hardware_Component_Model_Extended struct {
	SoftLayer_Hardware_Component_Model

	// CompatibleParentComponentModels - All the component models that a hardware component model is
	// compatible with.
	CompatibleParentComponentModels []*SoftLayer_Hardware_Component_Model `json:"compatibleParentComponentModels"`

	// HardwareComponents - A hardware component model's physical components in inventory.
	HardwareComponents []*SoftLayer_Hardware_Component `json:"hardwareComponents"`

	// RebootTime - no documentation
	RebootTime *SoftLayer_Hardware_Component_Motherboard_Reboot_Time `json:"rebootTime"`

	// ArchitectureType - <nil>
	ArchitectureType *SoftLayer_Hardware_Component_Model_Architecture_Type `json:"architectureType"`

	// Type - no documentation
	Type string `json:"type"`

	// CompatibleArrayTypeCount - no documentation
	CompatibleArrayTypeCount uint64 `json:"compatibleArrayTypeCount"`

	// CompatibleParentComponentModelCount - A count of all the component models that a hardware component
	// model is compatible with.
	CompatibleParentComponentModelCount uint64 `json:"compatibleParentComponentModelCount"`

	// CompatibleChildComponentModels - All the component models that are compatible with a hardware
	// component model.
	CompatibleChildComponentModels []*SoftLayer_Hardware_Component_Model `json:"compatibleChildComponentModels"`

	// InfinibandCompatibleAttribute - <nil>
	InfinibandCompatibleAttribute *SoftLayer_Hardware_Component_Model_Attribute `json:"infinibandCompatibleAttribute"`

	// IsInfinibandCompatible - <nil>
	IsInfinibandCompatible bool `json:"isInfinibandCompatible"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// CompatibleChildComponentModelCount - A count of all the component models that are compatible with a
	// hardware component model.
	CompatibleChildComponentModelCount uint64 `json:"compatibleChildComponentModelCount"`

	// HardwareGenericComponentModel - The non-vendor specific generic component model for a hardware
	// component model.
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel"`

	// ValidAttributeTypeCount - A count of the types of attributes that are allowed for a given hardware
	// component model.
	ValidAttributeTypeCount uint64 `json:"validAttributeTypeCount"`

	// CompatibleArrayTypes - <nil>
	CompatibleArrayTypes []*SoftLayer_Configuration_Storage_Group_Array_Type `json:"compatibleArrayTypes"`

	// ValidAttributeTypes - The types of attributes that are allowed for a given hardware component model.
	ValidAttributeTypes []*SoftLayer_Hardware_Component_Model_Attribute_Type `json:"validAttributeTypes"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Hardware_Component_Model_Attribute `json:"attributes"`
}

func (softlayer_hardware_component_model *SoftLayer_Hardware_Component_Model_Extended) String() string {
	return "SoftLayer_Hardware_Component_Model"
}
