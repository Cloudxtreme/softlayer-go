package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Resource_Group_Attribute - <nil>
type SoftLayer_Resource_Group_Attribute struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`
}

func (softlayer_resource_group_attribute *SoftLayer_Resource_Group_Attribute) String() string {
	return "SoftLayer_Resource_Group_Attribute"
}

// SoftLayer_Resource_Group_Attribute_Extended is SoftLayer_Resource_Group_Attribute with all maskable types.
type SoftLayer_Resource_Group_Attribute_Extended struct {
	SoftLayer_Resource_Group_Attribute

	// Group - no documentation
	Group *SoftLayer_Resource_Group `json:"group,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Resource_Group_Attribute_Type `json:"type,omitempty"`
}

func (softlayer_resource_group_attribute *SoftLayer_Resource_Group_Attribute_Extended) String() string {
	return "SoftLayer_Resource_Group_Attribute"
}
