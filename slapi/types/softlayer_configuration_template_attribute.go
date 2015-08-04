package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Attribute - Configuration template attribute class contains
// supplementary information for a configuration template.
type SoftLayer_Configuration_Template_Attribute struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`
}

func (softlayer_configuration_template_attribute *SoftLayer_Configuration_Template_Attribute) String() string {
	return "SoftLayer_Configuration_Template_Attribute"
}

// SoftLayer_Configuration_Template_Attribute_Extended is SoftLayer_Configuration_Template_Attribute with all maskable types.
type SoftLayer_Configuration_Template_Attribute_Extended struct {
	SoftLayer_Configuration_Template_Attribute

	// ConfigurationTemplate - <nil>
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate,omitempty"`
}

func (softlayer_configuration_template_attribute *SoftLayer_Configuration_Template_Attribute_Extended) String() string {
	return "SoftLayer_Configuration_Template_Attribute"
}
