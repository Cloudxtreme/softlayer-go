package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Member_Network_Subnet - <nil>
type SoftLayer_Resource_Group_Member_Network_Subnet struct {
}

// SoftLayer_Resource_Group_Member_Network_Subnet_Extended is SoftLayer_Resource_Group_Member_Network_Subnet with all maskable types.
type SoftLayer_Resource_Group_Member_Network_Subnet_Extended struct {
	SoftLayer_Resource_Group_Member_Network_Subnet

	// Resource - A resource group member's associated network subnet.
	Resource *SoftLayer_Network_Subnet `json:"resource"`
}

func (softlayer_resource_group_member_network_subnet *SoftLayer_Resource_Group_Member_Network_Subnet) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Subnet"
}
