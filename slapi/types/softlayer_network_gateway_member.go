package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Gateway_Member - <nil>
type SoftLayer_Network_Gateway_Member struct {

	// HardwareId - The internal identifier of the hardware for this member.
	HardwareId int `json:"hardwareId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// NetworkGatewayId - The internal identifier of the gateway this member belongs to.
	NetworkGatewayId int `json:"networkGatewayId,omitempty"`

	// Priority - The priority for this gateway member. This is set internally and cannot be provided on
	// create.
	Priority int `json:"priority,omitempty"`
}

func (softlayer_network_gateway_member *SoftLayer_Network_Gateway_Member) String() string {
	return "SoftLayer_Network_Gateway_Member"
}

// SoftLayer_Network_Gateway_Member_Extended is SoftLayer_Network_Gateway_Member with all maskable types.
type SoftLayer_Network_Gateway_Member_Extended struct {
	SoftLayer_Network_Gateway_Member

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// NetworkGateway - no documentation
	NetworkGateway *SoftLayer_Network_Gateway `json:"networkGateway,omitempty"`
}

func (softlayer_network_gateway_member *SoftLayer_Network_Gateway_Member_Extended) String() string {
	return "SoftLayer_Network_Gateway_Member"
}
