package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Generic_MarketingFeature - The
// SoftLayer_Hardware_Component_Model_Generic_MarketingFeature data type contains general information
// relating to all the advertising features of a single SoftLayer hardware generic component model.
type SoftLayer_Hardware_Component_Model_Generic_MarketingFeature struct {

	// Features - no documentation
	Features string `json:"features"`

	// HardwareGenericComponentModel - The generic component model for a list of advertising or marketing
	// features
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel"`

	// Price - no documentation
	Price string `json:"price"`
}

func (softlayer_hardware_component_model_generic_marketingfeature *SoftLayer_Hardware_Component_Model_Generic_MarketingFeature) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic_MarketingFeature"
}
