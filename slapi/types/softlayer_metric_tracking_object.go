package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Metric_Tracking_Object - Metric tracking objects provides a common interface to all
// metrics provided by SoftLayer. These metrics range from network component traffic for a server to
// aggregated Bandwidth Pooling traffic and more. Every object within SoftLayer's range of objects that
// has data that can be tracked over time has an associated tracking object. Use the
// [[SoftLayer_Metric_Tracking_Object]] service to retrieve raw and graph data from a tracking object.
type SoftLayer_Metric_Tracking_Object struct {

	// Data - no documentation
	Data []*SoftLayer_Metric_Tracking_Object_Data `json:"data"`

	// Id - no documentation
	Id int `json:"id"`

	// Label - no documentation
	Label string `json:"label"`

	// ResourceTableId - The identifier of the existing resource this object is attempting to track.
	ResourceTableId int `json:"resourceTableId"`

	// StartDate - The date this tracker began tracking this particular resource.
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) String() string {
	return "SoftLayer_Metric_Tracking_Object"
}

// SoftLayer_Metric_Tracking_Object_Extended is SoftLayer_Metric_Tracking_Object with all maskable types.
type SoftLayer_Metric_Tracking_Object_Extended struct {
	SoftLayer_Metric_Tracking_Object

	// Type - no documentation
	Type *SoftLayer_Metric_Tracking_Object_Type `json:"type"`
}

func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object_Extended) String() string {
	return "SoftLayer_Metric_Tracking_Object"
}
