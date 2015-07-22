package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Property - A property provides additional information about a volume which
// it is assigned to. This information can range from "Mountable" flags to utilized snapshot space.
type SoftLayer_Network_Storage_Property struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Type - The type provides a standardized definition for a property.
	Type *SoftLayer_Network_Storage_Property_Type `json:"type"`

	// Value - no documentation
	Value string `json:"value"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume"`

	// VolumeId - no documentation
	VolumeId int `json:"volumeId"`
}
