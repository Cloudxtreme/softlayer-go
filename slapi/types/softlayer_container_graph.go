package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Graph - <nil>
type SoftLayer_Container_Graph struct {

	// BaseUnit - no documentation
	BaseUnit string `json:"baseUnit,omitempty"`

	// Height - no documentation
	Height int `json:"height,omitempty"`

	// Image - no documentation
	Image string `json:"image,omitempty"`

	// Options - no documentation
	Options []*SoftLayer_Container_Graph_Option `json:"options,omitempty"`

	// EndDatetime - no documentation
	EndDatetime string `json:"endDatetime,omitempty"`

	// ReturnImage - no documentation
	ReturnImage bool `json:"returnImage,omitempty"`

	// Metrics - no documentation
	Metrics []*SoftLayer_Container_Metric_Data_Type `json:"metrics,omitempty"`

	// NormalizeFlag - Indicator to control whether the graph data is normalized.
	NormalizeFlag string `json:"normalizeFlag,omitempty"`

	// StartDatetime - no documentation
	StartDatetime string `json:"startDatetime,omitempty"`

	// Width - no documentation
	Width int `json:"width,omitempty"`

	// Interval - no documentation
	Interval int `json:"interval,omitempty"`

	// Plots - no documentation
	Plots []*SoftLayer_Container_Graph_Plot `json:"plots,omitempty"`

	// Template - no documentation
	Template string `json:"template,omitempty"`

	// Title - no documentation
	Title string `json:"title,omitempty"`
}

func (softlayer_container_graph *SoftLayer_Container_Graph) String() string {
	return "SoftLayer_Container_Graph"
}
