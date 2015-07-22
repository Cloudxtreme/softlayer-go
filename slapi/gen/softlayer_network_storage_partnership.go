package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Partnership - A network storage partnership is used to link multiple
// volumes to each other. These partnerships describe replication hierarchies or link volume snapshots
// to their associated storage volume.
type SoftLayer_Network_Storage_Partnership struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// PartnerVolume - no documentation
	PartnerVolume *SoftLayer_Network_Storage `json:"partnerVolume"`

	// PartnerVolumeId - The child volume id which a partnership is associated with.
	PartnerVolumeId int `json:"partnerVolumeId"`

	// Type - The type provides a standardized definition for a partnership.
	Type *SoftLayer_Network_Storage_Partnership_Type `json:"type"`

	// Volume - no documentation
	Volume *SoftLayer_Network_Storage `json:"volume"`

	// VolumeId - The volume id which a partnership is associated with.
	VolumeId int `json:"volumeId"`
}
