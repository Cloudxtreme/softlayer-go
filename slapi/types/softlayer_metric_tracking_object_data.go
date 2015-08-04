package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Metric_Tracking_Object_Data - SoftLayer_Metric_Tracking_Object_Data models an individual
// unit of data tracked by a SoftLayer tracking object, including the type of data polled, the date it
// was polled at, and the counter value that was measured at polling time.
type SoftLayer_Metric_Tracking_Object_Data struct {

	// Type - no documentation
	Type string `json:"type,omitempty"`

	// Counter - no documentation
	Counter slapi.Float64 `json:"counter,omitempty"`

	// DateTime - no documentation
	DateTime *time.Time `json:"dateTime,omitempty"`
}

func (softlayer_metric_tracking_object_data *SoftLayer_Metric_Tracking_Object_Data) String() string {
	return "SoftLayer_Metric_Tracking_Object_Data"
}
