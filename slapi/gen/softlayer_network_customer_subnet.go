package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Customer_Subnet - The SoftLayer_Network_Customer_Subnet data type contains general
// information relating to a single customer subnet (remote).
type SoftLayer_Network_Customer_Subnet struct {

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Cidr - A subnet's Classless Inter-Domain Routing prefix. This is a number between 0 and 32
	// signifying the number of bits in a subnet's netmask. These bits separate a subnet's network address
	// from it's host addresses. It performs the same function as the ''netmask'' property, but is
	// represented as an integer.
	Cidr int `json:"cidr"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddressCount - A count of all ip addresses associated with a subnet.
	IpAddressCount uint64 `json:"ipAddressCount"`

	// IpAddresses - no documentation
	IpAddresses []*SoftLayer_Network_Customer_Subnet_IpAddress `json:"ipAddresses"`

	// Netmask - A bitmask in dotted-quad format that is used to separate a subnet's network address from
	// it's host addresses. This performs the same function as the ''cidr'' property, but is expressed in a
	// string format.
	Netmask string `json:"netmask"`

	// NetworkIdentifier - A subnet's network identifier. This is the first IP address of a subnet.
	NetworkIdentifier string `json:"networkIdentifier"`

	// TotalIpAddresses - no documentation
	TotalIpAddresses int `json:"totalIpAddresses"`
}

// CreateObject - For IPSec network tunnels, customers can create their local subnets using this
// method. After the customer is created successfully, the customer subnet can then be added to the
// IPSec network tunnel.
func (softlayer_network_customer_subnet *SoftLayer_Network_Customer_Subnet) CreateObject(templateObject SoftLayer_Network_Customer_Subnet) (*SoftLayer_Network_Customer_Subnet, error) {
	var returnValue *SoftLayer_Network_Customer_Subnet
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Customer_Subnet object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Network_Customer_Subnet
// service. You can only retrieve the subnet whose account matches the account that your portal user is
// assigned to.
func (softlayer_network_customer_subnet *SoftLayer_Network_Customer_Subnet) GetObject() (*SoftLayer_Network_Customer_Subnet, error) {
	var returnValue *SoftLayer_Network_Customer_Subnet
	return returnValue, nil
}
