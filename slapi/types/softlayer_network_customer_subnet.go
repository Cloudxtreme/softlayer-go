package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Customer_Subnet - The SoftLayer_Network_Customer_Subnet data type contains general
// information relating to a single customer subnet (remote).
type SoftLayer_Network_Customer_Subnet struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Netmask - A bitmask in dotted-quad format that is used to separate a subnet's network address from
	// it's host addresses. This performs the same function as the ''cidr'' property, but is expressed in a
	// string format.
	Netmask string `json:"netmask,omitempty"`

	// NetworkIdentifier - A subnet's network identifier. This is the first IP address of a subnet.
	NetworkIdentifier string `json:"networkIdentifier,omitempty"`

	// TotalIpAddresses - no documentation
	TotalIpAddresses int `json:"totalIpAddresses,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// Cidr - A subnet's Classless Inter-Domain Routing prefix. This is a number between 0 and 32
	// signifying the number of bits in a subnet's netmask. These bits separate a subnet's network address
	// from it's host addresses. It performs the same function as the ''netmask'' property, but is
	// represented as an integer.
	Cidr int `json:"cidr,omitempty"`

	// IpAddressCount - A count of all ip addresses associated with a subnet.
	IpAddressCount uint64 `json:"ipAddressCount,omitempty"`

	// IpAddresses - no documentation
	IpAddresses []*SoftLayer_Network_Customer_Subnet_IpAddress `json:"ipAddresses,omitempty"`
}

func (softlayer_network_customer_subnet *SoftLayer_Network_Customer_Subnet) String() string {
	return "SoftLayer_Network_Customer_Subnet"
}
