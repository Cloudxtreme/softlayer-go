package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Section_Type - The SoftLayer_Configuration_Template_Section_Type
// data type contains information of a configuration section type. Configuration can contain
// sub-sections.
type SoftLayer_Configuration_Template_Section_Type struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - Internal identifier of a configuration section type
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_configuration_template_section_type *SoftLayer_Configuration_Template_Section_Type) String() string {
	return "SoftLayer_Configuration_Template_Section_Type"
}
