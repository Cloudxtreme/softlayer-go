package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Storage_Allowed_Host_IpAddress - <nil>
type SoftLayer_Network_Storage_Allowed_Host_IpAddress struct {

	// Resource - The SoftLayer_Network_Subnet_IpAddress object which this
	// SoftLayer_Network_Storage_Allowed_Host is referencing.
	Resource *SoftLayer_Network_Subnet_IpAddress `json:"resource"`
}

func (softlayer_network_storage_allowed_host_ipaddress *SoftLayer_Network_Storage_Allowed_Host_IpAddress) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_IpAddress"
}

// GetObject - <nil>
func (softlayer_network_storage_allowed_host_ipaddress *SoftLayer_Network_Storage_Allowed_Host_IpAddress) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage_Allowed_Host_IpAddress, error) {
	var returnValue *SoftLayer_Network_Storage_Allowed_Host_IpAddress
	return returnValue, nil
}
