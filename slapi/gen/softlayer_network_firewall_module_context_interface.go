package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Module_Context_Interface - <nil>
type SoftLayer_Network_Firewall_Module_Context_Interface struct {

	// FirewallContextAccessControlListCount - no documentation
	FirewallContextAccessControlListCount uint64 `json:"firewallContextAccessControlListCount"`

	// FirewallContextAccessControlLists - <nil>
	FirewallContextAccessControlLists []*SoftLayer_Network_Firewall_AccessControlList `json:"firewallContextAccessControlLists"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`

	// NetworkVlan - <nil>
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`
}

// GetObject - <nil>
func (softlayer_network_firewall_module_context_interface *SoftLayer_Network_Firewall_Module_Context_Interface) GetObject() (*SoftLayer_Network_Firewall_Module_Context_Interface, error) {
	var returnValue *SoftLayer_Network_Firewall_Module_Context_Interface
	return returnValue, nil
}
