package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Template_Rule - The SoftLayer_Network_Component_Firewall_Rule type
// contains general information relating to a single SoftLayer firewall template rule. Use the
// [[SoftLayer Network Component Firewall]] service to view current rules. Use the [[SoftLayer Network
// Firewall Update Request]] service to submit a firewall update request.
type SoftLayer_Network_Firewall_Template_Rule struct {

	// Action - The action that this template rule is to take [permit or deny].
	Action string `json:"action"`

	// DestinationIpAddress - The destination IP address considered for determining rule application.
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationIpSubnetMask - The destination IP subnet mask considered for determining rule
	// application.
	DestinationIpSubnetMask string `json:"destinationIpSubnetMask"`

	// DestinationPortRangeEnd - The ending (upper end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeEnd int `json:"destinationPortRangeEnd"`

	// DestinationPortRangeStart - The starting (lower end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeStart int `json:"destinationPortRangeStart"`

	// FirewallTemplate - The firewall template that this rule is attached to.
	FirewallTemplate *SoftLayer_Network_Firewall_Template `json:"firewallTemplate"`

	// FirewallTemplateId - The unique identifier of the firewall template that a firewall template rule is
	// associated with.
	FirewallTemplateId int `json:"firewallTemplateId"`

	// Id - no documentation
	Id int `json:"id"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// OrderValue - The numeric value describing the order in which the rule set should be applied.
	OrderValue int `json:"orderValue"`

	// Protocol - The protocol considered for determining rule application.
	Protocol string `json:"protocol"`

	// SourceIpAddress - The source IP address considered for determining rule application.
	SourceIpAddress string `json:"sourceIpAddress"`

	// SourceIpSubnetMask - The source IP subnet mask considered for determining rule application.
	SourceIpSubnetMask string `json:"sourceIpSubnetMask"`
}

func (softlayer_network_firewall_template_rule *SoftLayer_Network_Firewall_Template_Rule) String() string {
	return "SoftLayer_Network_Firewall_Template_Rule"
}
