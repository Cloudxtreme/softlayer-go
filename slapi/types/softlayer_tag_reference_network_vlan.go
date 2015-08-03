package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Vlan - <nil>
type SoftLayer_Tag_Reference_Network_Vlan struct {
}

// SoftLayer_Tag_Reference_Network_Vlan_Extended is SoftLayer_Tag_Reference_Network_Vlan with all maskable types.
type SoftLayer_Tag_Reference_Network_Vlan_Extended struct {
	SoftLayer_Tag_Reference_Network_Vlan

	// Resource - <nil>
	Resource *SoftLayer_Network_Vlan `json:"resource"`
}

func (softlayer_tag_reference_network_vlan *SoftLayer_Tag_Reference_Network_Vlan) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan"
}
