package sl

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

	// Member - A resource group member attribute's resource group member.
	Member *SoftLayer_Resource_Group_Member `json:"member"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Member_Attribute_Type `json:"type"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_resource_group_member_attribute *SoftLayer_Resource_Group_Member_Attribute) String() string {
	return "SoftLayer_Resource_Group_Member_Attribute"
}
