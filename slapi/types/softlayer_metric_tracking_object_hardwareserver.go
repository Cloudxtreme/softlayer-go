package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Metric_Tracking_Object_HardwareServer - SoftLayer_Metric_Tracking_Object_HardwareServer
// models tracking objects specific to physical hardware and the data that are recorded by those
// servers.
type SoftLayer_Metric_Tracking_Object_HardwareServer struct {

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// BillingCyclePrivateBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw bandwidth usage data for the current
	// billing cycle. One object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount"`

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn"`

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut"`

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal"`

	// BillingCyclePublicBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware_Server `json:"resource"`
}

func (softlayer_metric_tracking_object_hardwareserver *SoftLayer_Metric_Tracking_Object_HardwareServer) String() string {
	return "SoftLayer_Metric_Tracking_Object_HardwareServer"
}
