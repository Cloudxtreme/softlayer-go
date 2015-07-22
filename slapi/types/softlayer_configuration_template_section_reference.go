package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

	// Section - <nil>
	Section *SoftLayer_Configuration_Template_Section `json:"section"`

	// SectionId - no documentation
	SectionId int `json:"sectionId"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId"`
}

// GetObject - <nil>
func (softlayer_configuration_template_section_reference *SoftLayer_Configuration_Template_Section_Reference) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Configuration_Template_Section_Reference, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Reference
	return returnValue, nil
}
