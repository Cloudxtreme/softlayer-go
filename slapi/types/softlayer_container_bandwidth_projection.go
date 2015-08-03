package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Bandwidth_Projection - SoftLayer_Container_Bandwidth_Projection models projected
// bandwidth use over a time range.
type SoftLayer_Container_Bandwidth_Projection struct {

	// AllowedUsage - no documentation
	AllowedUsage string `json:"allowedUsage"`

	// EstimatedUsage - Estimated bandwidth usage so far this billing cycle.
	EstimatedUsage string `json:"estimatedUsage"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`

	// ProjectedUsage - Projected usage for this hardware based on previous usage this billing cycle.
	ProjectedUsage string `json:"projectedUsage"`

	// ServerName - no documentation
	ServerName string `json:"serverName"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_container_bandwidth_projection *SoftLayer_Container_Bandwidth_Projection) String() string {
	return "SoftLayer_Container_Bandwidth_Projection"
}
