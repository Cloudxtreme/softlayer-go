package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Daily_Usage - <nil>
type SoftLayer_Network_Storage_Daily_Usage struct {

	// BytesUsed - <nil>
	BytesUsed uint64 `json:"bytesUsed,omitempty"`

	// CdnHttpBandwidth - <nil>
	CdnHttpBandwidth uint64 `json:"cdnHttpBandwidth,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// NasVolumeId - <nil>
	NasVolumeId int `json:"nasVolumeId,omitempty"`

	// PublicBandwidthOut - <nil>
	PublicBandwidthOut uint64 `json:"publicBandwidthOut,omitempty"`
}

func (softlayer_network_storage_daily_usage *SoftLayer_Network_Storage_Daily_Usage) String() string {
	return "SoftLayer_Network_Storage_Daily_Usage"
}

// SoftLayer_Network_Storage_Daily_Usage_Extended is SoftLayer_Network_Storage_Daily_Usage with all maskable types.
type SoftLayer_Network_Storage_Daily_Usage_Extended struct {
	SoftLayer_Network_Storage_Daily_Usage

	// NasVolume - <nil>
	NasVolume *SoftLayer_Network_Storage `json:"nasVolume,omitempty"`
}

func (softlayer_network_storage_daily_usage *SoftLayer_Network_Storage_Daily_Usage_Extended) String() string {
	return "SoftLayer_Network_Storage_Daily_Usage"
}
