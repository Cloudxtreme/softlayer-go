package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute - The
// SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute data type contains information relating
// to a single firewall rule.
type SoftLayer_Container_Utility_Network_Firewall_Rule_Attribute struct {

	// Actions - no documentation
	Actions []string `json:"actions"`

	// MaximumRuleCount - no documentation
	MaximumRuleCount int `json:"maximumRuleCount"`

	// Protocols - no documentation
	Protocols []string `json:"protocols"`

	// SourceIpSubnetMasks - The valid source ip subnet masks for use with rules.
	SourceIpSubnetMasks []*SoftLayer_Container_Utility_Network_Subnet_Mask_Generic_Detail `json:"sourceIpSubnetMasks"`
}
