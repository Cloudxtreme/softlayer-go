package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Generic - The SoftLayer_Hardware_Component_Model_Generic data
// type contains general information relating to a single SoftLayer generic component model. A generic
// component model represents a non-vendor specific representation of a hardware component. Frequently
// SoftLayer utilizes components from various vendors in the servers they provision. For Example:
// Several different vendors produce 6GB DDR2 memory. The generic component model for the 6GB stick of
// RAM encompasses every instance of this component regardless of make and model.
type SoftLayer_Hardware_Component_Model_Generic struct {

	// Id - A generic component model's internal identification number.
	Id int `json:"id,omitempty"`

	// HardwareComponentTypeId - The internal identifier of the component type for a generic component
	// model.
	HardwareComponentTypeId int `json:"hardwareComponentTypeId,omitempty"`

	// Units - The unit of measurement for the capacity of a generic component model.
	Units string `json:"units,omitempty"`

	// UpgradePriority - A generic component model's upgrade priority. The upgrade priority indicates the
	// order a generic component model should be considered over other generic component models. A higher
	// number indicates that a generic component model receives a higher upgrade preference in comparison
	// to a generic component model with a lower priority number.
	UpgradePriority int `json:"upgradePriority,omitempty"`

	// Capacity - A generic component model's capacity. The capacity of a generic component model depends
	// on the model itself. For Example: Hard drives have a capacity that reflects the amount of data that
	// hard drive can store.
	Capacity float64 `json:"capacity,omitempty"`

	// Description - A brief description for a generic component model that typically defines it's
	// function.
	Description string `json:"description,omitempty"`
}

func (softlayer_hardware_component_model_generic *SoftLayer_Hardware_Component_Model_Generic) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic"
}

// SoftLayer_Hardware_Component_Model_Generic_Extended is SoftLayer_Hardware_Component_Model_Generic with all maskable types.
type SoftLayer_Hardware_Component_Model_Generic_Extended struct {
	SoftLayer_Hardware_Component_Model_Generic

	// HardwareComponentType - no documentation
	HardwareComponentType *SoftLayer_Hardware_Component_Type `json:"hardwareComponentType,omitempty"`

	// HardwareComponentModelCount - A count of a generic component model's hardware component model.
	HardwareComponentModelCount uint64 `json:"hardwareComponentModelCount,omitempty"`

	// HardwareComponentModels - A generic component model's hardware component model.
	HardwareComponentModels []*SoftLayer_Hardware_Component_Model `json:"hardwareComponentModels,omitempty"`

	// MarketingFeatures - A list of features that a generic component model can provide.
	MarketingFeatures *SoftLayer_Hardware_Component_Model_Generic_MarketingFeature `json:"marketingFeatures,omitempty"`
}

func (softlayer_hardware_component_model_generic *SoftLayer_Hardware_Component_Model_Generic_Extended) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic"
}
