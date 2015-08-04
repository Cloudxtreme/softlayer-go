package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Section_Attribute - Configuration section attribute class contains
// supplementary information for a configuration section.
type SoftLayer_Configuration_Template_Section_Attribute struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// ConfigurationSection - <nil>
	ConfigurationSection *SoftLayer_Configuration_Template_Section `json:"configurationSection,omitempty"`
}

func (softlayer_configuration_template_section_attribute *SoftLayer_Configuration_Template_Section_Attribute) String() string {
	return "SoftLayer_Configuration_Template_Section_Attribute"
}
