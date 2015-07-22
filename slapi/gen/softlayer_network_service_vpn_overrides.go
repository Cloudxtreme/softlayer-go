package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Service_Vpn_Overrides - The SoftLayer_Network_Service_Vpn_Overrides data type
// contains information relating user ids to subnet ids when VPN access is manually configured. It is
// essentially an entry in a 'white list' of subnets a SoftLayer portal VPN user may access.
type SoftLayer_Network_Service_Vpn_Overrides struct {

	// Id - no documentation
	Id int `json:"id"`

	// Subnet - Subnet components accessible by a SoftLayer VPN portal user.
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// SubnetId - The identifier of a subnet accessible by the SoftLayer portal VPN user.
	SubnetId int `json:"subnetId"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - no documentation
	UserId int `json:"userId"`
}

// CreateObjects - no documentation
func (softlayer_network_service_vpn_overrides *SoftLayer_Network_Service_Vpn_Overrides) CreateObjects(templateObjects []SoftLayer_Network_Service_Vpn_Overrides) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - Use this method to delete a single SoftLayer portal VPN user subnet override.
func (softlayer_network_service_vpn_overrides *SoftLayer_Network_Service_Vpn_Overrides) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObjects - Use this method to delete a collection of SoftLayer portal VPN user subnet
// overrides.
func (softlayer_network_service_vpn_overrides *SoftLayer_Network_Service_Vpn_Overrides) DeleteObjects(templateObjects []SoftLayer_Network_Service_Vpn_Overrides) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_service_vpn_overrides *SoftLayer_Network_Service_Vpn_Overrides) GetObject() (*SoftLayer_Network_Service_Vpn_Overrides, error) {
	var returnValue *SoftLayer_Network_Service_Vpn_Overrides
	return returnValue, nil
}
