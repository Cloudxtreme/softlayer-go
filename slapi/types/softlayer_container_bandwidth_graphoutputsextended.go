package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Bandwidth_GraphOutputsExtended - SoftLayer_Container_Bandwidth_GraphOutputs
// models an individual bandwidth graph image and certain details about that graph image.
type SoftLayer_Container_Bandwidth_GraphOutputsExtended struct {

	// MaxEndDate - The ending date of the data represented in a bandwidth graph.
	MaxEndDate *time.Time `json:"maxEndDate,omitempty"`

	// MinStartDate - The beginning date of the data represented in a bandwidth graph.
	MinStartDate *time.Time `json:"minStartDate,omitempty"`

	// OutBoundTotalBytes - The amount of outbound traffic reported on a bandwidth graph image.
	OutBoundTotalBytes uint64 `json:"outBoundTotalBytes,omitempty"`

	// GraphImage - The raw PNG binary data of a bandwidth graph image.
	GraphImage string `json:"graphImage,omitempty"`

	// GraphTitle - no documentation
	GraphTitle string `json:"graphTitle,omitempty"`

	// InBoundTotalBytes - The amount of inbound traffic reported on a bandwidth graph image.
	InBoundTotalBytes uint64 `json:"inBoundTotalBytes,omitempty"`
}

func (softlayer_container_bandwidth_graphoutputsextended *SoftLayer_Container_Bandwidth_GraphOutputsExtended) String() string {
	return "SoftLayer_Container_Bandwidth_GraphOutputsExtended"
}
