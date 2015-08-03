package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template_Section_Definition_Group - Configuration definition group gives you
// details of the definition and allows extra functionality.
type SoftLayer_Configuration_Template_Section_Definition_Group struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`
}

// SoftLayer_Configuration_Template_Section_Definition_Group_Extended is SoftLayer_Configuration_Template_Section_Definition_Group with all maskable types.
type SoftLayer_Configuration_Template_Section_Definition_Group_Extended struct {
	SoftLayer_Configuration_Template_Section_Definition_Group

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template_Section_Definition_Group `json:"parent"`
}

func (softlayer_configuration_template_section_definition_group *SoftLayer_Configuration_Template_Section_Definition_Group) String() string {
	return "SoftLayer_Configuration_Template_Section_Definition_Group"
}
