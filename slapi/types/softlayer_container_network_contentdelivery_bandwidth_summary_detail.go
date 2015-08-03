package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail -
// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_File models a CDN account's overall
// bandwidth usage and overages within a given date range.
type SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail struct {

	// Duration - no documentation
	Duration float32 `json:"duration"`

	// ViewCount - no documentation
	ViewCount int `json:"viewCount"`
}

func (softlayer_container_network_contentdelivery_bandwidth_summary_detail *SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail"
}
