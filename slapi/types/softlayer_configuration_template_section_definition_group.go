package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template_Section_Definition_Group `json:"parent"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`
}

// GetAllGroups - Get all configuration definition group objects. ''getAllGroups'' returns an array of
// SoftLayer_Configuration_Template_Section_Definition_Group objects upon success.
func (softlayer_configuration_template_section_definition_group *SoftLayer_Configuration_Template_Section_Definition_Group) GetAllGroups(ctx *slapi.RequestContext) ([]*SoftLayer_Configuration_Template_Section_Definition_Group, error) {
	var returnValue []*SoftLayer_Configuration_Template_Section_Definition_Group
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_configuration_template_section_definition_group *SoftLayer_Configuration_Template_Section_Definition_Group) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Configuration_Template_Section_Definition_Group, error) {
	var returnValue *SoftLayer_Configuration_Template_Section_Definition_Group
	return returnValue, nil
}
