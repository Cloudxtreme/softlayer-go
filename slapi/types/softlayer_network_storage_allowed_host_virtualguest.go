package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Storage_Allowed_Host_VirtualGuest - <nil>
type SoftLayer_Network_Storage_Allowed_Host_VirtualGuest struct {

	// Resource - The SoftLayer_Virtual_Guest object which this SoftLayer_Network_Storage_Allowed_Host is
	// referencing.
	Resource *SoftLayer_Virtual_Guest `json:"resource"`
}

func (softlayer_network_storage_allowed_host_virtualguest *SoftLayer_Network_Storage_Allowed_Host_VirtualGuest) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_VirtualGuest"
}

// GetObject - <nil>
func (softlayer_network_storage_allowed_host_virtualguest *SoftLayer_Network_Storage_Allowed_Host_VirtualGuest) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage_Allowed_Host_VirtualGuest, error) {
	var returnValue *SoftLayer_Network_Storage_Allowed_Host_VirtualGuest
	return returnValue, nil
}
