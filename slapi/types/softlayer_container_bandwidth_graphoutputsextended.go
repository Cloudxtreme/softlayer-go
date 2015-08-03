package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Bandwidth_GraphOutputsExtended - SoftLayer_Container_Bandwidth_GraphOutputs
// models an individual bandwidth graph image and certain details about that graph image.
type SoftLayer_Container_Bandwidth_GraphOutputsExtended struct {

	// GraphImage - The raw PNG binary data of a bandwidth graph image.
	GraphImage string `json:"graphImage"`

	// GraphTitle - no documentation
	GraphTitle string `json:"graphTitle"`

	// InBoundTotalBytes - The amount of inbound traffic reported on a bandwidth graph image.
	InBoundTotalBytes uint64 `json:"inBoundTotalBytes"`

	// MaxEndDate - The ending date of the data represented in a bandwidth graph.
	MaxEndDate *time.Time `json:"maxEndDate"`

	// MinStartDate - The beginning date of the data represented in a bandwidth graph.
	MinStartDate *time.Time `json:"minStartDate"`

	// OutBoundTotalBytes - The amount of outbound traffic reported on a bandwidth graph image.
	OutBoundTotalBytes uint64 `json:"outBoundTotalBytes"`
}

func (softlayer_container_bandwidth_graphoutputsextended *SoftLayer_Container_Bandwidth_GraphOutputsExtended) String() string {
	return "SoftLayer_Container_Bandwidth_GraphOutputsExtended"
}
