package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Auxiliary_Network_Status - <nil>
type SoftLayer_Auxiliary_Network_Status struct {
}

func (softlayer_auxiliary_network_status *SoftLayer_Auxiliary_Network_Status) String() string {
	return "SoftLayer_Auxiliary_Network_Status"
}

// GetNetworkStatus - Return the current network status of and latency information for a given target
// from numerous points around the world. Valid Targets: * ALL * * * * * * * * * *
func (softlayer_auxiliary_network_status *SoftLayer_Auxiliary_Network_Status) GetNetworkStatus(ctx *slapi.RequestContext, target string) ([]*SoftLayer_Container_Auxiliary_Network_Status_Reading, error) {
	var returnValue []*SoftLayer_Container_Auxiliary_Network_Status_Reading
	return returnValue, nil
}
