package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Definition_Value - SoftLayer_Configuration_Section_Value is
// used to set the value for a configuration definition
type SoftLayer_Configuration_Template_Section_Definition_Value struct {

	// DefinitionId - Internal identifier of a configuration definition that this configuration value if
	// defined by
	DefinitionId int `json:"definitionId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// TemplateId - Internal identifier of a configuration template that this configuration value belongs
	// to
	TemplateId int `json:"templateId"`

	// Value - no documentation
	Value string `json:"value"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_configuration_template_section_definition_value *SoftLayer_Configuration_Template_Section_Definition_Value) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Value"
}

// SoftLayer_Configuration_Template_Section_Definition_Value_Extended is SoftLayer_Configuration_Template_Section_Definition_Value with all maskable types.
type SoftLayer_Configuration_Template_Section_Definition_Value_Extended struct {
	SoftLayer_Configuration_Template_Section_Definition_Value

	// Definition - <nil>
	Definition *SoftLayer_Configuration_Template_Section_Definition `json:"definition"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template"`
}

func (softlayer_configuration_template_section_definition_value *SoftLayer_Configuration_Template_Section_Definition_Value_Extended) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Value"
}
