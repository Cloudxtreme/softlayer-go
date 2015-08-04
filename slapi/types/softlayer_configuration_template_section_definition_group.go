package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Definition_Group - Configuration definition group gives you
// details of the definition and allows extra functionality.
type SoftLayer_Configuration_Template_Section_Definition_Group struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder,omitempty"`
}

func (softlayer_configuration_template_section_definition_group *SoftLayer_Configuration_Template_Section_Definition_Group) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Group"
}

// SoftLayer_Configuration_Template_Section_Definition_Group_Extended is SoftLayer_Configuration_Template_Section_Definition_Group with all maskable types.
type SoftLayer_Configuration_Template_Section_Definition_Group_Extended struct {
	SoftLayer_Configuration_Template_Section_Definition_Group

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template_Section_Definition_Group `json:"parent,omitempty"`
}

func (softlayer_configuration_template_section_definition_group *SoftLayer_Configuration_Template_Section_Definition_Group_Extended) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Group"
}
