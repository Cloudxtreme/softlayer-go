package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Member_Network_Storage - <nil>
type SoftLayer_Resource_Group_Member_Network_Storage struct {
}

// SoftLayer_Resource_Group_Member_Network_Storage_Extended is SoftLayer_Resource_Group_Member_Network_Storage with all maskable types.
type SoftLayer_Resource_Group_Member_Network_Storage_Extended struct {
	SoftLayer_Resource_Group_Member_Network_Storage

	// Resource - A resource group member's associated network storage.
	Resource *SoftLayer_Network_Storage `json:"resource"`
}

func (softlayer_resource_group_member_network_storage *SoftLayer_Resource_Group_Member_Network_Storage) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Storage"
}
