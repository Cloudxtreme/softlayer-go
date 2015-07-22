package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Firewall - The SoftLayer_Network_Component_Firewall data type contains
// general information relating to a single SoftLayer network component firewall. This is the object
// which ties the running rules to a specific downstream server. Use the [[SoftLayer Network Firewall
// Template]] service to pull SoftLayer recommended rule set templates. Use the [[SoftLayer Network
// Firewall Update Request]] service to submit a firewall update request.
type SoftLayer_Network_Component_Firewall struct {

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

	// GuestNetworkComponentId - Unique ID for the network component of the switch interface that this
	// network component firewall is attached to.
	GuestNetworkComponentId int `json:"guestNetworkComponentId"`

	// Id - no documentation
	Id int `json:"id"`

	// NetworkComponent - The network component of the switch interface that this network component
	// firewall belongs to.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// NetworkComponentId - Unique ID for the network component of the switch interface that this network
	// component firewall is attached to.
	NetworkComponentId int `json:"networkComponentId"`

	// NetworkFirewallUpdateRequest - no documentation
	NetworkFirewallUpdateRequest []*SoftLayer_Network_Firewall_Update_Request `json:"networkFirewallUpdateRequest"`

	// NetworkFirewallUpdateRequestCount - A count of the update requests made for this firewall.
	NetworkFirewallUpdateRequestCount uint64 `json:"networkFirewallUpdateRequestCount"`

	// RuleCount - A count of the currently running rule set of this network component firewall.
	RuleCount uint64 `json:"ruleCount"`

	// Rules - The currently running rule set of this network component firewall.
	Rules []*SoftLayer_Network_Component_Firewall_Rule `json:"rules"`

	// Status - no documentation
	Status string `json:"status"`

	// SubnetCount - A count of the additional subnets linked to this network component firewall.
	SubnetCount uint64 `json:"subnetCount"`

	// Subnets - The additional subnets linked to this network component firewall.
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`
}

// GetObject - getObject returns a
// SoftLayer_Network_Firewall_Module_Context_Interface_AccessControlList_Network_Component object. You
// can only get objects for servers attached to your account that have a network firewall enabled.
func (softlayer_network_component_firewall *SoftLayer_Network_Component_Firewall) GetObject() (*SoftLayer_Network_Component_Firewall, error) {
	var returnValue *SoftLayer_Network_Component_Firewall
	return returnValue, nil
}
