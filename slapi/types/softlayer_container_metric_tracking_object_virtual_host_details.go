package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details -
// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details This container details a virtual
// host's metric data.
type SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details struct {

	// Day - no documentation
	Day *time.Time `json:"day"`

	// MaxInstances - The maximum number of guests hosted by this platform for the given day.
	MaxInstances int `json:"maxInstances"`

	// MaxMemoryUsage - The maximum amount of memory utilized by this platform for the given day.
	MaxMemoryUsage int `json:"maxMemoryUsage"`

	// MeanInstances - The mean number of guests hosted by this platform for the given day.
	MeanInstances float32 `json:"meanInstances"`

	// MeanMemoryUsage - The mean amount of memory utilized by this platform for the given day.
	MeanMemoryUsage float32 `json:"meanMemoryUsage"`

	// MinInstances - The minimum number of guests hosted by this platform for the given day.
	MinInstances int `json:"minInstances"`

	// MinMemoryUsage - The minimum amount of memory utilized by this platform for the given day.
	MinMemoryUsage int `json:"minMemoryUsage"`
}

func (softlayer_container_metric_tracking_object_virtual_host_details *SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details) String() string {
	return "SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details"
}
