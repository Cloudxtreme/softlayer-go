package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Attribute - Configuration template attribute class contains
// supplementary information for a configuration template.
type SoftLayer_Configuration_Template_Attribute struct {

	// ConfigurationTemplate - <nil>
	ConfigurationTemplate *SoftLayer_Configuration_Template `json:"configurationTemplate"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_configuration_template_attribute *SoftLayer_Configuration_Template_Attribute) String() string {
	return "SoftLayer_Configuration_Template_Attribute"
}
