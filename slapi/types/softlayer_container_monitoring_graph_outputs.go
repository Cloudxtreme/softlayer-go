package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Monitoring_Graph_Outputs - SoftLayer_Container_Monitoring_Graph_Outputs models a
// single outbound object for a graph of given data sets.
type SoftLayer_Container_Monitoring_Graph_Outputs struct {

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// GraphError - no documentation
	GraphError string `json:"graphError,omitempty"`

	// GraphImage - The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage string `json:"graphImage,omitempty"`
}

func (softlayer_container_monitoring_graph_outputs *SoftLayer_Container_Monitoring_Graph_Outputs) String() string {
	return "SoftLayer_Container_Monitoring_Graph_Outputs"
}
