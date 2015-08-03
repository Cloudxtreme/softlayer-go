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
	AvgMemoryUsageInBillingCycle int `json:"avgMemoryUsageInBillingCycle"`

	// CurrentBillCycleEnd - no documentation
	CurrentBillCycleEnd *time.Time `json:"currentBillCycleEnd"`

	// CurrentBillCycleStart - no documentation
	CurrentBillCycleStart *time.Time `json:"currentBillCycleStart"`

	// LastInstanceCount - The last count of instances this platform was hosting.
	LastInstanceCount int `json:"lastInstanceCount"`

	// LastMemoryUsageAmount - no documentation
	LastMemoryUsageAmount int `json:"lastMemoryUsageAmount"`

	// LastPollTime - The last time this virtual host was polled for metrics.
	LastPollTime *time.Time `json:"lastPollTime"`

	// MaxInstanceInBillingCycle - The max number of instances hosted thus far in this billing cycle.
	MaxInstanceInBillingCycle int `json:"maxInstanceInBillingCycle"`

	// PreviousBillCycleEnd - no documentation
	PreviousBillCycleEnd *time.Time `json:"previousBillCycleEnd"`

	// PreviousBillCycleStart - no documentation
	PreviousBillCycleStart *time.Time `json:"previousBillCycleStart"`

	// VirtualPlatformName - no documentation
	VirtualPlatformName string `json:"virtualPlatformName"`
}

func (softlayer_container_metric_tracking_object_virtual_host_summary *SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary) String() string {
	return "SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Summary"
}
