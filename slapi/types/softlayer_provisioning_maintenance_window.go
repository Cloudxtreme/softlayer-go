package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Provisioning_Maintenance_Window - The SoftLayer_Provisioning_Maintenance_Window represent
// a time window that SoftLayer performs a hardware or software maintenance and upgrades.
type SoftLayer_Provisioning_Maintenance_Window struct {

	// BeginDate - The date and time a maintenance window is scheduled to begin.
	BeginDate *time.Time `json:"beginDate,omitempty"`

	// DayOfWeek - An ISO-8601 numeric representation of the day of the week that a maintenance window is
	// performed. 1: Monday, 7: Sunday
	DayOfWeek int `json:"dayOfWeek,omitempty"`

	// EndDate - The date and time a maintenance window is scheduled to end.
	EndDate *time.Time `json:"endDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// LocationId - An internal identifier of the location (data center) record that a maintenance window
	// takes place in.
	LocationId int `json:"locationId,omitempty"`

	// PortalTzId - no documentation
	PortalTzId int `json:"portalTzId,omitempty"`
}

func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) String() string {
	return "SoftLayer_Provisioning_Maintenance_Window"
}
