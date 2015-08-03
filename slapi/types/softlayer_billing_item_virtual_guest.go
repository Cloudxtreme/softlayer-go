package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Virtual_Guest - The SoftLayer_Billing_Item_Virtual_Guest data type contains
// general information relating to a single SoftLayer billing item for guests.
type SoftLayer_Billing_Item_Virtual_Guest struct {

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

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn"`

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut"`

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this virtual server for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicBandwidthUsageCount - A count of the raw public bandwidth usage data for the
	// current billing cycle.
	BillingCyclePublicBandwidthUsageCount uint64 `json:"billingCyclePublicBandwidthUsageCount"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this virtual server for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`

	// MonitoringBillingItemCount - no documentation
	MonitoringBillingItemCount uint64 `json:"monitoringBillingItemCount"`

	// MonitoringBillingItems - <nil>
	MonitoringBillingItems []*SoftLayer_Billing_Item `json:"monitoringBillingItems"`

	// Resource - no documentation
	Resource *SoftLayer_Virtual_Guest `json:"resource"`

	// ResourceTableId - The resource (unique identifier) for a server billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_virtual_guest *SoftLayer_Billing_Item_Virtual_Guest) String() string {
	return "SoftLayer_Billing_Item_Virtual_Guest"
}
