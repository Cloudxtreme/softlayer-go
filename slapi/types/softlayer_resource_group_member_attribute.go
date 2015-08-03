package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Member_Attribute - <nil>
type SoftLayer_Resource_Group_Member_Attribute struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Value - no documentation
	Value string `json:"value"`
}

// SoftLayer_Resource_Group_Member_Attribute_Extended is SoftLayer_Resource_Group_Member_Attribute with all maskable types.
type SoftLayer_Resource_Group_Member_Attribute_Extended struct {
	SoftLayer_Resource_Group_Member_Attribute

	// Member - A resource group member attribute's resource group member.
	Member *SoftLayer_Resource_Group_Member `json:"member"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Attribute_Type `json:"type"`
}

func (softlayer_resource_group_member_attribute *SoftLayer_Resource_Group_Member_Attribute) String() string {
	return "SoftLayer_Resource_Group_Member_Attribute"
}
