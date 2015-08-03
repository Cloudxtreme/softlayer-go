package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Reference - The
// SoftLayer_Configuration_Template_Section_Reference data type contains information of a configuration
// section and its associated configuration template.
type SoftLayer_Configuration_Template_Section_Reference struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - Internal identifier of a configuration section reference.
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// SectionId - no documentation
	SectionId int `json:"sectionId"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId"`
}

// SoftLayer_Configuration_Template_Section_Reference_Extended is SoftLayer_Configuration_Template_Section_Reference with all maskable types.
type SoftLayer_Configuration_Template_Section_Reference_Extended struct {
	SoftLayer_Configuration_Template_Section_Reference

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template"`
}

func (softlayer_configuration_template_section_reference *SoftLayer_Configuration_Template_Section_Reference) String() string {
	return "SoftLayer_Configuration_Template_Section_Reference"
}
