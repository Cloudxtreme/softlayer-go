package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Usage_Detail - The SoftLayer_Network_Bandwidth_Usage_Detail data type
// contains specific information relating to bandwidth utilization at a specific point in time on a
// given network interface.
type SoftLayer_Network_Bandwidth_Usage_Detail struct {

	// AmountIn - no documentation
	AmountIn float64 `json:"amountIn,omitempty"`

	// AmountOut - no documentation
	AmountOut float64 `json:"amountOut,omitempty"`

	// BandwidthUsageDetailTypeId - ID of the bandwidth usage detail type for this record.
	BandwidthUsageDetailTypeId float64 `json:"bandwidthUsageDetailTypeId,omitempty"`
}

func (softlayer_network_bandwidth_usage_detail *SoftLayer_Network_Bandwidth_Usage_Detail) String() string {
	return "SoftLayer_Network_Bandwidth_Usage_Detail"
}

// SoftLayer_Network_Bandwidth_Usage_Detail_Extended is SoftLayer_Network_Bandwidth_Usage_Detail with all maskable types.
type SoftLayer_Network_Bandwidth_Usage_Detail_Extended struct {
	SoftLayer_Network_Bandwidth_Usage_Detail

	// TrackingObject - The tracking object this bandwidth usage record describes.
	TrackingObject *SoftLayer_Metric_Tracking_Object `json:"trackingObject,omitempty"`

	// Type - In and out bandwidth utilization for a specified time stamp.
	Type *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type `json:"type,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_network_bandwidth_usage_detail *SoftLayer_Network_Bandwidth_Usage_Detail_Extended) String() string {
	return "SoftLayer_Network_Bandwidth_Usage_Detail"
}
