package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Firewall_AccessControlList - The SoftLayer_Network_Firewall_AccessControlList data
// type contains general information relating to a single SoftLayer firewall access to controll list.
// This is the object which ties the running rules to a specific context. Use the [[SoftLayer Network
// Firewall Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer
// Network Firewall Update Request]] service to submit a firewall update request.
type SoftLayer_Network_Firewall_AccessControlList struct {

	// Direction - <nil>
	Direction string `json:"direction"`

	// FirewallContextInterfaceId - <nil>
	FirewallContextInterfaceId int `json:"firewallContextInterfaceId"`

	// Id - <nil>
	Id int `json:"id"`

	// NetworkFirewallUpdateRequestCount - A count of the update requests made for this firewall.
	NetworkFirewallUpdateRequestCount uint64 `json:"networkFirewallUpdateRequestCount"`

	// NetworkFirewallUpdateRequests - no documentation
	NetworkFirewallUpdateRequests []*SoftLayer_Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests"`

	// NetworkVlan - <nil>
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// RuleCount - A count of the currently running rule set of this context access control list firewall.
	RuleCount uint64 `json:"ruleCount"`

	// Rules - The currently running rule set of this context access control list firewall.
	Rules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"rules"`
}

func (softlayer_network_firewall_accesscontrollist *SoftLayer_Network_Firewall_AccessControlList) String() string {
	return "SoftLayer_Network_Firewall_AccessControlList"
}

// GetObject - getObject returns a SoftLayer_Network_Firewall_AccessControlList object. You can only
// get objects for servers attached to your account that have a network firewall enabled.
func (softlayer_network_firewall_accesscontrollist *SoftLayer_Network_Firewall_AccessControlList) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Firewall_AccessControlList, error) {
	var returnValue *SoftLayer_Network_Firewall_AccessControlList
	return returnValue, nil
}
