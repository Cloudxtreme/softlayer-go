package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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
}

func (softlayer_network_firewall_accesscontrollist *SoftLayer_Network_Firewall_AccessControlList) String() string {
	return "SoftLayer_Network_Firewall_AccessControlList"
}

// SoftLayer_Network_Firewall_AccessControlList_Extended is SoftLayer_Network_Firewall_AccessControlList with all maskable types.
type SoftLayer_Network_Firewall_AccessControlList_Extended struct {
	SoftLayer_Network_Firewall_AccessControlList

	// NetworkFirewallUpdateRequestCount - A count of the update requests made for this firewall.
	NetworkFirewallUpdateRequestCount uint64 `json:"networkFirewallUpdateRequestCount"`

	// RuleCount - A count of the currently running rule set of this context access control list firewall.
	RuleCount uint64 `json:"ruleCount"`

	// NetworkFirewallUpdateRequests - no documentation
	NetworkFirewallUpdateRequests []*SoftLayer_Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests"`

	// NetworkVlan - <nil>
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// Rules - The currently running rule set of this context access control list firewall.
	Rules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"rules"`
}

func (softlayer_network_firewall_accesscontrollist *SoftLayer_Network_Firewall_AccessControlList_Extended) String() string {
	return "SoftLayer_Network_Firewall_AccessControlList"
}
