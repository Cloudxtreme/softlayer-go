package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total - The
// SoftLayer_Network_Bandwidth_Usage_Detail data type contains specific information relating to
// bandwidth utilization at a specific point in time on a given network interface.
type SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total struct {

	// AmountIn - no documentation
	AmountIn float64 `json:"amountIn"`

	// AmountOut - no documentation
	AmountOut float64 `json:"amountOut"`

	// BandwidthUsageDetailTypeId - ID of the bandwidth usage detail type for this record.
	BandwidthUsageDetailTypeId float64 `json:"bandwidthUsageDetailTypeId"`
}

func (softlayer_network_bandwidth_version1_usage_detail_total *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total"
}

// SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total_Extended is SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total with all maskable types.
type SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total_Extended struct {
	SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// TrackingObject - The tracking object this bandwidth usage record describes.
	TrackingObject *SoftLayer_Metric_Tracking_Object `json:"trackingObject"`

	// Type - In and out bandwidth utilization for a specified time stamp.
	Type *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Type `json:"type"`
}

func (softlayer_network_bandwidth_version1_usage_detail_total *SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total_Extended) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage_Detail_Total"
}
