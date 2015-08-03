package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Firewall_Update_Request - The SoftLayer_Network_Firewall_Update_Request data type
// contains information relating to a SoftLayer network firewall update request. Use the [[SoftLayer
// Network Component Firewall]] service to view current rules. Use the [[SoftLayer Network Firewall
// Template]] service to pull SoftLayer recommended rule set templates.
type SoftLayer_Network_Firewall_Update_Request struct {

	// FirewallContextAccessControlListId - The unique identifier of the firewall access control list that
	// the rule set is destined for.
	FirewallContextAccessControlListId int `json:"firewallContextAccessControlListId"`

	// NetworkComponentFirewallId - The unique identifier of the network component firewall that the rule
	// set is destined for.
	NetworkComponentFirewallId int `json:"networkComponentFirewallId"`

	// AuthorizingUserId - The unique identifier of the user that authorized the update request.
	AuthorizingUserId int `json:"authorizingUserId"`

	// ApplyDate - Timestamp of when the rules from the update request were applied to the firewall.
	ApplyDate *time.Time `json:"applyDate"`

	// AuthorizingUserType - The type of user that authorized the update request or
	AuthorizingUserType string `json:"authorizingUserType"`

	// BypassFlag - Flag indicating whether the request is for a rule bypass configuration [0 or 1].
	BypassFlag bool `json:"bypassFlag"`

	// HardwareId - The unique identifier of the server that the rule set is destined to protect.
	HardwareId int `json:"hardwareId"`

	// Id - The unique identifier of the firewall update request.
	Id int `json:"id"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_network_firewall_update_request *SoftLayer_Network_Firewall_Update_Request) String() string {
	return "SoftLayer_Network_Firewall_Update_Request"
}

// SoftLayer_Network_Firewall_Update_Request_Extended is SoftLayer_Network_Firewall_Update_Request with all maskable types.
type SoftLayer_Network_Firewall_Update_Request_Extended struct {
	SoftLayer_Network_Firewall_Update_Request

	// Hardware - The downstream server that the rule set will be applied to.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// AuthorizingUser - The user that authorized this firewall update request.
	AuthorizingUser *SoftLayer_User_Interface `json:"authorizingUser"`

	// Rules - The group of rules contained within the update request.
	Rules []*SoftLayer_Network_Firewall_Update_Request_Rule `json:"rules"`

	// NetworkComponentFirewall - The network component firewall that the rule set will be applied to.
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// RuleCount - A count of the group of rules contained within the update request.
	RuleCount uint64 `json:"ruleCount"`

	// Guest - The downstream virtual server that the rule set will be applied to.
	Guest *SoftLayer_Virtual_Guest `json:"guest"`
}

func (softlayer_network_firewall_update_request *SoftLayer_Network_Firewall_Update_Request_Extended) String() string {
	return "SoftLayer_Network_Firewall_Update_Request"
}
