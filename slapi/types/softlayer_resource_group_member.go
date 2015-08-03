package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member - <nil>
type SoftLayer_Resource_Group_Member struct {

	// Status - no documentation
	Status string `json:"status"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`
}

// SoftLayer_Resource_Group_Member_Extended is SoftLayer_Resource_Group_Member with all maskable types.
type SoftLayer_Resource_Group_Member_Extended struct {
	SoftLayer_Resource_Group_Member

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount"`
}

func (softlayer_resource_group_member *SoftLayer_Resource_Group_Member) String() string {
	return "SoftLayer_Resource_Group_Member"
}
