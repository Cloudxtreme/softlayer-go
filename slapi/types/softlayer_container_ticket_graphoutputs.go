package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Ticket_GraphOutputs - SoftLayer_Container_Ticket_GraphOutputs models a single
// outbound object for a given bandwidth graph.
type SoftLayer_Container_Ticket_GraphOutputs struct {

	// GraphImage - The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage string `json:"graphImage"`

	// GraphTitle - The title that ended up being displayed as part of the graph image.
	GraphTitle string `json:"graphTitle"`

	// MaxEndDate - no documentation
	MaxEndDate *time.Time `json:"maxEndDate"`

	// MinStartDate - no documentation
	MinStartDate *time.Time `json:"minStartDate"`
}
