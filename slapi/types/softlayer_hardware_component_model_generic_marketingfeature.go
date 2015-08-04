package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_Model_Generic_MarketingFeature - The
// SoftLayer_Hardware_Component_Model_Generic_MarketingFeature data type contains general information
// relating to all the advertising features of a single SoftLayer hardware generic component model.
type SoftLayer_Hardware_Component_Model_Generic_MarketingFeature struct {

	// Features - no documentation
	Features string `json:"features,omitempty"`

	// Price - no documentation
	Price string `json:"price,omitempty"`
}

func (softlayer_hardware_component_model_generic_marketingfeature *SoftLayer_Hardware_Component_Model_Generic_MarketingFeature) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic_MarketingFeature"
}

// SoftLayer_Hardware_Component_Model_Generic_MarketingFeature_Extended is SoftLayer_Hardware_Component_Model_Generic_MarketingFeature with all maskable types.
type SoftLayer_Hardware_Component_Model_Generic_MarketingFeature_Extended struct {
	SoftLayer_Hardware_Component_Model_Generic_MarketingFeature

	// HardwareGenericComponentModel - The generic component model for a list of advertising or marketing
	// features
	HardwareGenericComponentModel *SoftLayer_Hardware_Component_Model_Generic `json:"hardwareGenericComponentModel,omitempty"`
}

func (softlayer_hardware_component_model_generic_marketingfeature *SoftLayer_Hardware_Component_Model_Generic_MarketingFeature_Extended) String() string {
	return "SoftLayer_Hardware_Component_Model_Generic_MarketingFeature"
}
