package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod - <nil>
type SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_resource_group_member_cloudstack_version3_pod *SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod) String() string {
	return "SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod"
}

// SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod_Extended is SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod with all maskable types.
type SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod_Extended struct {
	SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod

	// Resource - no documentation
	Resource *SoftLayer_Resource_Group `json:"resource,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`
}

func (softlayer_resource_group_member_cloudstack_version3_pod *SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_CloudStack_Version3_Pod"
}
