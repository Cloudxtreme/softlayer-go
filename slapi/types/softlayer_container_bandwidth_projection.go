package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Bandwidth_Projection - SoftLayer_Container_Bandwidth_Projection models projected
// bandwidth use over a time range.
type SoftLayer_Container_Bandwidth_Projection struct {

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// AllowedUsage - no documentation
	AllowedUsage string `json:"allowedUsage,omitempty"`

	// EstimatedUsage - Estimated bandwidth usage so far this billing cycle.
	EstimatedUsage string `json:"estimatedUsage,omitempty"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId,omitempty"`

	// ProjectedUsage - Projected usage for this hardware based on previous usage this billing cycle.
	ProjectedUsage string `json:"projectedUsage,omitempty"`

	// ServerName - no documentation
	ServerName string `json:"serverName,omitempty"`
}

func (softlayer_container_bandwidth_projection *SoftLayer_Container_Bandwidth_Projection) String() string {
	return "SoftLayer_Container_Bandwidth_Projection"
}
