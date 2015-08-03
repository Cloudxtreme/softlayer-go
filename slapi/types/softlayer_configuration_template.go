package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template - The SoftLayer_Configuration_Template data type contains general
// information of an arbitrary resource.
type SoftLayer_Configuration_Template struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`

	// ItemId - Internal identifier of a product item that this configuration template is associated with
	ItemId int `json:"itemId"`

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`

	// ParentId - Internal identifier of the parent configuration template
	ParentId int `json:"parentId"`

	// UserRecordId - Internal identifier of a user that last modified this configuration template
	UserRecordId int `json:"userRecordId"`
}

// SoftLayer_Configuration_Template_Extended is SoftLayer_Configuration_Template with all maskable types.
type SoftLayer_Configuration_Template_Extended struct {
	SoftLayer_Configuration_Template

	// LinkedSectionReferences - <nil>
	LinkedSectionReferences *SoftLayer_Configuration_Template_Section_Reference `json:"linkedSectionReferences"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount"`

	// DefaultValues - <nil>
	DefaultValues []*SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValues"`

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template `json:"parent"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ConfigurationSections - <nil>
	ConfigurationSections []*SoftLayer_Configuration_Template_Section `json:"configurationSections"`

	// ConfigurationTemplateReference - <nil>
	ConfigurationTemplateReference []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions"`

	// DefaultValueCount - no documentation
	DefaultValueCount uint64 `json:"defaultValueCount"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ConfigurationSectionCount - no documentation
	ConfigurationSectionCount uint64 `json:"configurationSectionCount"`
}

func (softlayer_configuration_template *SoftLayer_Configuration_Template) String() string {
	return "SoftLayer_Configuration_Template"
}
