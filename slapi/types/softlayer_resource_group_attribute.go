package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Attribute - <nil>
type SoftLayer_Resource_Group_Attribute struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group"`

	// Id - no documentation
	Id int `json:"id"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Attribute_Type `json:"type"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_resource_group_attribute *SoftLayer_Resource_Group_Attribute) String() string {
	return "SoftLayer_Resource_Group_Attribute"
}
