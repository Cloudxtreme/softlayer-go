package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Firewall_Update_Request - The SoftLayer_Network_Firewall_Update_Request data type
// contains information relating to a SoftLayer network firewall update request. Use the [[SoftLayer
// Network Component Firewall]] service to view current rules. Use the [[SoftLayer Network Firewall
// Template]] service to pull SoftLayer recommended rule set templates.
type SoftLayer_Network_Firewall_Update_Request struct {

	// ApplyDate - Timestamp of when the rules from the update request were applied to the firewall.
	ApplyDate *time.Time `json:"applyDate"`

	// AuthorizingUser - The user that authorized this firewall update request.
	AuthorizingUser *SoftLayer_User_Interface `json:"authorizingUser"`

	// AuthorizingUserId - The unique identifier of the user that authorized the update request.
	AuthorizingUserId int `json:"authorizingUserId"`

	// AuthorizingUserType - The type of user that authorized the update request or
	AuthorizingUserType string `json:"authorizingUserType"`

	// BypassFlag - Flag indicating whether the request is for a rule bypass configuration [0 or 1].
	BypassFlag bool `json:"bypassFlag"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// FirewallContextAccessControlListId - The unique identifier of the firewall access control list that
	// the rule set is destined for.
	FirewallContextAccessControlListId int `json:"firewallContextAccessControlListId"`

	// Guest - The downstream virtual server that the rule set will be applied to.
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// Hardware - The downstream server that the rule set will be applied to.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The unique identifier of the server that the rule set is destined to protect.
	HardwareId int `json:"hardwareId"`

	// Id - The unique identifier of the firewall update request.
	Id int `json:"id"`

	// NetworkComponentFirewall - The network component firewall that the rule set will be applied to.
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// NetworkComponentFirewallId - The unique identifier of the network component firewall that the rule
	// set is destined for.
	NetworkComponentFirewallId int `json:"networkComponentFirewallId"`

	// RuleCount - A count of the group of rules contained within the update request.
	RuleCount uint64 `json:"ruleCount"`

	// Rules - The group of rules contained within the update request.
	Rules []*SoftLayer_Network_Firewall_Update_Request_Rule `json:"rules"`
}

// CreateObject - Create a new firewall update request. The SoftLayer_Network_Firewall_Update_Request
// object passed to this function must have at least one rule. ''createObject'' returns a Boolean
// ''true'' on successful object creation or ''false'' if your firewall update request was unable to be
// created.
func (softlayer_network_firewall_update_request *SoftLayer_Network_Firewall_Update_Request) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Network_Firewall_Update_Request) (*SoftLayer_Network_Firewall_Update_Request, error) {
	var returnValue *SoftLayer_Network_Firewall_Update_Request
	return returnValue, nil
}

// GetFirewallUpdateRequestRuleAttributes - Get the possible attribute values for a firewall update
// request rule. These are the valid values which may be submitted as rule parameters for a firewall
// update request. ''getFirewallUpdateRequestRuleAttributes'' returns a
// SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute object upon success.
func (softlayer_network_firewall_update_request *SoftLayer_Network_Firewall_Update_Request) GetFirewallUpdateRequestRuleAttributes(commonOptions *slapi.CommonOptions) (*SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute, error) {
	var returnValue *SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute
	return returnValue, nil
}

// GetObject - ''getObject'' returns a SoftLayer_Network_Firewall_Update_Request object. You can only
// get historical objects for servers attached to your account that have a network firewall enabled.
// ''createObject'' inserts a new SoftLayer_Network_Firewall_Update_Request object. You can only insert
// requests for servers attached to your account that have a network firewall enabled.
// ''getFirewallUpdateRequestRuleAttributes'' Get the possible attribute values for a firewall update
// request rule.
func (softlayer_network_firewall_update_request *SoftLayer_Network_Firewall_Update_Request) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Firewall_Update_Request, error) {
	var returnValue *SoftLayer_Network_Firewall_Update_Request
	return returnValue, nil
}

// UpdateRuleNote - <nil>
func (softlayer_network_firewall_update_request *SoftLayer_Network_Firewall_Update_Request) UpdateRuleNote(commonOptions *slapi.CommonOptions, fwRule SoftLayer_Network_Component_Firewall_Rule, note string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
