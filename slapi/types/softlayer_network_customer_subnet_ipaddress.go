package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Customer_Subnet_IpAddress - The SoftLayer_Network_Customer_Subnet_IpAddress data
// type contains general information relating to a single Customer Subnet (Remote) IPv4 address.
type SoftLayer_Network_Customer_Subnet_IpAddress struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// SubnetId - The unique identifier for the customer subnet (remote) the ip address belongs to.
	SubnetId int `json:"subnetId,omitempty"`

	// TranslationCount - A count of all the address translations that are tied to an IP address.
	TranslationCount uint64 `json:"translationCount,omitempty"`

	// Subnet - The customer subnet (remote) that the ip address belongs to.
	Subnet *SoftLayer_Network_Customer_Subnet `json:"subnet,omitempty"`

	// Translations - All the address translations that are tied to an IP address.
	Translations []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation `json:"translations,omitempty"`
}

func (softlayer_network_customer_subnet_ipaddress *SoftLayer_Network_Customer_Subnet_IpAddress) String() string {
	return "SoftLayer_Network_Customer_Subnet_IpAddress"
}
