package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Gateway_Status - <nil>
type SoftLayer_Network_Gateway_Status struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_network_gateway_status *SoftLayer_Network_Gateway_Status) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Gateway_Status, error) {
	var returnValue *SoftLayer_Network_Gateway_Status
	return returnValue, nil
}
