package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section - The SoftLayer_Configuration_Template_Section data type
// contains information of a configuration section. Configuration can contain sub-sections.
type SoftLayer_Configuration_Template_Section struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions"`

	// Description - no documentation
	Description string `json:"description"`

	// DisallowedDeletionFlag - <nil>
	DisallowedDeletionFlag bool `json:"disallowedDeletionFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// LinkedTemplate - <nil>
	LinkedTemplate *SoftLayer_Configuration_Template `json:"linkedTemplate"`

	// LinkedTemplateId - Internal identifier of a sub configuration template that this section points to.
	// Use this property if you wish to create a reference to a sub configuration template when creating a
	// linked section.
	LinkedTemplateId string `json:"linkedTemplateId"`

	// LinkedTemplateReference - <nil>
	LinkedTemplateReference *SoftLayer_Configuration_Template_Section_Reference `json:"linkedTemplateReference"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// ParentId - Internal identifier of the parent configuration section
	ParentId int `json:"parentId"`

	// ProfileCount - no documentation
	ProfileCount uint64 `json:"profileCount"`

	// Profiles - <nil>
	Profiles []*SoftLayer_Configuration_Template_Section_Profile `json:"profiles"`

	// SectionType - <nil>
	SectionType *SoftLayer_Configuration_Template_Section_Type `json:"sectionType"`

	// SectionTypeName - <nil>
	SectionTypeName string `json:"sectionTypeName"`

	// Sort - no documentation
	Sort int `json:"sort"`

	// SubSectionCount - no documentation
	SubSectionCount uint64 `json:"subSectionCount"`

	// SubSections - <nil>
	SubSections []*SoftLayer_Configuration_Template_Section `json:"subSections"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template"`

	// TemplateId - Internal identifier of a configuration template that this section belongs to
	TemplateId string `json:"templateId"`

	// TypeId - Internal identifier of the configuration section type
	TypeId int `json:"typeId"`
}

// GetObject - <nil>
func (softlayer_configuration_template_section *SoftLayer_Configuration_Template_Section) GetObject() (*SoftLayer_Configuration_Template_Section, error) {
	var returnValue *SoftLayer_Configuration_Template_Section
	return returnValue, nil
}

// HasSubSections - no documentation
func (softlayer_configuration_template_section *SoftLayer_Configuration_Template_Section) HasSubSections() (bool, error) {
	var returnValue bool
	return returnValue, nil
}
