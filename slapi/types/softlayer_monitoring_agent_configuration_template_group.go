package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Monitoring_Agent_Configuration_Template_Group - The
// SoftLayer_Monitoring_Agent_Configuration_Template_Group class is consisted of configuration
// templates for agents in a monitoring package.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group struct {

	// Description - Description of a monitoring agent configuration group
	Description string `json:"description,omitempty"`

	// Id - Internal identifier of a monitoring agent configuration group
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ItemId - Internal identifier of a configuration template type
	ItemId int `json:"itemId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Template_Group"
}

// SoftLayer_Monitoring_Agent_Configuration_Template_Group_Extended is SoftLayer_Monitoring_Agent_Configuration_Template_Group with all maskable types.
type SoftLayer_Monitoring_Agent_Configuration_Template_Group_Extended struct {
	SoftLayer_Monitoring_Agent_Configuration_Template_Group

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount,omitempty"`

	// ConfigurationTemplateCount - no documentation
	ConfigurationTemplateCount uint64 `json:"configurationTemplateCount,omitempty"`

	// ConfigurationTemplateReferences - <nil>
	ConfigurationTemplateReferences []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReferences,omitempty"`

	// ConfigurationTemplates - <nil>
	ConfigurationTemplates []*SoftLayer_Configuration_Template `json:"configurationTemplates,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`
}

func (softlayer_monitoring_agent_configuration_template_group *SoftLayer_Monitoring_Agent_Configuration_Template_Group_Extended) String() string {
	return "SoftLayer_Monitoring_Agent_Configuration_Template_Group"
}
