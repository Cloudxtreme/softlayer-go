package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Bandwidth_Usage_Detail - The SoftLayer_Network_Bandwidth_Usage_Detail data type
// contains specific information relating to bandwidth utilization at a specific point in time on a
// given network interface.
type SoftLayer_Network_Bandwidth_Usage_Detail struct {

	// AmountIn - no documentation
	AmountIn slapi.Float64 `json:"amountIn,omitempty"`

	// AmountOut - no documentation
	AmountOut slapi.Float64 `json:"amountOut,omitempty"`

	// BandwidthUsageDetailTypeId - ID of the bandwidth usage detail type for this record.
	BandwidthUsageDetailTypeId slapi.Float64 `json:"bandwidthUsageDetailTypeId,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// TrackingObject - The tracking object this bandwidth usage record describes.
	TrackingObject *SoftLayer_Metric_Tracking_Object `json:"trackingObject,omitempty"`

	// Type - In and out bandwidth utilization for a specified time stamp.
	Type *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type `json:"type,omitempty"`
}

func (softlayer_network_bandwidth_usage_detail *SoftLayer_Network_Bandwidth_Usage_Detail) String() string {
	return "SoftLayer_Network_Bandwidth_Usage_Detail"
}
