package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Update_Request_Rule - The SoftLayer_Network_Firewall_Update_Request_Rule
// type contains information relating to a SoftLayer network firewall update request rule. This rule is
// a member of a [[SoftLayer Network Firewall Update Request]]. Use the [[SoftLayer Network Component
// Firewall]] service to view current rules. Use the [[SoftLayer Network Firewall Template]] service to
// pull SoftLayer recommended rule set templates.
type SoftLayer_Network_Firewall_Update_Request_Rule struct {

	// SourceIpSubnetMask - The source IP subnet mask considered for determining rule application.
	SourceIpSubnetMask string `json:"sourceIpSubnetMask"`

	// DestinationPortRangeEnd - The ending (upper end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeEnd int `json:"destinationPortRangeEnd"`

	// OrderValue - The numeric value describing the order in which the rule should be applied.
	OrderValue int `json:"orderValue"`

	// SourceIpAddress - The source IP address considered for determining rule application.
	SourceIpAddress string `json:"sourceIpAddress"`

	// SourceIpCidr - The is used for determining rule application. This value will
	SourceIpCidr int `json:"sourceIpCidr"`

	// Version - Whether this rule is an IPv4 rule or an IPv6 rule. If
	Version int `json:"version"`

	// DestinationIpCidr - The is used for determining rule application. This value will
	DestinationIpCidr int `json:"destinationIpCidr"`

	// FirewallUpdateRequestId - The unique identifier of the firewall update request that a firewall
	// update request rule is associated with.
	FirewallUpdateRequestId int `json:"firewallUpdateRequestId"`

	// Notes - The notes field for the firewall update request rule.
	Notes string `json:"notes"`

	// Protocol - The protocol considered for determining rule application.
	Protocol string `json:"protocol"`

	// Action - The action that this update request rule is to take [permit or deny].
	Action string `json:"action"`

	// DestinationIpAddress - The destination IP address considered for determining rule application.
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationIpSubnetMask - The destination IP subnet mask considered for determining rule
	// application.
	DestinationIpSubnetMask string `json:"destinationIpSubnetMask"`

	// DestinationPortRangeStart - The starting (lower end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeStart int `json:"destinationPortRangeStart"`

	// Id - A Firewall update request rule's internal identifier.
	Id int `json:"id"`
}

// SoftLayer_Network_Firewall_Update_Request_Rule_Extended is SoftLayer_Network_Firewall_Update_Request_Rule with all maskable types.
type SoftLayer_Network_Firewall_Update_Request_Rule_Extended struct {
	SoftLayer_Network_Firewall_Update_Request_Rule

	// FirewallUpdateRequest - no documentation
	FirewallUpdateRequest *SoftLayer_Network_Firewall_Update_Request `json:"firewallUpdateRequest"`
}

func (softlayer_network_firewall_update_request_rule *SoftLayer_Network_Firewall_Update_Request_Rule) String() string {
	return "SoftLayer_Network_Firewall_Update_Request_Rule"
}
