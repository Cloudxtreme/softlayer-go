package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Bandwidth_Version1_Usage - The SoftLayer_Network_Bandwidth_Version1_Usage data
// type contains general information relating to a single bandwidth usage record.
type SoftLayer_Network_Bandwidth_Version1_Usage struct {
}

func (softlayer_network_bandwidth_version1_usage *SoftLayer_Network_Bandwidth_Version1_Usage) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage"
}

// SoftLayer_Network_Bandwidth_Version1_Usage_Extended is SoftLayer_Network_Bandwidth_Version1_Usage with all maskable types.
type SoftLayer_Network_Bandwidth_Version1_Usage_Extended struct {
	SoftLayer_Network_Bandwidth_Version1_Usage

	// BandwidthAllotmentDetail - no documentation
	BandwidthAllotmentDetail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail `json:"bandwidthAllotmentDetail,omitempty"`

	// BandwidthUsageDetail - no documentation
	BandwidthUsageDetail []*SoftLayer_Network_Bandwidth_Version1_Usage_Detail `json:"bandwidthUsageDetail,omitempty"`

	// BandwidthUsageDetailCount - A count of bandwidth usage details for this hardware.
	BandwidthUsageDetailCount uint64 `json:"bandwidthUsageDetailCount,omitempty"`
}

func (softlayer_network_bandwidth_version1_usage *SoftLayer_Network_Bandwidth_Version1_Usage_Extended) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Usage"
}
