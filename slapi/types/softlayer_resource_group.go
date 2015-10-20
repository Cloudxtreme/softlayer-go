package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group - <nil>
type SoftLayer_Resource_Group struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// RootResourceGroupId - <nil>
	RootResourceGroupId int `json:"rootResourceGroupId,omitempty"`

	// TemplateId - no documentation
	TemplateId int `json:"templateId,omitempty"`

	// AttributeCount - A count of a resource group's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount,omitempty"`

	// SubnetMemberCount - A count of a resource group's associated subnet members.
	SubnetMemberCount uint64 `json:"subnetMemberCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Attribute `json:"attributes,omitempty"`

	// RootResourceGroup - no documentation
	RootResourceGroup *SoftLayer_Resource_Group `json:"rootResourceGroup,omitempty"`

	// SubnetMembers - no documentation
	SubnetMembers []*SoftLayer_Resource_Group_Member `json:"subnetMembers,omitempty"`

	// Template - no documentation
	Template *SoftLayer_Resource_Group_Template `json:"template,omitempty"`

	// AncestorGroupCount - A count of a resource group's associated group ancestors.
	AncestorGroupCount uint64 `json:"ancestorGroupCount,omitempty"`

	// HardwareMembers - no documentation
	HardwareMembers []*SoftLayer_Resource_Group_Member `json:"hardwareMembers,omitempty"`

	// Members - no documentation
	Members []*SoftLayer_Resource_Group_Member `json:"members,omitempty"`

	// AncestorGroups - no documentation
	AncestorGroups []*SoftLayer_Resource_Group `json:"ancestorGroups,omitempty"`

	// VlanMembers - no documentation
	VlanMembers []*SoftLayer_Resource_Group_Member `json:"vlanMembers,omitempty"`

	// HardwareMemberCount - A count of a resource group's associated hardware members.
	HardwareMemberCount uint64 `json:"hardwareMemberCount,omitempty"`

	// VlanMemberCount - no documentation
	VlanMemberCount uint64 `json:"vlanMemberCount,omitempty"`
}

func (softlayer_resource_group *SoftLayer_Resource_Group) String() string {
	return "SoftLayer_Resource_Group"
}
