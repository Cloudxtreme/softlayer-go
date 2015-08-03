package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Bandwidth_Version1_Usage_Detail - The
// SoftLayer_Network_Bandwidth_Version1_Usage_Detail data type contains specific information relating
// to bandwidth utilization at a specific point in time on a given network interface.
type SoftLayer_Network_Bandwidth_Version1_Usage_Detail struct {

	// Day - Day and time this bandwidth utilization event was recorded.
	Day *time.Time `json:"day"`

	// AmountIn - no documentation
	AmountIn float64 `json:"amountIn"`

	// AmountOut - no documentation
	AmountOut float64 `json:"amountOut"`
}

func (softlayer_network_bandwidth_version1_usage_detail *SoftLayer_Network_Bandwidth_Version1_Usage_Detail) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage_Detail"
}

// SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Extended is SoftLayer_Network_Bandwidth_Version1_Usage_Detail with all maskable types.
type SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Extended struct {
	SoftLayer_Network_Bandwidth_Version1_Usage_Detail

	// BandwidthUsage - In and out bandwidth utilization for a specified time stamp.
	BandwidthUsage *SoftLayer_Network_Bandwidth_Version1_Usage `json:"bandwidthUsage"`

	// BandwidthUsageDetailType - Describes this bandwidth utilization record as on the public or private
	// network interface.
	BandwidthUsageDetailType *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type `json:"bandwidthUsageDetailType"`
}

func (softlayer_network_bandwidth_version1_usage_detail *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Extended) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage_Detail"
}
