package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_AccessControlList - The SoftLayer_Network_Firewall_AccessControlList data
// type contains general information relating to a single SoftLayer firewall access to controll list.
// This is the object which ties the running rules to a specific context. Use the [[SoftLayer Network
// Firewall Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer
// Network Firewall Update Request]] service to submit a firewall update request.
type SoftLayer_Network_Firewall_AccessControlList struct {

	// Direction - <nil>
	Direction string `json:"direction,omitempty"`

	// FirewallContextInterfaceId - <nil>
	FirewallContextInterfaceId int `json:"firewallContextInterfaceId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_network_firewall_accesscontrollist *SoftLayer_Network_Firewall_AccessControlList) String() string {
	return "SoftLayer_Network_Firewall_AccessControlList"
}

// SoftLayer_Network_Firewall_AccessControlList_Extended is SoftLayer_Network_Firewall_AccessControlList with all maskable types.
type SoftLayer_Network_Firewall_AccessControlList_Extended struct {
	SoftLayer_Network_Firewall_AccessControlList

	// Rules - The currently running rule set of this context access control list firewall.
	Rules []*SoftLayer_Network_Vlan_Firewall_Rule `json:"rules,omitempty"`

	// NetworkFirewallUpdateRequestCount - A count of the update requests made for this firewall.
	NetworkFirewallUpdateRequestCount uint64 `json:"networkFirewallUpdateRequestCount,omitempty"`

	// RuleCount - A count of the currently running rule set of this context access control list firewall.
	RuleCount uint64 `json:"ruleCount,omitempty"`

	// NetworkFirewallUpdateRequests - no documentation
	NetworkFirewallUpdateRequests []*SoftLayer_Network_Firewall_Update_Request `json:"networkFirewallUpdateRequests,omitempty"`

	// NetworkVlan - <nil>
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`
}

func (softlayer_network_firewall_accesscontrollist *SoftLayer_Network_Firewall_AccessControlList_Extended) String() string {
	return "SoftLayer_Network_Firewall_AccessControlList"
}
