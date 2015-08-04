package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail -
// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_File models a CDN account's overall
// bandwidth usage and overages within a given date range.
type SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail struct {

	// ViewCount - no documentation
	ViewCount int `json:"viewCount,omitempty"`

	// Usage - The amount of bandwidth used by a CDN account in between a given starting and ending date.
	Usage float32 `json:"usage,omitempty"`

	// UsageUnits - The unit of measurement used in a CDN bandwidth summary.
	UsageUnits string `json:"usageUnits,omitempty"`

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId,omitempty"`

	// EndDateTime - no documentation
	EndDateTime *time.Time `json:"endDateTime,omitempty"`

	// Duration - no documentation
	Duration float32 `json:"duration,omitempty"`

	// FileName - no documentation
	FileName string `json:"fileName,omitempty"`

	// MediaType - no documentation
	MediaType string `json:"mediaType,omitempty"`

	// StartDateTime - no documentation
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}

func (softlayer_container_network_contentdelivery_bandwidth_summary_detail *SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary_Detail"
}
