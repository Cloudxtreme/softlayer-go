package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

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

func (softlayer_network_firewall_module_context_interface *SoftLayer_Network_Firewall_Module_Context_Interface) String() string {
	return "SoftLayer_Network_Firewall_Module_Context_Interface"
}

// GetObject - <nil>
func (softlayer_network_firewall_module_context_interface *SoftLayer_Network_Firewall_Module_Context_Interface) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Firewall_Module_Context_Interface, error) {
	var returnValue *SoftLayer_Network_Firewall_Module_Context_Interface
	return returnValue, nil
}
