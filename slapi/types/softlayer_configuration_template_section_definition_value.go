package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Definition_Value - SoftLayer_Configuration_Section_Value is
// used to set the value for a configuration definition
type SoftLayer_Configuration_Template_Section_Definition_Value struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DefinitionId - Internal identifier of a configuration definition that this configuration value if
	// defined by
	DefinitionId int `json:"definitionId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// TemplateId - Internal identifier of a configuration template that this configuration value belongs
	// to
	TemplateId int `json:"templateId,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// Definition - <nil>
	Definition *SoftLayer_Configuration_Template_Section_Definition `json:"definition,omitempty"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template,omitempty"`
}

func (softlayer_configuration_template_section_definition_value *SoftLayer_Configuration_Template_Section_Definition_Value) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Value"
}
