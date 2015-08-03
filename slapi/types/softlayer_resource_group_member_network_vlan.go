package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Group_Member_Network_Vlan - <nil>
type SoftLayer_Resource_Group_Member_Network_Vlan struct {
}

func (softlayer_resource_group_member_network_vlan *SoftLayer_Resource_Group_Member_Network_Vlan) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Vlan"
}

// SoftLayer_Resource_Group_Member_Network_Vlan_Extended is SoftLayer_Resource_Group_Member_Network_Vlan with all maskable types.
type SoftLayer_Resource_Group_Member_Network_Vlan_Extended struct {
	SoftLayer_Resource_Group_Member_Network_Vlan

	// Resource - no documentation
	Resource *SoftLayer_Network_Vlan `json:"resource"`
}

func (softlayer_resource_group_member_network_vlan *SoftLayer_Resource_Group_Member_Network_Vlan_Extended) String() string {
	return "SoftLayer_Resource_Group_Member_Network_Vlan"
}
