package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Metric_Tracking_Object_Details -
// SoftLayer_Container_Metric_Tracking_Object_Details This container is a parent class for detailing
// diverse metrics.
type SoftLayer_Container_Metric_Tracking_Object_Details struct {

	// MetricName - The name that best describes the metric being collected.
	MetricName string `json:"metricName"`
}

func (softlayer_container_metric_tracking_object_details *SoftLayer_Container_Metric_Tracking_Object_Details) String() string {
	return "SoftLayer_Container_Metric_Tracking_Object_Details"
}
