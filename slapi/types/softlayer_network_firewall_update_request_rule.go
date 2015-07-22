package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Firewall_Update_Request_Rule - The SoftLayer_Network_Firewall_Update_Request_Rule
// type contains information relating to a SoftLayer network firewall update request rule. This rule is
// a member of a [[SoftLayer Network Firewall Update Request]]. Use the [[SoftLayer Network Component
// Firewall]] service to view current rules. Use the [[SoftLayer Network Firewall Template]] service to
// pull SoftLayer recommended rule set templates.
type SoftLayer_Network_Firewall_Update_Request_Rule struct {

	// Action - The action that this update request rule is to take [permit or deny].
	Action string `json:"action"`

	// DestinationIpAddress - The destination IP address considered for determining rule application.
	DestinationIpAddress string `json:"destinationIpAddress"`

	// DestinationIpCidr - The is used for determining rule application. This value will
	DestinationIpCidr int `json:"destinationIpCidr"`

	// DestinationIpSubnetMask - The destination IP subnet mask considered for determining rule
	// application.
	DestinationIpSubnetMask string `json:"destinationIpSubnetMask"`

	// DestinationPortRangeEnd - The ending (upper end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeEnd int `json:"destinationPortRangeEnd"`

	// DestinationPortRangeStart - The starting (lower end of range) destination port considered for
	// determining rule application.
	DestinationPortRangeStart int `json:"destinationPortRangeStart"`

	// FirewallUpdateRequest - no documentation
	FirewallUpdateRequest *SoftLayer_Network_Firewall_Update_Request `json:"firewallUpdateRequest"`

	// FirewallUpdateRequestId - The unique identifier of the firewall update request that a firewall
	// update request rule is associated with.
	FirewallUpdateRequestId int `json:"firewallUpdateRequestId"`

	// Id - A Firewall update request rule's internal identifier.
	Id int `json:"id"`

	// Notes - The notes field for the firewall update request rule.
	Notes string `json:"notes"`

	// OrderValue - The numeric value describing the order in which the rule should be applied.
	OrderValue int `json:"orderValue"`

	// Protocol - The protocol considered for determining rule application.
	Protocol string `json:"protocol"`

	// SourceIpAddress - The source IP address considered for determining rule application.
	SourceIpAddress string `json:"sourceIpAddress"`

	// SourceIpCidr - The is used for determining rule application. This value will
	SourceIpCidr int `json:"sourceIpCidr"`

	// SourceIpSubnetMask - The source IP subnet mask considered for determining rule application.
	SourceIpSubnetMask string `json:"sourceIpSubnetMask"`

	// Version - Whether this rule is an IPv4 rule or an IPv6 rule. If
	Version int `json:"version"`
}

// CreateObject - Create a new firewall update request. The SoftLayer_Network_Firewall_Update_Request
// object passed to this function must have at least one rule. ''createObject'' returns a Boolean
// ''true'' on successful object creation or ''false'' if your firewall update request was unable to be
// created..
func (softlayer_network_firewall_update_request_rule *SoftLayer_Network_Firewall_Update_Request_Rule) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Firewall_Update_Request_Rule) (*SoftLayer_Network_Firewall_Update_Request_Rule, error) {
	var returnValue *SoftLayer_Network_Firewall_Update_Request_Rule
	return returnValue, nil
}

// GetObject - getObject returns a SoftLayer_Network_Firewall_Update_Request_Rule object. You can only
// get historical objects for servers attached to your account that have a network firewall enabled.
// createObject inserts a new SoftLayer_Network_Firewall_Update_Request_Rule object. Use the
// SoftLayer_Network_Firewall_Update_Request to create groups of rules for an update request.
func (softlayer_network_firewall_update_request_rule *SoftLayer_Network_Firewall_Update_Request_Rule) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Firewall_Update_Request_Rule, error) {
	var returnValue *SoftLayer_Network_Firewall_Update_Request_Rule
	return returnValue, nil
}

// ValidateRule - Validate the supplied firewall request rule against the object it will apply to. For
// IPv4 rules, pass in an instance of SoftLayer_Network_Firewall_Update_Request_Rule. for IPv6 rules,
// pass in an instance of SoftLayer_Network_Firewall_Update_Request_Rule_Version6. The ID of the
// applied to object can either be applyToComponentId (an ID of a SoftLayer_Network_Component_Firewall)
// or applyToAclId (an ID of a SoftLayer_Network_Firewall_Module_Context_Interface_AccessControlList).
// One, and only one, of applyToComponentId and applyToAclId can be specified. If validation is
// successful, nothing is returned. If validation is unsuccessful, an exception is thrown explaining
// the nature of the validation error.
func (softlayer_network_firewall_update_request_rule *SoftLayer_Network_Firewall_Update_Request_Rule) ValidateRule(ctx *slapi.RequestContext, rule SoftLayer_Network_Firewall_Update_Request_Rule, applyToComponentId int, applyToAclId int) error {
	return nil
}
