package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Subnet_IpAddress - SoftLayer_Container_Subnet_IPAddress models an IP v4
// address as it exists as a member of it's subnet, letting the user know if it is a network
// identifier, gateway, broadcast, or useable address. Addresses that are neither the network
// identifier nor the gateway nor the broadcast addresses are usable by SoftLayer servers.
type SoftLayer_Container_Network_Subnet_IpAddress struct {

	// Hardware - The hardware that an IP address is associated with.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// IsBroadcastAddress - Whether an IP address is its subnet's broadcast address.
	IsBroadcastAddress bool `json:"isBroadcastAddress"`

	// IsGatewayAddress - Whether an IP address is its subnet's gateway address. Gateway addresses exist on
	// SoftLayer's routers and many not be assigned to servers.
	IsGatewayAddress bool `json:"isGatewayAddress"`

	// IsNetworkAddress - Whether an IP address is its subnet's network identifier address.
	IsNetworkAddress bool `json:"isNetworkAddress"`
}
