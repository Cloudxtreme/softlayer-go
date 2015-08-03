package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Member_Resource_Group - <nil>
type SoftLayer_Resource_Group_Member_Resource_Group struct {
}

// SoftLayer_Resource_Group_Member_Resource_Group_Extended is SoftLayer_Resource_Group_Member_Resource_Group with all maskable types.
type SoftLayer_Resource_Group_Member_Resource_Group_Extended struct {
	SoftLayer_Resource_Group_Member_Resource_Group

	// Resource - A resource group member's associated resource group.
	Resource *SoftLayer_Resource_Group `json:"resource"`
}

func (softlayer_resource_group_member_resource_group *SoftLayer_Resource_Group_Member_Resource_Group) String() string {
	return "SoftLayer_Resource_Group_Member_Resource_Group"
}
