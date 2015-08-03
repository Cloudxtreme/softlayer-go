package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section - The SoftLayer_Configuration_Template_Section data type
// contains information of a configuration section. Configuration can contain sub-sections.
type SoftLayer_Configuration_Template_Section struct {

	// Id - no documentation
	Id int `json:"id"`

	// TypeId - Internal identifier of the configuration section type
	TypeId int `json:"typeId"`

	// Name - no documentation
	Name string `json:"name"`

	// ParentId - Internal identifier of the parent configuration section
	ParentId int `json:"parentId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// TemplateId - Internal identifier of a configuration template that this section belongs to
	TemplateId string `json:"templateId"`

	// LinkedTemplateId - Internal identifier of a sub configuration template that this section points to.
	// Use this property if you wish to create a reference to a sub configuration template when creating a
	// linked section.
	LinkedTemplateId string `json:"linkedTemplateId"`

	// Sort - no documentation
	Sort int `json:"sort"`
}

// SoftLayer_Configuration_Template_Section_Extended is SoftLayer_Configuration_Template_Section with all maskable types.
type SoftLayer_Configuration_Template_Section_Extended struct {
	SoftLayer_Configuration_Template_Section

	// DisallowedDeletionFlag - <nil>
	DisallowedDeletionFlag bool `json:"disallowedDeletionFlag"`

	// LinkedTemplateReference - <nil>
	LinkedTemplateReference *SoftLayer_Configuration_Template_Section_Reference `json:"linkedTemplateReference"`

	// LinkedTemplate - <nil>
	LinkedTemplate *SoftLayer_Configuration_Template `json:"linkedTemplate"`

	// Profiles - <nil>
	Profiles []*SoftLayer_Configuration_Template_Section_Profile `json:"profiles"`

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount"`

	// ProfileCount - no documentation
	ProfileCount uint64 `json:"profileCount"`

	// SubSectionCount - no documentation
	SubSectionCount uint64 `json:"subSectionCount"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions"`

	// SectionType - <nil>
	SectionType *SoftLayer_Configuration_Template_Section_Type `json:"sectionType"`

	// SectionTypeName - <nil>
	SectionTypeName string `json:"sectionTypeName"`

	// SubSections - <nil>
	SubSections []*SoftLayer_Configuration_Template_Section `json:"subSections"`
}

func (softlayer_configuration_template_section *SoftLayer_Configuration_Template_Section) String() string {
	return "SoftLayer_Configuration_Template_Section"
}
