package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Vpn_Overrides - The SoftLayer_Network_Service_Vpn_Overrides data type
// contains information relating user ids to subnet ids when VPN access is manually configured. It is
// essentially an entry in a 'white list' of subnets a SoftLayer portal VPN user may access.
type SoftLayer_Network_Service_Vpn_Overrides struct {

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// SubnetId - The identifier of a subnet accessible by the SoftLayer portal VPN user.
	SubnetId int `json:"subnetId,omitempty"`

	// Subnet - Subnet components accessible by a SoftLayer VPN portal user.
	Subnet *SoftLayer_Network_Subnet `json:"subnet,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_network_service_vpn_overrides *SoftLayer_Network_Service_Vpn_Overrides) String() string {
	return "SoftLayer_Network_Service_Vpn_Overrides"
}
