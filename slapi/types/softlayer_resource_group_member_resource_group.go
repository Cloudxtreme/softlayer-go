package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Resource_Group - <nil>
type SoftLayer_Resource_Group_Member_Resource_Group struct {

	// Status - no documentation
	Status string `json:"status,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_resource_group_member_resource_group *SoftLayer_Resource_Group_Member_Resource_Group) String() string {
	return "SoftLayer_Resource_Group_Member_Resource_Group"
}

// SoftLayer_Resource_Group_Member_Resource_Group_Extended is SoftLayer_Resource_Group_Member_Resource_Group with all maskable types.
type SoftLayer_Resource_Group_Member_Resource_Group_Extended struct {
	SoftLayer_Resource_Group_Member_Resource_Group

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// Resource - A resource group member's associated resource group.
	Resource *SoftLayer_Resource_Group `json:"resource,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`
}

func (softlayer_resource_group_member_resource_group *SoftLayer_Resource_Group_Member_Resource_Group_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Resource_Group"
}
