package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack -
// SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack models tracking objects specific to virtual
// dedicated racks. Bandwidth Pooling aggregate the bandwidth used by multiple servers within the rack.
type SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// StartDate - The date this tracker began tracking this particular resource.
	StartDate *time.Time `json:"startDate,omitempty"`

	// ResourceTableId - The identifier of the existing resource this object is attempting to track.
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`
}

func (softlayer_metric_tracking_object_virtualdedicatedrack *SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack) String() string {
	return "SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack"
}

// SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack_Extended is SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack with all maskable types.
type SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack_Extended struct {
	SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`

	// BillingCyclePublicBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageIn float64 `json:"billingCyclePublicUsageIn,omitempty"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageOut float64 `json:"billingCyclePublicUsageOut,omitempty"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal,omitempty"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw bandwidth usage data for the current
	// billing cycle. One object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`

	// BillingCyclePrivateBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageIn float64 `json:"billingCyclePrivateUsageIn,omitempty"`

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"resource,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Metric_Tracking_Object_Type `json:"type,omitempty"`

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageOut float64 `json:"billingCyclePrivateUsageOut,omitempty"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount,omitempty"`
}

func (softlayer_metric_tracking_object_virtualdedicatedrack *SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack_Extended) String() string {
	return "SoftLayer_Metric_Tracking_Object_VirtualDedicatedRack"
}
