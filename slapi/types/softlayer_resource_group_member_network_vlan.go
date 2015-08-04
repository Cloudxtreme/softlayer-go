package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Network_Vlan - <nil>
type SoftLayer_Resource_Group_Member_Network_Vlan struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_resource_group_member_network_vlan *SoftLayer_Resource_Group_Member_Network_Vlan) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Vlan"
}

// SoftLayer_Resource_Group_Member_Network_Vlan_Extended is SoftLayer_Resource_Group_Member_Network_Vlan with all maskable types.
type SoftLayer_Resource_Group_Member_Network_Vlan_Extended struct {
	SoftLayer_Resource_Group_Member_Network_Vlan

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Network_Vlan `json:"resource,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`
}

func (softlayer_resource_group_member_network_vlan *SoftLayer_Resource_Group_Member_Network_Vlan_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Vlan"
}
