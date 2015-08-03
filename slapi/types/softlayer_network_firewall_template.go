package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Firewall_Template - The SoftLayer_Network_Firewall_Template type contains general
// information for a SoftLayer network firewall template. Firewall templates are recommend rule sets
// for use with SoftLayer Hardware Firewall (Dedicated). These optimized templates are designed to
// balance security restriction with application availability. The templates given may be altered to
// provide custom network security, or may be used as-is for basic security. At least one rule set be
// applied for the firewall to block traffic. Use the [[SoftLayer Network Component Firewall]] service
// to view current rules. Use the [[SoftLayer Network Firewall Update Request]] service to submit a
// firewall update request.
type SoftLayer_Network_Firewall_Template struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_network_firewall_template *SoftLayer_Network_Firewall_Template) String() string {
	return "SoftLayer_Network_Firewall_Template"
}

// SoftLayer_Network_Firewall_Template_Extended is SoftLayer_Network_Firewall_Template with all maskable types.
type SoftLayer_Network_Firewall_Template_Extended struct {
	SoftLayer_Network_Firewall_Template

	// RuleCount - A count of the rule set that belongs to this firewall rules template.
	RuleCount uint64 `json:"ruleCount"`

	// Rules - The rule set that belongs to this firewall rules template.
	Rules []*SoftLayer_Network_Firewall_Template_Rule `json:"rules"`
}

func (softlayer_network_firewall_template *SoftLayer_Network_Firewall_Template_Extended) String() string {
	return "SoftLayer_Network_Firewall_Template"
}
