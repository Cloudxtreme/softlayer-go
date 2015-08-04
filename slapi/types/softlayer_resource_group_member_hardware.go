package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Hardware - <nil>
type SoftLayer_Resource_Group_Member_Hardware struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Status - no documentation
	Status string `json:"status,omitempty"`
}

func (softlayer_resource_group_member_hardware *SoftLayer_Resource_Group_Member_Hardware) String() string {
	return "SoftLayer_Resource_Group_Member_Hardware"
}

// SoftLayer_Resource_Group_Member_Hardware_Extended is SoftLayer_Resource_Group_Member_Hardware with all maskable types.
type SoftLayer_Resource_Group_Member_Hardware_Extended struct {
	SoftLayer_Resource_Group_Member_Hardware

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Type `json:"type,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware `json:"resource,omitempty"`

	// ServerArbiterOnly - A resource group hardware member's associated server arbiter-only state.
	ServerArbiterOnly *SoftLayer_Resource_Group_Member_Attribute `json:"serverArbiterOnly,omitempty"`

	// ServerPriority - A resource group hardware member's associated server priority.
	ServerPriority *SoftLayer_Resource_Group_Member_Attribute `json:"serverPriority,omitempty"`

	// ServerTags - A resource group hardware member's associated server tags (in format).
	ServerTags *SoftLayer_Resource_Group_Member_Attribute `json:"serverTags,omitempty"`

	// Roles - no documentation
	Roles []*SoftLayer_Resource_Group_Role `json:"roles,omitempty"`

	// RoleCount - A count of a resource group member's associated roles.
	RoleCount uint64 `json:"roleCount,omitempty"`

	// ServerHidden - A resource group hardware member's associated server hidden state.
	ServerHidden *SoftLayer_Resource_Group_Member_Attribute `json:"serverHidden,omitempty"`

	// ServerSlaveDelay - A resource group hardware member's associated server slave delay (in seconds).
	ServerSlaveDelay *SoftLayer_Resource_Group_Member_Attribute `json:"serverSlaveDelay,omitempty"`

	// ServerVotes - A resource group hardware member's associated server vote count.
	ServerVotes *SoftLayer_Resource_Group_Member_Attribute `json:"serverVotes,omitempty"`

	// DescendantMembers - A resource group member's associated member descendants.
	DescendantMembers []*SoftLayer_Resource_Group_Member `json:"descendantMembers,omitempty"`

	// AttributeCount - A count of a resource group member's associated attributes.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// DescendantMemberCount - A count of a resource group member's associated member descendants.
	DescendantMemberCount uint64 `json:"descendantMemberCount,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_Resource_Group_Member_Attribute `json:"attributes,omitempty"`
}

func (softlayer_resource_group_member_hardware *SoftLayer_Resource_Group_Member_Hardware_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Hardware"
}
