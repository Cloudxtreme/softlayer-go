package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute - The
// SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute data type contains information relating
// to a single firewall rule.
type SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute struct {

	// Actions - no documentation
	Actions []string `json:"actions,omitempty"`

	// MaximumRuleCount - no documentation
	MaximumRuleCount int `json:"maximumRuleCount,omitempty"`

	// Protocols - no documentation
	Protocols []string `json:"protocols,omitempty"`

	// SourceIpSubnetMasks - The valid source ip subnet masks for use with rules.
	SourceIpSubnetMasks []*SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail `json:"sourceIpSubnetMasks,omitempty"`
}

func (softlayer_container_utility_network_firewall_rule_attribute *SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute) String() string {
	return "SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute"
}
