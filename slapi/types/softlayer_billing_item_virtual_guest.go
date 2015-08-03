package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Virtual_Guest - The SoftLayer_Billing_Item_Virtual_Guest data type contains
// general information relating to a single SoftLayer billing item for guests.
type SoftLayer_Billing_Item_Virtual_Guest struct {

	// ResourceTableId - The resource (unique identifier) for a server billing item.
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_billing_item_virtual_guest *SoftLayer_Billing_Item_Virtual_Guest) String() string {
	return "SoftLayer_Billing_Item_Virtual_Guest"
}

// SoftLayer_Billing_Item_Virtual_Guest_Extended is SoftLayer_Billing_Item_Virtual_Guest with all maskable types.
type SoftLayer_Billing_Item_Virtual_Guest_Extended struct {
	SoftLayer_Billing_Item_Virtual_Guest

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this virtual server for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut"`

	// Resource - no documentation
	Resource *SoftLayer_Virtual_Guest `json:"resource"`

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut"`

	// MonitoringBillingItems - <nil>
	MonitoringBillingItems []*SoftLayer_Billing_Item `json:"monitoringBillingItems"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object will be returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw private bandwidth usage data for the
	// current billing cycle.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount"`

	// MonitoringBillingItemCount - no documentation
	MonitoringBillingItemCount uint64 `json:"monitoringBillingItemCount"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// will be returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this virtual server for the
	// current billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn"`

	// BillingCyclePublicBandwidthUsageCount - A count of the raw public bandwidth usage data for the
	// current billing cycle.
	BillingCyclePublicBandwidthUsageCount uint64 `json:"billingCyclePublicBandwidthUsageCount"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this virtual server for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`
}

func (softlayer_billing_item_virtual_guest *SoftLayer_Billing_Item_Virtual_Guest_Extended) String() string {
	return "SoftLayer_Billing_Item_Virtual_Guest"
}
