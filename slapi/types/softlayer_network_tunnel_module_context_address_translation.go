package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Tunnel_Module_Context_Address_Translation - The
// SoftLayer_Network_Tunnel_Module_Context_Address_Translation data type contains general information
// relating to a single address translation. Information such as notes, ip addresses, along with record
// information, and network tunnel data may be retrieved.
type SoftLayer_Network_Tunnel_Module_Context_Address_Translation struct {

	// CustomerIpAddress - The ip address record that will receive the encrypted traffic.
	CustomerIpAddress string `json:"customerIpAddress"`

	// CustomerIpAddressId - The unique identifier for the ip address record that will receive the
	// encrypted traffic.
	CustomerIpAddressId int `json:"customerIpAddressId"`

	// Id - no documentation
	Id int `json:"id"`

	// InternalIpAddress - The ip address record that will deliver the encrypted traffic.
	InternalIpAddress string `json:"internalIpAddress"`

	// Notes - A name or description given to an address translation to help identify the address
	// translation.
	Notes string `json:"notes"`

	// InternalIpAddressId - The unique identifier for the ip address record that will deliver the
	// encrypted traffic.
	InternalIpAddressId int `json:"internalIpAddressId"`

	// NetworkTunnelContextId - An address translation's network tunnel identifier.
	NetworkTunnelContextId int `json:"networkTunnelContextId"`
}

// SoftLayer_Network_Tunnel_Module_Context_Address_Translation_Extended is SoftLayer_Network_Tunnel_Module_Context_Address_Translation with all maskable types.
type SoftLayer_Network_Tunnel_Module_Context_Address_Translation_Extended struct {
	SoftLayer_Network_Tunnel_Module_Context_Address_Translation

	// InternalIpAddressRecord - The ip address record for the ip that will deliver the encrypted traffic
	// from the IPSec network tunnel.
	InternalIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"internalIpAddressRecord"`

	// CustomerIpAddressRecord - The ip address record for the ip that will receive the encrypted traffic
	// from the IPSec network tunnel.
	CustomerIpAddressRecord *SoftLayer_Network_Customer_Subnet_IpAddress `json:"customerIpAddressRecord"`

	// NetworkTunnelContext - The IPSec network tunnel an address translation belongs to.
	NetworkTunnelContext *SoftLayer_Network_Tunnel_Module_Context `json:"networkTunnelContext"`
}

func (softlayer_network_tunnel_module_context_address_translation *SoftLayer_Network_Tunnel_Module_Context_Address_Translation) String() string {
	return "SoftLayer_Network_Tunnel_Module_Context_Address_Translation"
}
