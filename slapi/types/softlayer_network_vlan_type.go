package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Vlan_Type - <nil>
type SoftLayer_Network_Vlan_Type struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_network_vlan_type *SoftLayer_Network_Vlan_Type) String() string {
	return "SoftLayer_Network_Vlan_Type"
}

// GetObject - <nil>
func (softlayer_network_vlan_type *SoftLayer_Network_Vlan_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Vlan_Type, error) {
	var returnValue *SoftLayer_Network_Vlan_Type
	return returnValue, nil
}
