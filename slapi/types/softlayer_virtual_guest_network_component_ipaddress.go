package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Network_Component_IpAddress - The
// SoftLayer_Virtual_Guest_Network_Component_IpAddress data type contains general information relating
// to the binding of a single network component to a single SoftLayer IP address.
type SoftLayer_Virtual_Guest_Network_Component_IpAddress struct {

	// IpAddress - The IP address associated with this object's network component.
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// IpAddressId - The unique ID of the [[SoftLayer_Network_Subnet_ipAddress|ip address]] this virtual IP
	// address is associated with.
	IpAddressId int `json:"ipAddressId"`

	// NetworkComponent - The network component associated with this object's IP address.
	NetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"networkComponent"`

	// Port - The port that a network component has reserved. This field is only required for some IP
	// address types.
	Port int `json:"port"`

	// Type - The type of IP that this IP address record references. Some examples are for the network
	// component's primary IP address and which represents the IP information for logging into a computing
	// instance's console.
	Type string `json:"type"`
}

func (softlayer_virtual_guest_network_component_ipaddress *SoftLayer_Virtual_Guest_Network_Component_IpAddress) String() string {
	return "SoftLayer_Virtual_Guest_Network_Component_IpAddress"
}
