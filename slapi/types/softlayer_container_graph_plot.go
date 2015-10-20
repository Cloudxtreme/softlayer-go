package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Graph_Plot - <nil>
type SoftLayer_Container_Graph_Plot struct {

	// Metric - <nil>
	Metric *SoftLayer_Container_Metric_Data_Type `json:"metric,omitempty"`

	// Unit - <nil>
	Unit string `json:"unit,omitempty"`

	// Data - <nil>
	Data []*SoftLayer_Container_Graph_Plot_Coordinate `json:"data,omitempty"`
}

func (softlayer_container_graph_plot *SoftLayer_Container_Graph_Plot) String() string {
	return "SoftLayer_Container_Graph_Plot"
}
