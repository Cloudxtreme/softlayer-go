package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Software_Component_Password - <nil>
type SoftLayer_Resource_Group_Member_Software_Component_Password struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`
}

func (softlayer_resource_group_member_software_component_password *SoftLayer_Resource_Group_Member_Software_Component_Password) String() string {
	return "SoftLayer_Resource_Group_Member_Software_Component_Password"
}

// SoftLayer_Resource_Group_Member_Software_Component_Password_Extended is SoftLayer_Resource_Group_Member_Software_Component_Password with all maskable types.
type SoftLayer_Resource_Group_Member_Software_Component_Password_Extended struct {
	SoftLayer_Resource_Group_Member_Software_Component_Password

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// Resource - A resource group member's associated software component password.
	Resource *SoftLayer_Software_Component_Password `json:"resource,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`
}

func (softlayer_resource_group_member_software_component_password *SoftLayer_Resource_Group_Member_Software_Component_Password_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Software_Component_Password"
}
