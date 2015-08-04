package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template - The SoftLayer_Configuration_Template data type contains general
// information of an arbitrary resource.
type SoftLayer_Configuration_Template struct {

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// ItemId - Internal identifier of a product item that this configuration template is associated with
	ItemId int `json:"itemId,omitempty"`

	// ParentId - Internal identifier of the parent configuration template
	ParentId int `json:"parentId,omitempty"`

	// UserRecordId - Internal identifier of a user that last modified this configuration template
	UserRecordId int `json:"userRecordId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_configuration_template *SoftLayer_Configuration_Template) String() string {
	return "SoftLayer_Configuration_Template"
}

// SoftLayer_Configuration_Template_Extended is SoftLayer_Configuration_Template with all maskable types.
type SoftLayer_Configuration_Template_Extended struct {
	SoftLayer_Configuration_Template

	// LinkedSectionReferences - <nil>
	LinkedSectionReferences *SoftLayer_Configuration_Template_Section_Reference `json:"linkedSectionReferences,omitempty"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ConfigurationSections - <nil>
	ConfigurationSections []*SoftLayer_Configuration_Template_Section `json:"configurationSections,omitempty"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions,omitempty"`

	// ConfigurationSectionCount - no documentation
	ConfigurationSectionCount uint64 `json:"configurationSectionCount,omitempty"`

	// ConfigurationTemplateReference - <nil>
	ConfigurationTemplateReference []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference,omitempty"`

	// DefaultValues - <nil>
	DefaultValues []*SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValues,omitempty"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount,omitempty"`

	// DefaultValueCount - no documentation
	DefaultValueCount uint64 `json:"defaultValueCount,omitempty"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount,omitempty"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item,omitempty"`

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template `json:"parent,omitempty"`
}

func (softlayer_configuration_template *SoftLayer_Configuration_Template_Extended) String() string {
	return "SoftLayer_Configuration_Template"
}
