package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template - The SoftLayer_Configuration_Template data type contains general
// information of an arbitrary resource.
type SoftLayer_Configuration_Template struct {

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// UserRecordId - Internal identifier of a user that last modified this configuration template
	UserRecordId int `json:"userRecordId"`

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId"`

	// Name - no documentation
	Name string `json:"name"`

	// ParentId - Internal identifier of the parent configuration template
	ParentId int `json:"parentId"`

	// ItemId - Internal identifier of a product item that this configuration template is associated with
	ItemId int `json:"itemId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - no documentation
	Description string `json:"description"`
}

func (softlayer_configuration_template *SoftLayer_Configuration_Template) String() string {
	return "SoftLayer_Configuration_Template"
}

// SoftLayer_Configuration_Template_Extended is SoftLayer_Configuration_Template with all maskable types.
type SoftLayer_Configuration_Template_Extended struct {
	SoftLayer_Configuration_Template

	// DefaultValues - <nil>
	DefaultValues []*SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValues"`

	// ConfigurationSectionCount - no documentation
	ConfigurationSectionCount uint64 `json:"configurationSectionCount"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// LinkedSectionReferences - <nil>
	LinkedSectionReferences *SoftLayer_Configuration_Template_Section_Reference `json:"linkedSectionReferences"`

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template `json:"parent"`

	// DefaultValueCount - no documentation
	DefaultValueCount uint64 `json:"defaultValueCount"`

	// ConfigurationTemplateReference - <nil>
	ConfigurationTemplateReference []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ConfigurationSections - <nil>
	ConfigurationSections []*SoftLayer_Configuration_Template_Section `json:"configurationSections"`
}

func (softlayer_configuration_template *SoftLayer_Configuration_Template_Extended) String() string {
	return "SoftLayer_Configuration_Template"
}
