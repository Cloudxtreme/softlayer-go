package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Configuration_Storage_Group_Template_Group - Single storage group(array) used in a storage
// group template. If a server configuration requires a raid configuration this object will describe a
// single array to be configured.
type SoftLayer_Configuration_Storage_Group_Template_Group struct {

	// HardDrivesString - Comma delimited integers of drive indexes for the array. This can also be the
	// string 'all' to specify all drives in the server
	HardDrivesString string `json:"hardDrivesString,omitempty"`

	// OrderIndex - no documentation
	OrderIndex int `json:"orderIndex,omitempty"`

	// Size - Size of array. Must be within limitations of the smallest drive and raid mode
	Size slapi.Float64 `json:"size,omitempty"`

	// Grow - no documentation
	Grow bool `json:"grow,omitempty"`

	// Type - <nil>
	Type *SoftLayer_Configuration_Storage_Group_Array_Type `json:"type,omitempty"`
}

func (softlayer_configuration_storage_group_template_group *SoftLayer_Configuration_Storage_Group_Template_Group) String() string {
	return "SoftLayer_Configuration_Storage_Group_Template_Group"
}
