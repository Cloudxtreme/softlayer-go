package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Usage - The SoftLayer_Network_Bandwidth_Usage data type contains
// specific information relating to bandwidth utilization at a specific point in time on a given
// network interface.
type SoftLayer_Network_Bandwidth_Usage struct {

	// AmountIn - no documentation
	AmountIn float64 `json:"amountIn"`

	// AmountOut - no documentation
	AmountOut float64 `json:"amountOut"`

	// BandwidthUsageDetailTypeId - ID of the bandwidth usage detail type for this record.
	BandwidthUsageDetailTypeId float64 `json:"bandwidthUsageDetailTypeId"`

	// TrackingObject - The tracking object this bandwidth usage record describes.
	TrackingObject *SoftLayer_Metric_Tracking_Object `json:"trackingObject"`

	// Type - In and out bandwidth utilization for a specified time stamp.
	Type *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type `json:"type"`
}

func (softlayer_network_bandwidth_usage *SoftLayer_Network_Bandwidth_Usage) String() string {
	return "SoftLayer_Network_Bandwidth_Usage"
}
