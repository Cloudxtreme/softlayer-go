package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Module_Context_Interface - <nil>
type SoftLayer_Network_Firewall_Module_Context_Interface struct {

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_network_firewall_module_context_interface *SoftLayer_Network_Firewall_Module_Context_Interface) String() string {
	return "SoftLayer_Network_Firewall_Module_Context_Interface"
}

// SoftLayer_Network_Firewall_Module_Context_Interface_Extended is SoftLayer_Network_Firewall_Module_Context_Interface with all maskable types.
type SoftLayer_Network_Firewall_Module_Context_Interface_Extended struct {
	SoftLayer_Network_Firewall_Module_Context_Interface

	// FirewallContextAccessControlLists - <nil>
	FirewallContextAccessControlLists []*SoftLayer_Network_Firewall_AccessControlList `json:"firewallContextAccessControlLists"`

	// NetworkVlan - <nil>
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// FirewallContextAccessControlListCount - no documentation
	FirewallContextAccessControlListCount uint64 `json:"firewallContextAccessControlListCount"`
}

func (softlayer_network_firewall_module_context_interface *SoftLayer_Network_Firewall_Module_Context_Interface_Extended) String() string {
	return "SoftLayer_Network_Firewall_Module_Context_Interface"
}
