package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group - <nil>
type SoftLayer_Resource_Group struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// RootResourceGroupId - <nil>
	RootResourceGroupId int `json:"rootResourceGroupId,omitempty"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_resource_group *SoftLayer_Resource_Group) String() string {
	return "SoftLayer_Resource_Group"
}

// SoftLayer_Resource_Group_Extended is SoftLayer_Resource_Group with all maskable types.
type SoftLayer_Resource_Group_Extended struct {
	SoftLayer_Resource_Group

	// SubnetMemberCount - A count of a resource group's associated subnet members.
	SubnetMemberCount uint64 `json:"subnetMemberCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Attribute `json:"attributes,omitempty"`

	// HardwareMembers - no documentation
	HardwareMembers []*SoftLayer_Resource_Group_Member `json:"hardwareMembers,omitempty"`

	// Members - no documentation
	Members []*SoftLayer_Resource_Group_Member `json:"members,omitempty"`

	// VlanMembers - no documentation
	VlanMembers []*SoftLayer_Resource_Group_Member `json:"vlanMembers,omitempty"`

	// AncestorGroups - no documentation
	AncestorGroups []*SoftLayer_Resource_Group `json:"ancestorGroups,omitempty"`

	// AncestorGroupCount - A count of a resource group's associated group ancestors.
	AncestorGroupCount uint64 `json:"ancestorGroupCount,omitempty"`

	// HardwareMemberCount - A count of a resource group's associated hardware members.
	HardwareMemberCount uint64 `json:"hardwareMemberCount,omitempty"`

	// VlanMemberCount - no documentation
	VlanMemberCount uint64 `json:"vlanMemberCount,omitempty"`

	// RootResourceGroup - no documentation
	RootResourceGroup *SoftLayer_Resource_Group `json:"rootResourceGroup,omitempty"`

	// Template - no documentation
	Template *SoftLayer_Resource_Group_Template `json:"template,omitempty"`

	// AttributeCount - A count of a resource group's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount,omitempty"`

	// SubnetMembers - no documentation
	SubnetMembers []*SoftLayer_Resource_Group_Member `json:"subnetMembers,omitempty"`
}

func (softlayer_resource_group *SoftLayer_Resource_Group_Extended) String() string {
	return "SoftLayer_Resource_Group"
}
