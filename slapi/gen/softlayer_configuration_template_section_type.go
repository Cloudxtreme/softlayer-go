package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Configuration_Template_Section_Type - The SoftLayer_Configuration_Template_Section_Type
// data type contains information of a configuration section type. Configuration can contain
// sub-sections.
type SoftLayer_Configuration_Template_Section_Type struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - Internal identifier of a configuration section type
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_configuration_template_section_type *SoftLayer_Configuration_Template_Section_Type) GetObject() (*SoftLayer_Configuration_Template_Section_Type, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Type
	return returnValue, nil
}
