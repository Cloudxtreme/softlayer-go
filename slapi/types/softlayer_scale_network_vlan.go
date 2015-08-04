package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Network_Vlan - <nil>
type SoftLayer_Scale_Network_Vlan struct {

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// NetworkVlanId - no documentation
	NetworkVlanId int `json:"networkVlanId,omitempty"`

	// ScaleGroupId - The identifier of the group this network reference applies to.
	ScaleGroupId int `json:"scaleGroupId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_scale_network_vlan *SoftLayer_Scale_Network_Vlan) String() string {
	return "SoftLayer_Scale_Network_Vlan"
}

// SoftLayer_Scale_Network_Vlan_Extended is SoftLayer_Scale_Network_Vlan with all maskable types.
type SoftLayer_Scale_Network_Vlan_Extended struct {
	SoftLayer_Scale_Network_Vlan

	// NetworkVlan - no documentation
	NetworkVlan *SoftLayer_Network_Vlan `json:"networkVlan,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`
}

func (softlayer_scale_network_vlan *SoftLayer_Scale_Network_Vlan_Extended) String() string {
	return "SoftLayer_Scale_Network_Vlan"
}
