package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Monitoring_Agent_Configuration_Template_Group - The
// SoftLayer_Monitoring_Agent_Configuration_Template_Group class is consisted of configuration
// templates for agents in a monitoring package.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId"`

	// ConfigurationTemplateCount - no documentation
	ConfigurationTemplateCount uint64 `json:"configurationTemplateCount"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount"`

	// ConfigurationTemplateReferences - <nil>
	ConfigurationTemplateReferences []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReferences"`

	// ConfigurationTemplates - <nil>
	ConfigurationTemplates []*SoftLayer_Configuration_Template `json:"configurationTemplates"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - Description of a monitoring agent configuration group
	Description string `json:"description"`

	// Id - Internal identifier of a monitoring agent configuration group
	Id int `json:"id"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ItemId - Internal identifier of a configuration template type
	ItemId int `json:"itemId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`
}

// CreateObject - This method creates a SoftLayer_Monitoring_Agent_Configuration_Template_Group using
// the values provided in the template object. The template objects accountId will be overridden to use
// the active user's accountId as it shows on their associated SoftLayer_User_Customer object.
func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Monitoring_Agent_Configuration_Template_Group) (*SoftLayer_Monitoring_Agent_Configuration_Template_Group, error) {
	var returnValue *SoftLayer_Monitoring_Agent_Configuration_Template_Group
	return returnValue, nil
}

// DeleteObject - no documentation
func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method edits an existing SoftLayer_Monitoring_Agent_Configuration_Template_Group
// using the values passed in the $object parameter. The $object parameter should use the same
// structure as a SoftLayer_Monitoring_Agent_Configuration_Template_Group object.
func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Monitoring_Agent_Configuration_Template_Group) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - <nil>
func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Monitoring_Agent_Configuration_Template_Group, error) {
	var returnValue []*SoftLayer_Monitoring_Agent_Configuration_Template_Group
	return returnValue, nil
}

// GetConfigurationGroups - This method retrieves an array of
// SoftLayer_Monitoring_Agent_Configuration_Template_Group objects that are available to the active
// user's account. The packageId parameter is not currently used.
func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) GetConfigurationGroups(commonOptions *slapi.CommonOptions, packageId int) ([]*SoftLayer_Monitoring_Agent_Configuration_Template_Group, error) {
	var returnValue []*SoftLayer_Monitoring_Agent_Configuration_Template_Group
	return returnValue, nil
}

// GetObject - This method retrieves a monitoring agent configuration template group whose identifier
// corresponds to the value provided in the initialization parameter passed to the
// SoftLayer_Monitoring_Agent_Configuration_Template_Group service.
func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Monitoring_Agent_Configuration_Template_Group, error) {
	var returnValue *SoftLayer_Monitoring_Agent_Configuration_Template_Group
	return returnValue, nil
}
