package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Network_Storage - <nil>
type SoftLayer_Resource_Group_Member_Network_Storage struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`
}

func (softlayer_resource_group_member_network_storage *SoftLayer_Resource_Group_Member_Network_Storage) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Storage"
}

// SoftLayer_Resource_Group_Member_Network_Storage_Extended is SoftLayer_Resource_Group_Member_Network_Storage with all maskable types.
type SoftLayer_Resource_Group_Member_Network_Storage_Extended struct {
	SoftLayer_Resource_Group_Member_Network_Storage

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// Resource - A resource group member's associated network storage.
	Resource *SoftLayer_Network_Storage `json:"resource,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`
}

func (softlayer_resource_group_member_network_storage *SoftLayer_Resource_Group_Member_Network_Storage_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Storage"
}
