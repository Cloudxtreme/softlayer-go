package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary -
// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary This container summarizes a virtual
// host's metric data.
type SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary struct {

	// AvgMemoryUsageInBillingCycle - The average amount of memory usage thus far in this billing cycle.
	AvgMemoryUsageInBillingCycle int `json:"avgMemoryUsageInBillingCycle,omitempty"`

	// LastMemoryUsageAmount - no documentation
	LastMemoryUsageAmount int `json:"lastMemoryUsageAmount,omitempty"`

	// MaxInstanceInBillingCycle - The max number of instances hosted thus far in this billing cycle.
	MaxInstanceInBillingCycle int `json:"maxInstanceInBillingCycle,omitempty"`

	// PreviousBillCycleStart - no documentation
	PreviousBillCycleStart *time.Time `json:"previousBillCycleStart,omitempty"`

	// VirtualPlatformName - no documentation
	VirtualPlatformName string `json:"virtualPlatformName,omitempty"`

	// CurrentBillCycleEnd - no documentation
	CurrentBillCycleEnd *time.Time `json:"currentBillCycleEnd,omitempty"`

	// CurrentBillCycleStart - no documentation
	CurrentBillCycleStart *time.Time `json:"currentBillCycleStart,omitempty"`

	// LastInstanceCount - The last count of instances this platform was hosting.
	LastInstanceCount int `json:"lastInstanceCount,omitempty"`

	// LastPollTime - The last time this virtual host was polled for metrics.
	LastPollTime *time.Time `json:"lastPollTime,omitempty"`

	// PreviousBillCycleEnd - no documentation
	PreviousBillCycleEnd *time.Time `json:"previousBillCycleEnd,omitempty"`

	// MetricName - The name that best describes the metric being collected.
	MetricName string `json:"metricName,omitempty"`
}

func (softlayer_container_metric_tracking_object_virtual_host_summary *SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary) String() string {
	return "SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary"
}
