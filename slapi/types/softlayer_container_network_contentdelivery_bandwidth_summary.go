package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary -
// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary models a CDN account's overall
// bandwidth usage and overages within a given date range.
type SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary struct {

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId"`

	// EndDateTime - no documentation
	EndDateTime *time.Time `json:"endDateTime"`

	// FileName - no documentation
	FileName string `json:"fileName"`

	// MediaType - no documentation
	MediaType string `json:"mediaType"`

	// StartDateTime - no documentation
	StartDateTime *time.Time `json:"startDateTime"`

	// Usage - The amount of bandwidth used by a CDN account in between a given starting and ending date.
	Usage float32 `json:"usage"`

	// UsageUnits - The unit of measurement used in a CDN bandwidth summary.
	UsageUnits string `json:"usageUnits"`
}