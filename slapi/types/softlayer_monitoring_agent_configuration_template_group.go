package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Monitoring_Agent_Configuration_Template_Group - The
// SoftLayer_Monitoring_Agent_Configuration_Template_Group class is consisted of configuration
// templates for agents in a monitoring package.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group struct {

	// ItemId - Internal identifier of a configuration template type
	ItemId int `json:"itemId"`

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - Description of a monitoring agent configuration group
	Description string `json:"description"`

	// Id - Internal identifier of a monitoring agent configuration group
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`
}

// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Extended is SoftLayer_Monitoring_Agent_Configuration_Template_Group with all maskable types.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group_Extended struct {
	SoftLayer_Monitoring_Agent_Configuration_Template_Group

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ConfigurationTemplateCount - no documentation
	ConfigurationTemplateCount uint64 `json:"configurationTemplateCount"`

	// ConfigurationTemplateReferences - <nil>
	ConfigurationTemplateReferences []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReferences"`

	// ConfigurationTemplates - <nil>
	ConfigurationTemplates []*SoftLayer_Configuration_Template `json:"configurationTemplates"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount"`
}

func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Template_Group"
}
