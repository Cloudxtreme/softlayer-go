package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Network_Bandwidth_Version1_Usage_Detail - The
// SoftLayer_Network_Bandwidth_Version1_Usage_Detail data type contains specific information relating
// to bandwidth utilization at a specific point in time on a given network interface.
type SoftLayer_Network_Bandwidth_Version1_Usage_Detail struct {

	// AmountIn - no documentation
	AmountIn slapi.Float64 `json:"amountIn,omitempty"`

	// AmountOut - no documentation
	AmountOut slapi.Float64 `json:"amountOut,omitempty"`

	// Day - Day and time this bandwidth utilization event was recorded.
	Day *time.Time `json:"day,omitempty"`

	// BandwidthUsage - In and out bandwidth utilization for a specified time stamp.
	BandwidthUsage *SoftLayer_Network_Bandwidth_Version1_Usage `json:"bandwidthUsage,omitempty"`

	// BandwidthUsageDetailType - Describes this bandwidth utilization record as on the public or private
	// network interface.
	BandwidthUsageDetailType *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type `json:"bandwidthUsageDetailType,omitempty"`
}

func (softlayer_network_bandwidth_version1_usage_detail *SoftLayer_Network_Bandwidth_Version1_Usage_Detail) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage_Detail"
}
