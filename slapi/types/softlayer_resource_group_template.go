package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Template - <nil>
type SoftLayer_Resource_Group_Template struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`
}

func (softlayer_resource_group_template *SoftLayer_Resource_Group_Template) String() string {
	return "SoftLayer_Resource_Group_Template"
}

// SoftLayer_Resource_Group_Template_Extended is SoftLayer_Resource_Group_Template with all maskable types.
type SoftLayer_Resource_Group_Template_Extended struct {
	SoftLayer_Resource_Group_Template

	// Members - <nil>
	Members []*SoftLayer_Resource_Group_Template_Member `json:"members"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount"`

	// Children - <nil>
	Children []*SoftLayer_Resource_Group_Template `json:"children"`
}

func (softlayer_resource_group_template *SoftLayer_Resource_Group_Template_Extended) String() string {
	return "SoftLayer_Resource_Group_Template"
}
