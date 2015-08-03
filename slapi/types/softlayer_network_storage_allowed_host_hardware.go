package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Storage_Allowed_Host_Hardware - <nil>
type SoftLayer_Network_Storage_Allowed_Host_Hardware struct {

	// Resource - The SoftLayer_Hardware object which this SoftLayer_Network_Storage_Allowed_Host is
	// referencing.
	Resource *SoftLayer_Hardware `json:"resource"`
}

func (softlayer_network_storage_allowed_host_hardware *SoftLayer_Network_Storage_Allowed_Host_Hardware) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_Hardware"
}

// GetObject - <nil>
func (softlayer_network_storage_allowed_host_hardware *SoftLayer_Network_Storage_Allowed_Host_Hardware) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage_Allowed_Host_Hardware, error) {
	var returnValue *SoftLayer_Network_Storage_Allowed_Host_Hardware
	return returnValue, nil
}
