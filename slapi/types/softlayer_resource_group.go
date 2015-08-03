package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group - <nil>
type SoftLayer_Resource_Group struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// RootResourceGroupId - <nil>
	RootResourceGroupId int `json:"rootResourceGroupId"`

	// Id - no documentation
	Id int `json:"id"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId"`

	// Description - no documentation
	Description string `json:"description"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_resource_group *SoftLayer_Resource_Group) String() string {
	return "SoftLayer_Resource_Group"
}

// SoftLayer_Resource_Group_Extended is SoftLayer_Resource_Group with all maskable types.
type SoftLayer_Resource_Group_Extended struct {
	SoftLayer_Resource_Group

	// HardwareMembers - no documentation
	HardwareMembers []*SoftLayer_Resource_Group_Member `json:"hardwareMembers"`

	// AncestorGroupCount - A count of a resource group's associated group ancestors.
	AncestorGroupCount uint64 `json:"ancestorGroupCount"`

	// VlanMemberCount - no documentation
	VlanMemberCount uint64 `json:"vlanMemberCount"`

	// Members - no documentation
	Members []*SoftLayer_Resource_Group_Member `json:"members"`

	// RootResourceGroup - no documentation
	RootResourceGroup *SoftLayer_Resource_Group `json:"rootResourceGroup"`

	// VlanMembers - no documentation
	VlanMembers []*SoftLayer_Resource_Group_Member `json:"vlanMembers"`

	// AttributeCount - A count of a resource group's associated attributes.
	AttributeCount uint64 `json:"attributeCount"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount"`

	// SubnetMemberCount - A count of a resource group's associated subnet members.
	SubnetMemberCount uint64 `json:"subnetMemberCount"`

	// Template - no documentation
	Template *SoftLayer_Resource_Group_Template `json:"template"`

	// AncestorGroups - no documentation
	AncestorGroups []*SoftLayer_Resource_Group `json:"ancestorGroups"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Attribute `json:"attributes"`

	// SubnetMembers - no documentation
	SubnetMembers []*SoftLayer_Resource_Group_Member `json:"subnetMembers"`

	// HardwareMemberCount - A count of a resource group's associated hardware members.
	HardwareMemberCount uint64 `json:"hardwareMemberCount"`
}

func (softlayer_resource_group *SoftLayer_Resource_Group_Extended) String() string {
	return "SoftLayer_Resource_Group"
}
