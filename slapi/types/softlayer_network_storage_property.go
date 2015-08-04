package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Property - A property provides additional information about a volume which
// it is assigned to. This information can range from "Mountable" flags to utilized snapshot space.
type SoftLayer_Network_Storage_Property struct {

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// VolumeId - no documentation
	VolumeId int `json:"volumeId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`
}

func (softlayer_network_storage_property *SoftLayer_Network_Storage_Property) String() string {
	return "SoftLayer_Network_Storage_Property"
}

// SoftLayer_Network_Storage_Property_Extended is SoftLayer_Network_Storage_Property with all maskable types.
type SoftLayer_Network_Storage_Property_Extended struct {
	SoftLayer_Network_Storage_Property

	// Type - The type provides a standardized definition for a property.
	Type *SoftLayer_Network_Storage_Property_Type `json:"type,omitempty"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume,omitempty"`
}

func (softlayer_network_storage_property *SoftLayer_Network_Storage_Property_Extended) String() string {
	return "SoftLayer_Network_Storage_Property"
}
