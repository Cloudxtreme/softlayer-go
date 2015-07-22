package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Monitor - <nil>
type SoftLayer_Network_Monitor struct {
}

// GetIpAddressesByHardware - This will return an arrayObject of objects containing the ipaddresses.
// Using an string parameter you can send a partial ipaddress to search within a given ipaddress. You
// can also set the max limit as well using the setting the resultLimit.
func (softlayer_network_monitor *SoftLayer_Network_Monitor) GetIpAddressesByHardware(ctx *slapi.RequestContext, hardware SoftLayer_Hardware, partialIpAddress string) ([]*SoftLayer_Network_Subnet_IpAddress, error) {
	var returnValue []*SoftLayer_Network_Subnet_IpAddress
	return returnValue, nil
}

// GetIpAddressesByVirtualGuest - This will return an arrayObject of objects containing the
// ipaddresses. Using an string parameter you can send a partial ipaddress to search within a given
// ipaddress. You can also set the max limit as well using the setting the resultLimit.
func (softlayer_network_monitor *SoftLayer_Network_Monitor) GetIpAddressesByVirtualGuest(ctx *slapi.RequestContext, guest SoftLayer_Virtual_Guest, partialIpAddress string) ([]*SoftLayer_Network_Subnet_IpAddress, error) {
	var returnValue []*SoftLayer_Network_Subnet_IpAddress
	return returnValue, nil
}
