package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Metric_Tracking_Object_Bandwidth_Summary - This data type provides commonly used bandwidth
// summary components for the current billing cycle.
type SoftLayer_Metric_Tracking_Object_Bandwidth_Summary struct {

	// CurrentlyOverAllocationFlag - A flag that tells whether or not this tracking object's bandwidth
	// usage is already over the allocation. 1 means yes, 0 means no.
	CurrentlyOverAllocationFlag int `json:"currentlyOverAllocationFlag,omitempty"`

	// OutboundBandwidthAmount - The amount of outbound bandwidth (measured in gigabytes) currently used
	// this billing period
	OutboundBandwidthAmount float64 `json:"outboundBandwidthAmount,omitempty"`

	// ProjectedBandwidthUsage - The amount of bandwidth (measured in gigabytes) of projected usage, using
	// a basic average calculation of daily usage.
	ProjectedBandwidthUsage float32 `json:"projectedBandwidthUsage,omitempty"`

	// ProjectedOverAllocationFlag - A flag that tells whether or not this tracking object's bandwidth
	// usage is projected to go over the allocation, based on daily average usage. 1 means yes, 0 means no.
	ProjectedOverAllocationFlag int `json:"projectedOverAllocationFlag,omitempty"`

	// AverageDailyUsage - The daily average amount of outbound bandwidth usage.
	AverageDailyUsage float32 `json:"averageDailyUsage,omitempty"`

	// AllocationId - <nil>
	AllocationId int `json:"allocationId,omitempty"`

	// AmountOut - The amount of outbound bandwidth (measured in gigabytes) currently used this billing
	// period. Same as $outboundBandwidthAmount. Aliased for backward compatability.
	AmountOut float64 `json:"amountOut,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// AllocationAmount - This is the amount of bandwidth (measured in gigabytes) allocated for this
	// tracking object.
	AllocationAmount float32 `json:"allocationAmount,omitempty"`
}

func (softlayer_metric_tracking_object_bandwidth_summary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary) String() string {
	return "SoftLayer_Metric_Tracking_Object_Bandwidth_Summary"
}
