package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary -
// SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary models an individual
// CDN point of presence's bandwidth usage for a CDN account within a given date range. CDN POPs are
// located throughout the world, so individual POP usage may be beneficial in determining who is
// downloading your CDN hosted content.
type SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary struct {

	// Bandwidth - no documentation
	Bandwidth uint64 `json:"bandwidth,omitempty"`

	// EndDateTime - no documentation
	EndDateTime *time.Time `json:"endDateTime,omitempty"`

	// PopName - A CDN POP's name. This is typically the city the POP resides in.
	PopName string `json:"popName,omitempty"`

	// StartDateTime - no documentation
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// UsageUnits - The unit of measurement used in a CDN POP bandwidth summary.
	UsageUnits string `json:"usageUnits,omitempty"`

	// ViewCount - no documentation
	ViewCount uint64 `json:"viewCount,omitempty"`
}

func (softlayer_container_network_contentdelivery_bandwidth_pointsofpresence_summary *SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary"
}
