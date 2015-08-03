package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Virtual_Dedicated_Rack - A SoftLayer_Billing_Item_Virtual_Dedicated_Rack data
// type models the billing information for a single bandwidth pooling. Bandwidth pooling members share
// their public bandwidth allocations, and incur overage charges instead of the overages on individual
// rack members. Virtual rack billing items are the parent items for all of it's rack membership
// billing items.
type SoftLayer_Billing_Item_Virtual_Dedicated_Rack struct {
}

func (softlayer_billing_item_virtual_dedicated_rack *SoftLayer_Billing_Item_Virtual_Dedicated_Rack) String() string {
	return "SoftLayer_Billing_Item_Virtual_Dedicated_Rack"
}

// SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended is SoftLayer_Billing_Item_Virtual_Dedicated_Rack with all maskable types.
type SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended struct {
	SoftLayer_Billing_Item_Virtual_Dedicated_Rack

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage"`

	// BillingCyclePrivateBandwidthUsage - The raw private bandwidth usage data for the current billing
	// cycle.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage"`

	// BillingCyclePrivateUsageIn - The total private network inbound bandwidth for this virtual rack for
	// the current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn"`

	// BillingCyclePrivateUsageOut - The total private network outbound bandwidth for this virtual rack for
	// the current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut"`

	// BillingCyclePrivateUsageTotal - The total private network bandwidth for this virtual rack for the
	// current billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this virtual rack for the current
	// billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this virtual rack for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal"`

	// Resource - The virtual rack that a virtual rack billing item is associated with.
	Resource *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"resource"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw private bandwidth usage data for the
	// current billing cycle.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount"`

	// BillingCyclePublicBandwidthUsageCount - A count of the raw public bandwidth usage data for the
	// current billing cycle.
	BillingCyclePublicBandwidthUsageCount uint64 `json:"billingCyclePublicBandwidthUsageCount"`

	// BillingCyclePublicBandwidthUsage - The raw public bandwidth usage data for the current billing
	// cycle.
	BillingCyclePublicBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this virtual rack for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network a virtual rack is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount"`
}

func (softlayer_billing_item_virtual_dedicated_rack *SoftLayer_Billing_Item_Virtual_Dedicated_Rack_Extended) String() string {
	return "SoftLayer_Billing_Item_Virtual_Dedicated_Rack"
}
