package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Module_Context_Interface - <nil>
type SoftLayer_Network_Firewall_Module_Context_Interface struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// FirewallContextAccessControlListCount - no documentation
	FirewallContextAccessControlListCount uint64 `json:"firewallContextAccessControlListCount,omitempty"`

	// FirewallContextAccessControlLists - <nil>
	FirewallContextAccessControlLists []*SoftLayer_Network_Firewall_AccessControlList `json:"firewallContextAccessControlLists,omitempty"`

	// NetworkVlan - <nil>
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`
}

func (softlayer_network_firewall_module_context_interface *SoftLayer_Network_Firewall_Module_Context_Interface) String() string {
	return "SoftLayer_Network_Firewall_Module_Context_Interface"
}
