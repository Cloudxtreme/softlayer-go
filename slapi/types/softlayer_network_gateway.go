package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Gateway - <nil>
type SoftLayer_Network_Gateway struct {

	// PrivateIpAddressId - The internal identifier of the private IP address for this gateway.
	PrivateIpAddressId int `json:"privateIpAddressId"`

	// PublicIpAddressId - The internal identifier of the public IP address for this gateway.
	PublicIpAddressId int `json:"publicIpAddressId"`

	// StatusId - The current status of this gateway. This is always active unless there is a process
	// running to change the gateway. This can not be set on creation.
	StatusId int `json:"statusId"`

	// Name - A gateway's name. This is required on create and can be no more than 255 characters.
	Name string `json:"name"`

	// AccountId - The internal identifier of the account assigned to this gateway.
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// PrivateVlanId - The internal identifier of the private for this gateway.
	PrivateVlanId int `json:"privateVlanId"`

	// PublicVlanId - The internal identifier of the public for this gateway. This is set internally and
	// cannot be provided on create.
	PublicVlanId int `json:"publicVlanId"`

	// GroupNumber - The group number for this gateway. This is set internally and cannot be provided on
	// create.
	GroupNumber int `json:"groupNumber"`

	// NetworkSpace - A gateway's network space. Currently, only 'private' or 'both' is allowed. When this
	// value is 'private', it is a backend gateway only. Otherwise, it is a gateway for both frontend and
	// backend traffic.
	NetworkSpace string `json:"networkSpace"`

	// PublicIpv6AddressId - The internal identifier of the public IPv6 address for this gateway.
	PublicIpv6AddressId int `json:"publicIpv6AddressId"`
}

func (softlayer_network_gateway *SoftLayer_Network_Gateway) String() string {
	return "SoftLayer_Network_Gateway"
}

// SoftLayer_Network_Gateway_Extended is SoftLayer_Network_Gateway with all maskable types.
type SoftLayer_Network_Gateway_Extended struct {
	SoftLayer_Network_Gateway

	// PublicIpv6Address - no documentation
	PublicIpv6Address *SoftLayer_Network_Subnet_IpAddress `json:"publicIpv6Address"`

	// MemberCount - no documentation
	MemberCount uint64 `json:"memberCount"`

	// PrivateVlan - no documentation
	PrivateVlan *SoftLayer_Network_Vlan `json:"privateVlan"`

	// Members - no documentation
	Members []*SoftLayer_Network_Gateway_Member `json:"members"`

	// PublicIpAddress - no documentation
	PublicIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"publicIpAddress"`

	// InsideVlans - no documentation
	InsideVlans []*SoftLayer_Network_Gateway_Vlan `json:"insideVlans"`

	// PrivateIpAddress - no documentation
	PrivateIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"privateIpAddress"`

	// PublicVlan - no documentation
	PublicVlan *SoftLayer_Network_Vlan `json:"publicVlan"`

	// Status - no documentation
	Status *SoftLayer_Network_Gateway_Status `json:"status"`

	// InsideVlanCount - no documentation
	InsideVlanCount uint64 `json:"insideVlanCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_network_gateway *SoftLayer_Network_Gateway_Extended) String() string {
	return "SoftLayer_Network_Gateway"
}
