package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Role - <nil>
type SoftLayer_Resource_Group_Role struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_resource_group_role *SoftLayer_Resource_Group_Role) String() string {
	return "SoftLayer_Resource_Group_Role"
}

// SoftLayer_Resource_Group_Role_Extended is SoftLayer_Resource_Group_Role with all maskable types.
type SoftLayer_Resource_Group_Role_Extended struct {
	SoftLayer_Resource_Group_Role

	// MemberLinkCount - no documentation
	MemberLinkCount uint64 `json:"memberLinkCount"`

	// MemberLinks - no documentation
	MemberLinks []*SoftLayer_Resource_Group_Member_Role_Link `json:"memberLinks"`
}

func (softlayer_resource_group_role *SoftLayer_Resource_Group_Role_Extended) String() string {
	return "SoftLayer_Resource_Group_Role"
}
