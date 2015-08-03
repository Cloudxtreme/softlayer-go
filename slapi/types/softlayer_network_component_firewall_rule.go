package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Firewall_Rule - A SoftLayer_Network_Component_Firewall_Rule object type
// represents a currently running firewall rule and contains relative information. Use the [[SoftLayer
// Network Firewall Update Request]] service to submit a firewall update request. Use the [[SoftLayer
// Network Firewall Template]] service to pull SoftLayer recommended rule set templates.
type SoftLayer_Network_Component_Firewall_Rule struct {

	// Notes - no documentation
	Notes string `json:"notes"`

	// SourceIpAddress - The source IP address considered for determining rule application.
	SourceIpAddress string `json:"sourceIpAddress"`

	// SourceIpCidr - The is used for determining rule application. This value will
	SourceIpCidr int `json:"sourceIpCidr"`

	// SourceIpSubnetMask - The source IP subnet mask considered for determining rule application.
	SourceIpSubnetMask string `json:"sourceIpSubnetMask"`

	// Version - Whether this rule is an IPv4 rule or an IPv6 rule. If
	Version int `json:"version"`

	// Protocol - The protocol considered for determining rule application.
	Protocol string `json:"protocol"`

	// OrderValue - The numeric value describing the order in which the rule should be applied.
	OrderValue int `json:"orderValue"`

	// Id - no documentation
	Id int `json:"id"`

	// DestinationPortRangeEnd - The ending (upper end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeEnd int `json:"destinationPortRangeEnd"`

	// DestinationIpAddress - The destination IP address considered for determining rule application.
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationIpCidr - The is used for determining rule application. This value will
	DestinationIpCidr int `json:"destinationIpCidr"`

	// DestinationIpSubnetMask - The destination IP subnet mask considered for determining rule
	// application.
	DestinationIpSubnetMask string `json:"destinationIpSubnetMask"`

	// DestinationPortRangeStart - The starting (lower end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeStart int `json:"destinationPortRangeStart"`

	// Status - no documentation
	Status string `json:"status"`

	// Action - The action that the rule is to take [permit or deny].
	Action string `json:"action"`
}

func (softlayer_network_component_firewall_rule *SoftLayer_Network_Component_Firewall_Rule) String() string {
	return "SoftLayer_Network_Component_Firewall_Rule"
}

// SoftLayer_Network_Component_Firewall_Rule_Extended is SoftLayer_Network_Component_Firewall_Rule with all maskable types.
type SoftLayer_Network_Component_Firewall_Rule_Extended struct {
	SoftLayer_Network_Component_Firewall_Rule

	// NetworkComponentFirewall - The network component firewall that this rule belongs to.
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`
}

func (softlayer_network_component_firewall_rule *SoftLayer_Network_Component_Firewall_Rule_Extended) String() string {
	return "SoftLayer_Network_Component_Firewall_Rule"
}
