package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Subnet_IpAddress_Version6 - The SoftLayer_Network_Subnet_IpAddress data type
// contains general information relating to a single SoftLayer IPv6 address.
type SoftLayer_Network_Subnet_IpAddress_Version6 struct {

	// PublicVersion6NetworkGateway - The network gateway appliance using this address as the public IPv6
	// address.
	PublicVersion6NetworkGateway *SoftLayer_Network_Gateway `json:"publicVersion6NetworkGateway"`
}

func (softlayer_network_subnet_ipaddress_version6 *SoftLayer_Network_Subnet_IpAddress_Version6) String() string {
	return "SoftLayer_Network_Subnet_IpAddress_Version6"
}
