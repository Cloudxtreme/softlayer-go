package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Hardware - The SoftLayer_Billing_Item_Hardware data type contains general
// information relating to a single SoftLayer billing item for hardware.
type SoftLayer_Billing_Item_Hardware struct {

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw private bandwidth usage data for the
	// current billing cycle.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount"`

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this hardware for the current
	// billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn"`

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this hardware for the current
	// billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut"`

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this hardware for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicBandwidthUsageCount - A count of the raw public bandwidth usage data for the
	// current billing cycle.
	BillingCyclePublicBandwidthUsageCount uint64 `json:"billingCyclePublicBandwidthUsageCount"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this hardware for the current
	// billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this hardware for the current
	// billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this hardware for the current billing
	// cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`

	// LockboxNetworkStorage - no documentation
	LockboxNetworkStorage *SoftLayer_Billing_Item_Network_Storage `json:"lockboxNetworkStorage"`

	// MonitoringBillingItemCount - no documentation
	MonitoringBillingItemCount uint64 `json:"monitoringBillingItemCount"`

	// MonitoringBillingItems - <nil>
	MonitoringBillingItems []*SoftLayer_Billing_Item `json:"monitoringBillingItems"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware_Server `json:"resource"`

	// ResourceTableId - The resource (unique identifier) for a server billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_hardware *SoftLayer_Billing_Item_Hardware) String() string {
	return "SoftLayer_Billing_Item_Hardware"
}
