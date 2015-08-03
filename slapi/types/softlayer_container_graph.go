package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Graph - <nil>
type SoftLayer_Container_Graph struct {

	// Interval - no documentation
	Interval int `json:"interval"`

	// Metrics - no documentation
	Metrics []*SoftLayer_Container_Metric_Data_Type `json:"metrics"`

	// Width - no documentation
	Width int `json:"width"`

	// Height - no documentation
	Height int `json:"height"`

	// ReturnImage - no documentation
	ReturnImage bool `json:"returnImage"`

	// Template - no documentation
	Template string `json:"template"`

	// StartDatetime - no documentation
	StartDatetime string `json:"startDatetime"`

	// Title - no documentation
	Title string `json:"title"`

	// NormalizeFlag - Indicator to control whether the graph data is normalized.
	NormalizeFlag string `json:"normalizeFlag"`

	// Options - no documentation
	Options []*SoftLayer_Container_Graph_Option `json:"options"`

	// Plots - no documentation
	Plots []*SoftLayer_Container_Graph_Plot `json:"plots"`

	// BaseUnit - no documentation
	BaseUnit string `json:"baseUnit"`

	// EndDatetime - no documentation
	EndDatetime string `json:"endDatetime"`

	// Image - no documentation
	Image string `json:"image"`
}

func (softlayer_container_graph *SoftLayer_Container_Graph) String() string {
	return "SoftLayer_Container_Graph"
}
