package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Configuration_Template_Section_Definition_Type -
// SoftLayer_Configuration_Template_Section_Definition_Type further defines the value of a
// configuration definition.
type SoftLayer_Configuration_Template_Section_Definition_Type struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_configuration_template_section_definition_type *SoftLayer_Configuration_Template_Section_Definition_Type) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Type"
}

// GetObject - <nil>
func (softlayer_configuration_template_section_definition_type *SoftLayer_Configuration_Template_Section_Definition_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Configuration_Template_Section_Definition_Type, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Definition_Type
	return returnValue, nil
}
