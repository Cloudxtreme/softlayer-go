package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details -
// SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details This container details a virtual
// host's metric data.
type SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details struct {

	// Day - no documentation
	Day *time.Time `json:"day,omitempty"`

	// MaxInstances - The maximum number of guests hosted by this platform for the given day.
	MaxInstances int `json:"maxInstances,omitempty"`

	// MaxMemoryUsage - The maximum amount of memory utilized by this platform for the given day.
	MaxMemoryUsage int `json:"maxMemoryUsage,omitempty"`

	// MeanInstances - The mean number of guests hosted by this platform for the given day.
	MeanInstances slapi.Float64 `json:"meanInstances,omitempty"`

	// MeanMemoryUsage - The mean amount of memory utilized by this platform for the given day.
	MeanMemoryUsage slapi.Float64 `json:"meanMemoryUsage,omitempty"`

	// MinInstances - The minimum number of guests hosted by this platform for the given day.
	MinInstances int `json:"minInstances,omitempty"`

	// MinMemoryUsage - The minimum amount of memory utilized by this platform for the given day.
	MinMemoryUsage int `json:"minMemoryUsage,omitempty"`

	// MetricName - The name that best describes the metric being collected.
	MetricName string `json:"metricName,omitempty"`
}

func (softlayer_container_metric_tracking_object_virtual_host_details *SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details) String() string {
	return "SoftLayer_Container_Metric_Tracking_Object_Virtual_Host_Details"
}
