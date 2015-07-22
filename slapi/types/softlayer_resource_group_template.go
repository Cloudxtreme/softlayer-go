package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Resource_Group_Template - <nil>
type SoftLayer_Resource_Group_Template struct {

	// Children - <nil>
	Children []*SoftLayer_Resource_Group_Template `json:"children"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount"`

	// Members - <nil>
	Members []*SoftLayer_Resource_Group_Template_Member `json:"members"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`
}

// GetAllObjects - <nil>
func (softlayer_resource_group_template *SoftLayer_Resource_Group_Template) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Resource_Group_Template, error) {
	var returnValue []*SoftLayer_Resource_Group_Template
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_resource_group_template *SoftLayer_Resource_Group_Template) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Resource_Group_Template, error) {
	var returnValue *SoftLayer_Resource_Group_Template
	return returnValue, nil
}
