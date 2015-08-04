package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Gateway_Vlan - <nil>
type SoftLayer_Network_Gateway_Vlan struct {

	// BypassFlag - If true, this is bypassed. If false, it is routed through the gateway.
	BypassFlag bool `json:"bypassFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// NetworkGatewayId - The internal identifier of the gateway this is attached to.
	NetworkGatewayId int `json:"networkGatewayId,omitempty"`

	// NetworkVlanId - no documentation
	NetworkVlanId int `json:"networkVlanId,omitempty"`
}

func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan) String() string {
	return "SoftLayer_Network_Gateway_Vlan"
}

// SoftLayer_Network_Gateway_Vlan_Extended is SoftLayer_Network_Gateway_Vlan with all maskable types.
type SoftLayer_Network_Gateway_Vlan_Extended struct {
	SoftLayer_Network_Gateway_Vlan

	// NetworkGateway - no documentation
	NetworkGateway *SoftLayer_Network_Gateway `json:"networkGateway,omitempty"`

	// NetworkVlan - no documentation
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`
}

func (softlayer_network_gateway_vlan *SoftLayer_Network_Gateway_Vlan_Extended) String() string {
	return "SoftLayer_Network_Gateway_Vlan"
}
