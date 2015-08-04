package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Ticket_GraphOutputs - SoftLayer_Container_Ticket_GraphOutputs models a single
// outbound object for a given bandwidth graph.
type SoftLayer_Container_Ticket_GraphOutputs struct {

	// MaxEndDate - no documentation
	MaxEndDate *time.Time `json:"maxEndDate,omitempty"`

	// MinStartDate - no documentation
	MinStartDate *time.Time `json:"minStartDate,omitempty"`

	// GraphImage - The raw PNG binary data to be displayed once the graph is drawn.
	GraphImage string `json:"graphImage,omitempty"`

	// GraphTitle - The title that ended up being displayed as part of the graph image.
	GraphTitle string `json:"graphTitle,omitempty"`
}

func (softlayer_container_ticket_graphoutputs *SoftLayer_Container_Ticket_GraphOutputs) String() string {
	return "SoftLayer_Container_Ticket_GraphOutputs"
}
