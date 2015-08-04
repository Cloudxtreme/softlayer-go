package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Section_Definition_Attribute - Configuration definition attribute
// class contains supplementary information for a configuration definition.
type SoftLayer_Configuration_Template_Section_Definition_Attribute struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// ConfigurationDefinition - <nil>
	ConfigurationDefinition *SoftLayer_Configuration_Template_Section_Definition `json:"configurationDefinition,omitempty"`

	// AttributeType - <nil>
	AttributeType *SoftLayer_Configuration_Template_Section_Definition_Attribute_Type `json:"attributeType,omitempty"`
}

func (softlayer_configuration_template_section_definition_attribute *SoftLayer_Configuration_Template_Section_Definition_Attribute) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Attribute"
}
