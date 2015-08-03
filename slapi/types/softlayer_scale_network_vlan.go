package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Network_Vlan - <nil>
type SoftLayer_Scale_Network_Vlan struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// NetworkVlan - no documentation
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan"`

	// NetworkVlanId - no documentation
	NetworkVlanId int `json:"networkVlanId"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// ScaleGroupId - The identifier of the group this network reference applies to.
	ScaleGroupId int `json:"scaleGroupId"`
}

func (softlayer_scale_network_vlan *SoftLayer_Scale_Network_Vlan) String() string {
	return "SoftLayer_Scale_Network_Vlan"
}

// CreateObject - <nil>
func (softlayer_scale_network_vlan *SoftLayer_Scale_Network_Vlan) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Network_Vlan) (*SoftLayer_Scale_Network_Vlan, error) {
	var returnValue *SoftLayer_Scale_Network_Vlan
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_scale_network_vlan *SoftLayer_Scale_Network_Vlan) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_network_vlan *SoftLayer_Scale_Network_Vlan) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Network_Vlan, error) {
	var returnValue *SoftLayer_Scale_Network_Vlan
	return returnValue, nil
}
