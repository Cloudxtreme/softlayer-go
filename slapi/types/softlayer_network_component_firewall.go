package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Firewall - The SoftLayer_Network_Component_Firewall data type contains
// general information relating to a single SoftLayer network component firewall. This is the object
// which ties the running rules to a specific downstream server. Use the [[SoftLayer Network Firewall
// Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer Network
// Firewall Update Request]] service to submit a firewall update request.
type SoftLayer_Network_Component_Firewall struct {

	// NetworkComponentId - Unique ID for the network component of the switch interface that this network
	// component firewall is attached to.
	NetworkComponentId int `json:"networkComponentId"`

	// Id - no documentation
	Id int `json:"id"`

	// Status - no documentation
	Status string `json:"status"`

	// GuestNetworkComponentId - Unique ID for the network component of the switch interface that this
	// network component firewall is attached to.
	GuestNetworkComponentId int `json:"guestNetworkComponentId"`
}

func (softlayer_network_component_firewall *SoftLayer_Network_Component_Firewall) String() string {
	return "SoftLayer_Network_Component_Firewall"
}

// SoftLayer_Network_Component_Firewall_Extended is SoftLayer_Network_Component_Firewall with all maskable types.
type SoftLayer_Network_Component_Firewall_Extended struct {
	SoftLayer_Network_Component_Firewall

	// Rules - The currently running rule set of this network component firewall.
	Rules []*SoftLayer_Network_Component_Firewall_Rule `json:"rules"`

	// SubnetCount - A count of the additional subnets linked to this network component firewall.
	SubnetCount uint64 `json:"subnetCount"`

	// NetworkFirewallUpdateRequest - no documentation
	NetworkFirewallUpdateRequest []*SoftLayer_Network_Firewall_Update_Request `json:"networkFirewallUpdateRequest"`

	// NetworkFirewallUpdateRequestCount - A count of the update requests made for this firewall.
	NetworkFirewallUpdateRequestCount uint64 `json:"networkFirewallUpdateRequestCount"`

	// RuleCount - A count of the currently running rule set of this network component firewall.
	RuleCount uint64 `json:"ruleCount"`

	// Subnets - The additional subnets linked to this network component firewall.
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// ApplyServerRuleSubnetCount - A count of the additional subnets linked to this network component
	// firewall, that inherit rules from the host that the context slot is attached to.
	ApplyServerRuleSubnetCount uint64 `json:"applyServerRuleSubnetCount"`

	// ApplyServerRuleSubnets - The additional subnets linked to this network component firewall, that
	// inherit rules from the host that the context slot is attached to.
	ApplyServerRuleSubnets []*SoftLayer_Network_Subnet `json:"applyServerRuleSubnets"`

	// BillingItem - The billing item for a Hardware Firewall (Dedicated).
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// GuestNetworkComponent - The network component of the guest virtual server that this network
	// component firewall belongs to.
	GuestNetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"guestNetworkComponent"`

	// NetworkComponent - The network component of the switch interface that this network component
	// firewall belongs to.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`
}

func (softlayer_network_component_firewall *SoftLayer_Network_Component_Firewall_Extended) String() string {
	return "SoftLayer_Network_Component_Firewall"
}
