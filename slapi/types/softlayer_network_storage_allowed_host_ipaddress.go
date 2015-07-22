package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage_Allowed_Host_IpAddress - <nil>
type SoftLayer_Network_Storage_Allowed_Host_IpAddress struct {

	// Resource - The SoftLayer_Network_Subnet_IpAddress object which this
	// SoftLayer_Network_Storage_Allowed_Host is referencing.
	Resource *SoftLayer_Network_Subnet_IpAddress `json:"resource"`
}

// GetObject - <nil>
func (softlayer_network_storage_allowed_host_ipaddress *SoftLayer_Network_Storage_Allowed_Host_IpAddress) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Storage_Allowed_Host_IpAddress, error) {
	var returnValue *SoftLayer_Network_Storage_Allowed_Host_IpAddress
	return returnValue, nil
}
