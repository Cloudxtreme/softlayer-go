package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack -
// SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack models tracking objects specific to virtual
// dedicated racks. Bandwidth Pooling aggregate the bandwidth used by multiple servers within the rack.
type SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack struct {
}

func (softlayer_metric_tracking_object_virtualdedicatedrack *SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack) String() string {
	return "SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack"
}

// SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack_Extended is SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack with all maskable types.
type SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack_Extended struct {
	SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut"`

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// Resource - no documentation
	Resource *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"resource"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw bandwidth usage data for the current
	// billing cycle. One object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCyclePrivateBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn"`

	// BillingCyclePublicBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`
}

func (softlayer_metric_tracking_object_virtualdedicatedrack *SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack_Extended) String() string {
	return "SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack"
}
