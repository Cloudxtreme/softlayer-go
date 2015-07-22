package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Configuration_Template_Section_Definition_Value - SoftLayer_Configuration_Section_Value is
// used to set the value for a configuration definition
type SoftLayer_Configuration_Template_Section_Definition_Value struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Definition - <nil>
	Definition *SoftLayer_Configuration_Template_Section_Definition `json:"definition"`

	// DefinitionId - Internal identifier of a configuration definition that this configuration value if
	// defined by
	DefinitionId int `json:"definitionId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template"`

	// TemplateId - Internal identifier of a configuration template that this configuration value belongs
	// to
	TemplateId int `json:"templateId"`

	// Value - no documentation
	Value string `json:"value"`
}

// GetObject - <nil>
func (softlayer_configuration_template_section_definition_value *SoftLayer_Configuration_Template_Section_Definition_Value) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Configuration_Template_Section_Definition_Value, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Definition_Value
	return returnValue, nil
}
