package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

	// Type - no documentation
	Type *SoftLayer_Metric_Tracking_Object_Type `json:"type"`
}

// GetBackboneBandwidthGraph - Retrieve a PNG image of the last 24 hours of bandwidth usage of one of
// SoftLayer's network backbones.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetBackboneBandwidthGraph(commonOptions *slapi.CommonOptions, graphTitle string) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetBandwidthData - Retrieve a collection of raw bandwidth data from an individual public or private
// network tracking object. Raw data is ideal if you with to employ your own traffic storage and
// graphing systems.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetBandwidthData(commonOptions *slapi.CommonOptions, startDateTime time.Time, endDateTime time.Time, type_ string, rollupSeconds int) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthGraph - Retrieve a PNG image of a bandwidth graph representing the bandwidth usage over
// time recorded by SofTLayer's bandwidth pollers.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetBandwidthGraph(commonOptions *slapi.CommonOptions, startDateTime time.Time, endDateTime time.Time, graphType string, fontSize int, graphWidth int, graphHeight int, doNotShowTimeZone bool) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetBandwidthMinMax - Retrieve the total maximum and minimum amounts of bandwidth recorded by a
// tracking object. The first data object returned is minimum bandwidth, and the second object returned
// is the maximum data polled.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetBandwidthMinMax(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetBandwidthTotal - Retrieve the total amount of bandwidth recorded by a tracking object within the
// given date range. This method will only work on SoftLayer_Metric_Tracking_Object for
// SoftLayer_Hardware objects, and SoftLayer_Virtual_Guest objects.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetBandwidthTotal(commonOptions *slapi.CommonOptions, startDateTime time.Time, endDateTime time.Time, direction string, type_ string) (uint64, error) {
	var returnValue uint64
	return returnValue, nil
}

// GetCustomGraphData - Returns a graph container instance that is populated with metric data for the
// tracking object.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetCustomGraphData(commonOptions *slapi.CommonOptions, graphContainer SoftLayer_Container_Graph) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetDetailsForDateRange - Retrieve a collection of detailed metric data over a date range. Ideal if
// you want to employ your own graphing systems. Note not all metrics support this method. Those that
// do not return null.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetDetailsForDateRange(commonOptions *slapi.CommonOptions, startDate time.Time, endDate time.Time, graphType []string) ([]*SoftLayer_Container_Metric_Tracking_Object_Details, error) {
	var returnValue []*SoftLayer_Container_Metric_Tracking_Object_Details
	return returnValue, nil
}

// GetGraph - no documentation
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetGraph(commonOptions *slapi.CommonOptions, startDateTime time.Time, endDateTime time.Time, graphType []string) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetMetricDataTypes - Returns a collection of metric data types that can be retrieved for a metric
// tracking object.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetMetricDataTypes(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Container_Metric_Data_Type, error) {
	var returnValue []*SoftLayer_Container_Metric_Data_Type
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Metric_Tracking_Object object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Metric_Tracking_Object
// service. You can only tracking objects that are associated with your SoftLayer account or services.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Metric_Tracking_Object, error) {
	var returnValue *SoftLayer_Metric_Tracking_Object
	return returnValue, nil
}

// GetSummary - Retrieve a metric summary. Ideal if you want to employ your own graphing systems. Note
// not all metric types contain a summary. These return null.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetSummary(commonOptions *slapi.CommonOptions, graphType string) (*SoftLayer_Container_Metric_Tracking_Object_Summary, error) {
	var returnValue *SoftLayer_Container_Metric_Tracking_Object_Summary
	return returnValue, nil
}

// GetSummaryData - Returns summarized metric data for the date range, metric type and summary period
// provided.
func (softlayer_metric_tracking_object *SoftLayer_Metric_Tracking_Object) GetSummaryData(commonOptions *slapi.CommonOptions, startDateTime time.Time, endDateTime time.Time, validTypes []SoftLayer_Container_Metric_Data_Type, summaryPeriod int) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}
