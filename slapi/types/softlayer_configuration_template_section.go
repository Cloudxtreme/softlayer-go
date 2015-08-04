package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section - The SoftLayer_Configuration_Template_Section data type
// contains information of a configuration section. Configuration can contain sub-sections.
type SoftLayer_Configuration_Template_Section struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// ParentId - Internal identifier of the parent configuration section
	ParentId int `json:"parentId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// TemplateId - Internal identifier of a configuration template that this section belongs to
	TemplateId string `json:"templateId,omitempty"`

	// TypeId - Internal identifier of the configuration section type
	TypeId int `json:"typeId,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Sort - no documentation
	Sort int `json:"sort,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// LinkedTemplateId - Internal identifier of a sub configuration template that this section points to.
	// Use this property if you wish to create a reference to a sub configuration template when creating a
	// linked section.
	LinkedTemplateId string `json:"linkedTemplateId,omitempty"`
}

func (softlayer_configuration_template_section *SoftLayer_Configuration_Template_Section) String() string {
	return "SoftLayer_Configuration_Template_Section"
}

// SoftLayer_Configuration_Template_Section_Extended is SoftLayer_Configuration_Template_Section with all maskable types.
type SoftLayer_Configuration_Template_Section_Extended struct {
	SoftLayer_Configuration_Template_Section

	// Template - <nil>
	Template *SoftLayer_Configuration_Template `json:"template,omitempty"`

	// LinkedTemplateReference - <nil>
	LinkedTemplateReference *SoftLayer_Configuration_Template_Section_Reference `json:"linkedTemplateReference,omitempty"`

	// Profiles - <nil>
	Profiles []*SoftLayer_Configuration_Template_Section_Profile `json:"profiles,omitempty"`

	// SectionType - <nil>
	SectionType *SoftLayer_Configuration_Template_Section_Type `json:"sectionType,omitempty"`

	// SectionTypeName - <nil>
	SectionTypeName string `json:"sectionTypeName,omitempty"`

	// SubSections - <nil>
	SubSections []*SoftLayer_Configuration_Template_Section `json:"subSections,omitempty"`

	// SubSectionCount - no documentation
	SubSectionCount uint64 `json:"subSectionCount,omitempty"`

	// ProfileCount - no documentation
	ProfileCount uint64 `json:"profileCount,omitempty"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions,omitempty"`

	// LinkedTemplate - <nil>
	LinkedTemplate *SoftLayer_Configuration_Template `json:"linkedTemplate,omitempty"`

	// DisallowedDeletionFlag - <nil>
	DisallowedDeletionFlag bool `json:"disallowedDeletionFlag,omitempty"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount,omitempty"`
}

func (softlayer_configuration_template_section *SoftLayer_Configuration_Template_Section_Extended) String() string {
	return "SoftLayer_Configuration_Template_Section"
}
