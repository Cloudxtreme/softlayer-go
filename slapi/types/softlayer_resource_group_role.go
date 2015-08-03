package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Role - <nil>
type SoftLayer_Resource_Group_Role struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// MemberLinkCount - no documentation
	MemberLinkCount uint64 `json:"memberLinkCount"`

	// MemberLinks - no documentation
	MemberLinks []*SoftLayer_Resource_Group_Member_Role_Link `json:"memberLinks"`
}

func (softlayer_resource_group_role *SoftLayer_Resource_Group_Role) String() string {
	return "SoftLayer_Resource_Group_Role"
}
