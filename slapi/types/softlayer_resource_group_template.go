package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Template - <nil>
type SoftLayer_Resource_Group_Template struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Children - <nil>
	Children []*SoftLayer_Resource_Group_Template `json:"children,omitempty"`

	// Members - <nil>
	Members []*SoftLayer_Resource_Group_Template_Member `json:"members,omitempty"`

	// Package - <nil>
	Package *SoftLayer_Product_Package `json:"package,omitempty"`

	// ChildrenCount - no documentation
	ChildrenCount uint64 `json:"childrenCount,omitempty"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount,omitempty"`
}

func (softlayer_resource_group_template *SoftLayer_Resource_Group_Template) String() string {
	return "SoftLayer_Resource_Group_Template"
}
