package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Reference - The
// SoftLayer_Configuration_Template_Section_Reference data type contains information of a configuration
// section and its associated configuration template.
type SoftLayer_Configuration_Template_Section_Reference struct {

	// Id - Internal identifier of a configuration section reference.
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// SectionId - no documentation
	SectionId int `json:"sectionId,omitempty"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template,omitempty"`

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section,omitempty"`
}

func (softlayer_configuration_template_section_reference *SoftLayer_Configuration_Template_Section_Reference) String() string {
	return "SoftLayer_Configuration_Template_Section_Reference"
}
