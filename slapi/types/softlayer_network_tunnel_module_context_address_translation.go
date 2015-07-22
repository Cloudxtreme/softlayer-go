package sl

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

	// CustomerIpAddressRecord - The ip address record for the ip that will receive the encrypted traffic
	// from the IPSec network tunnel.
	CustomerIpAddressRecord *SoftLayer_Network_Customer_Subnet_IpAddress `json:"customerIpAddressRecord"`

	// Id - no documentation
	Id int `json:"id"`

	// InternalIpAddress - The ip address record that will deliver the encrypted traffic.
	InternalIpAddress string `json:"internalIpAddress"`

	// InternalIpAddressId - The unique identifier for the ip address record that will deliver the
	// encrypted traffic.
	InternalIpAddressId int `json:"internalIpAddressId"`

	// InternalIpAddressRecord - The ip address record for the ip that will deliver the encrypted traffic
	// from the IPSec network tunnel.
	InternalIpAddressRecord *SoftLayer_Network_Subnet_IpAddress `json:"internalIpAddressRecord"`

	// NetworkTunnelContext - The IPSec network tunnel an address translation belongs to.
	NetworkTunnelContext *SoftLayer_Network_Tunnel_Module_Context `json:"networkTunnelContext"`

	// NetworkTunnelContextId - An address translation's network tunnel identifier.
	NetworkTunnelContextId int `json:"networkTunnelContextId"`

	// Notes - A name or description given to an address translation to help identify the address
	// translation.
	Notes string `json:"notes"`
}
