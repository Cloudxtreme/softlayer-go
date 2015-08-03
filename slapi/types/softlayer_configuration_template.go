package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Configuration_Template - The SoftLayer_Configuration_Template data type contains general
// information of an arbitrary resource.
type SoftLayer_Configuration_Template struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - Internal identifier of a SoftLayer account that this configuration template belongs to
	AccountId int `json:"accountId"`

	// ConfigurationSectionCount - no documentation
	ConfigurationSectionCount uint64 `json:"configurationSectionCount"`

	// ConfigurationSections - <nil>
	ConfigurationSections []*SoftLayer_Configuration_Template_Section `json:"configurationSections"`

	// ConfigurationTemplateReference - <nil>
	ConfigurationTemplateReference []*SoftLayer_Monitoring_Agent_Configuration_Template_Group_Reference `json:"configurationTemplateReference"`

	// ConfigurationTemplateReferenceCount - no documentation
	ConfigurationTemplateReferenceCount uint64 `json:"configurationTemplateReferenceCount"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DefaultValueCount - no documentation
	DefaultValueCount uint64 `json:"defaultValueCount"`

	// DefaultValues - <nil>
	DefaultValues []*SoftLayer_Configuration_Template_Section_Definition_Value `json:"defaultValues"`

	// DefinitionCount - no documentation
	DefinitionCount uint64 `json:"definitionCount"`

	// Definitions - <nil>
	Definitions []*SoftLayer_Configuration_Template_Section_Definition `json:"definitions"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Item - <nil>
	Item *SoftLayer_Product_Item `json:"item"`

	// ItemId - Internal identifier of a product item that this configuration template is associated with
	ItemId int `json:"itemId"`

	// LinkedSectionReferences - <nil>
	LinkedSectionReferences *SoftLayer_Configuration_Template_Section_Reference `json:"linkedSectionReferences"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Parent - <nil>
	Parent *SoftLayer_Configuration_Template `json:"parent"`

	// ParentId - Internal identifier of the parent configuration template
	ParentId int `json:"parentId"`

	// User - <nil>
	User *SoftLayer_User_Customer `json:"user"`

	// UserRecordId - Internal identifier of a user that last modified this configuration template
	UserRecordId int `json:"userRecordId"`
}

func (softlayer_configuration_template *SoftLayer_Configuration_Template) String() string {
	return "SoftLayer_Configuration_Template"
}
