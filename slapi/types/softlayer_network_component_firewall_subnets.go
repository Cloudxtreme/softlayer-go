package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Component_Firewall_Subnets - A SoftLayer_Network_Component_Firewall_Subnets object
// type represents the current linked subnets and contains relative information. Use the [[SoftLayer
// Network Firewall Update Request]] service to submit a firewall update request. Use the [[SoftLayer
// Network Firewall Template]] service to pull SoftLayer recommended rule set templates.
type SoftLayer_Network_Component_Firewall_Subnets struct {

	// ApplyServerRulesFlag - A boolean flag that indicates whether the subnet should receive all the rules
	// intended for the host on this context slot.
	ApplyServerRulesFlag bool `json:"applyServerRulesFlag"`

	// NetworkComponentFirewall - The network component firewall that write rules for this subnet.
	NetworkComponentFirewall *SoftLayer_Network_Component_Firewall `json:"networkComponentFirewall"`

	// Subnet - The subnet that this link binds to the network component firewall.
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// SubnetId - The unique identifier of the subnet being linked to the network component firewall.
	SubnetId int `json:"subnetId"`
}
