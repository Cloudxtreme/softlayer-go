package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Network_Subnet - <nil>
type SoftLayer_Resource_Group_Member_Network_Subnet struct {

	// Status - no documentation
	Status string `json:"status,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_resource_group_member_network_subnet *SoftLayer_Resource_Group_Member_Network_Subnet) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Subnet"
}

// SoftLayer_Resource_Group_Member_Network_Subnet_Extended is SoftLayer_Resource_Group_Member_Network_Subnet with all maskable types.
type SoftLayer_Resource_Group_Member_Network_Subnet_Extended struct {
	SoftLayer_Resource_Group_Member_Network_Subnet

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`

	// Resource - A resource group member's associated network subnet.
	Resource *SoftLayer_Network_Subnet `json:"resource,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`
}

func (softlayer_resource_group_member_network_subnet *SoftLayer_Resource_Group_Member_Network_Subnet_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Subnet"
}
