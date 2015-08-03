package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Network_Vlan_Firewall - <nil>
type SoftLayer_Tag_Reference_Network_Vlan_Firewall struct {
}

func (softlayer_tag_reference_network_vlan_firewall *SoftLayer_Tag_Reference_Network_Vlan_Firewall) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan_Firewall"
}

// SoftLayer_Tag_Reference_Network_Vlan_Firewall_Extended is SoftLayer_Tag_Reference_Network_Vlan_Firewall with all maskable types.
type SoftLayer_Tag_Reference_Network_Vlan_Firewall_Extended struct {
	SoftLayer_Tag_Reference_Network_Vlan_Firewall

	// Resource - <nil>
	Resource *SoftLayer_Network_Vlan_Firewall `json:"resource"`
}

func (softlayer_tag_reference_network_vlan_firewall *SoftLayer_Tag_Reference_Network_Vlan_Firewall_Extended) String() string {
	return "SoftLayer_Tag_Reference_Network_Vlan_Firewall"
}
