package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Metric_Tracking_Object_Bandwidth_Summary - This data type provides commonly used bandwidth
// summary components for the current billing cycle.
type SoftLayer_Metric_Tracking_Object_Bandwidth_Summary struct {

	// AllocationAmount - This is the amount of bandwidth (measured in gigabytes) allocated for this
	// tracking object.
	AllocationAmount float32 `json:"allocationAmount"`

	// AllocationId - <nil>
	AllocationId int `json:"allocationId"`

	// AmountOut - The amount of outbound bandwidth (measured in gigabytes) currently used this billing
	// period. Same as $outboundBandwidthAmount. Aliased for backward compatability.
	AmountOut float64 `json:"amountOut"`

	// AverageDailyUsage - The daily average amount of outbound bandwidth usage.
	AverageDailyUsage float32 `json:"averageDailyUsage"`

	// CurrentlyOverAllocationFlag - A flag that tells whether or not this tracking object's bandwidth
	// usage is already over the allocation. 1 means yes, 0 means no.
	CurrentlyOverAllocationFlag int `json:"currentlyOverAllocationFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// OutboundBandwidthAmount - The amount of outbound bandwidth (measured in gigabytes) currently used
	// this billing period
	OutboundBandwidthAmount float64 `json:"outboundBandwidthAmount"`

	// ProjectedBandwidthUsage - The amount of bandwidth (measured in gigabytes) of projected usage, using
	// a basic average calculation of daily usage.
	ProjectedBandwidthUsage float32 `json:"projectedBandwidthUsage"`

	// ProjectedOverAllocationFlag - A flag that tells whether or not this tracking object's bandwidth
	// usage is projected to go over the allocation, based on daily average usage. 1 means yes, 0 means no.
	ProjectedOverAllocationFlag int `json:"projectedOverAllocationFlag"`
}

// GetObject - <nil>
func (softlayer_metric_tracking_object_bandwidth_summary *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Metric_Tracking_Object_Bandwidth_Summary, error) {
	var returnValue *SoftLayer_Metric_Tracking_Object_Bandwidth_Summary
	return returnValue, nil
}
