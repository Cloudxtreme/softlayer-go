package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member - <nil>
type SoftLayer_Resource_Group_Member struct {

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group"`

	// Id - no documentation
	Id int `json:"id"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles"`

	// Status - no documentation
	Status string `json:"status"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type"`
}

func (softlayer_resource_group_member *SoftLayer_Resource_Group_Member) String() string {
	return "SoftLayer_Resource_Group_Member"
}
