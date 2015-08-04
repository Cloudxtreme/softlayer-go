package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Metric_Tracking_Object_HardwareServer - SoftLayer_Metric_Tracking_Object_HardwareServer
// models tracking objects specific to physical hardware and the data that are recorded by those
// servers.
type SoftLayer_Metric_Tracking_Object_HardwareServer struct {

	// StartDate - The date this tracker began tracking this particular resource.
	StartDate *time.Time `json:"startDate,omitempty"`

	// ResourceTableId - The identifier of the existing resource this object is attempting to track.
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Label - no documentation
	Label string `json:"label,omitempty"`

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data,omitempty"`

	// BillingCyclePrivateUsageIn - The total private inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageIn slapi.Float64 `json:"billingCyclePrivateUsageIn,omitempty"`

	// BillingCyclePrivateUsageOut - The total private outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePrivateUsageOut slapi.Float64 `json:"billingCyclePrivateUsageOut,omitempty"`

	// BillingCyclePrivateUsageTotal - The total private bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePrivateUsageTotal uint `json:"billingCyclePrivateUsageTotal,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Hardware_Server `json:"resource,omitempty"`

	// BillingCyclePublicBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePublicBandwidthUsage *SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePublicBandwidthUsage,omitempty"`

	// BillingCyclePublicUsageTotal - The total public bandwidth for this item's resource for the current
	// billing cycle.
	BillingCyclePublicUsageTotal uint `json:"billingCyclePublicUsageTotal,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Metric_Tracking_Object_Type `json:"type,omitempty"`

	// BillingCyclePrivateBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One
	// object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCyclePrivateBandwidthUsage,omitempty"`

	// BillingCyclePublicUsageIn - The total public inbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageIn slapi.Float64 `json:"billingCyclePublicUsageIn,omitempty"`

	// BillingCyclePublicUsageOut - The total public outbound bandwidth for this item's resource for the
	// current billing cycle.
	BillingCyclePublicUsageOut slapi.Float64 `json:"billingCyclePublicUsageOut,omitempty"`

	// BillingCycleBandwidthUsageCount - A count of the raw bandwidth usage data for the current billing
	// cycle. One object is returned for each network this server is attached to.
	BillingCycleBandwidthUsageCount uint64 `json:"billingCycleBandwidthUsageCount,omitempty"`

	// BillingCycleBandwidthUsage - The raw bandwidth usage data for the current billing cycle. One object
	// is returned for each network this server is attached to.
	BillingCycleBandwidthUsage []*SoftLayer_Network_Bandwidth_Usage `json:"billingCycleBandwidthUsage,omitempty"`

	// BillingCyclePrivateBandwidthUsageCount - A count of the raw bandwidth usage data for the current
	// billing cycle. One object is returned for each network this server is attached to.
	BillingCyclePrivateBandwidthUsageCount uint64 `json:"billingCyclePrivateBandwidthUsageCount,omitempty"`
}

func (softlayer_metric_tracking_object_hardwareserver *SoftLayer_Metric_Tracking_Object_HardwareServer) String() string {
	return "SoftLayer_Metric_Tracking_Object_HardwareServer"
}
