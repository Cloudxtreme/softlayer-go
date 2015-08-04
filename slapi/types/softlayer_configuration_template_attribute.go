package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Attribute - Configuration template attribute class contains
// supplementary information for a configuration template.
type SoftLayer_Configuration_Template_Attribute struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// ConfigurationTemplate - <nil>
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate,omitempty"`
}

func (softlayer_configuration_template_attribute *SoftLayer_Configuration_Template_Attribute) String() string {
	return "SoftLayer_Configuration_Template_Attribute"
}
