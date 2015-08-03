package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Metric_Tracking_Object_Summary -
// SoftLayer_Container_Metric_Tracking_Object_Summary This container is a parent class for summarizing
// diverse metrics.
type SoftLayer_Container_Metric_Tracking_Object_Summary struct {

	// MetricName - The name that best describes the metric being collected.
	MetricName string `json:"metricName"`
}

func (softlayer_container_metric_tracking_object_summary *SoftLayer_Container_Metric_Tracking_Object_Summary) String() string {
	return "SoftLayer_Container_Metric_Tracking_Object_Summary"
}
