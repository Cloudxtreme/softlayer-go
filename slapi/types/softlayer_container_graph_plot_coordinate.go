package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Graph_Plot_Coordinate - <nil>
type SoftLayer_Container_Graph_Plot_Coordinate struct {

	// ZValue - <nil>
	ZValue slapi.Float64 `json:"zValue,omitempty"`

	// XValue - <nil>
	XValue slapi.Float64 `json:"xValue,omitempty"`

	// YValue - <nil>
	YValue slapi.Float64 `json:"yValue,omitempty"`
}

func (softlayer_container_graph_plot_coordinate *SoftLayer_Container_Graph_Plot_Coordinate) String() string {
	return "SoftLayer_Container_Graph_Plot_Coordinate"
}
