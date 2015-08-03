package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Customer_Subnet_IpAddress - The SoftLayer_Network_Customer_Subnet_IpAddress data
// type contains general information relating to a single Customer Subnet (Remote) IPv4 address.
type SoftLayer_Network_Customer_Subnet_IpAddress struct {

	// Id - no documentation
	Id int `json:"id"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// SubnetId - The unique identifier for the customer subnet (remote) the ip address belongs to.
	SubnetId int `json:"subnetId"`
}

func (softlayer_network_customer_subnet_ipaddress *SoftLayer_Network_Customer_Subnet_IpAddress) String() string {
	return "SoftLayer_Network_Customer_Subnet_IpAddress"
}

// SoftLayer_Network_Customer_Subnet_IpAddress_Extended is SoftLayer_Network_Customer_Subnet_IpAddress with all maskable types.
type SoftLayer_Network_Customer_Subnet_IpAddress_Extended struct {
	SoftLayer_Network_Customer_Subnet_IpAddress

	// Translations - All the address translations that are tied to an IP address.
	Translations []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation `json:"translations"`

	// TranslationCount - A count of all the address translations that are tied to an IP address.
	TranslationCount uint64 `json:"translationCount"`

	// Subnet - The customer subnet (remote) that the ip address belongs to.
	Subnet *SoftLayer_Network_Customer_Subnet `json:"subnet"`
}

func (softlayer_network_customer_subnet_ipaddress *SoftLayer_Network_Customer_Subnet_IpAddress_Extended) String() string {
	return "SoftLayer_Network_Customer_Subnet_IpAddress"
}
