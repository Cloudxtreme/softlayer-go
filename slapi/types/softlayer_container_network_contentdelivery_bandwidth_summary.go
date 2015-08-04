package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary -
// SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary models a CDN account's overall
// bandwidth usage and overages within a given date range.
type SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary struct {

	// StartDateTime - no documentation
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// Usage - The amount of bandwidth used by a CDN account in between a given starting and ending date.
	Usage slapi.Float64 `json:"usage,omitempty"`

	// UsageUnits - The unit of measurement used in a CDN bandwidth summary.
	UsageUnits string `json:"usageUnits,omitempty"`

	// CdnAccountId - no documentation
	CdnAccountId int `json:"cdnAccountId,omitempty"`

	// EndDateTime - no documentation
	EndDateTime *time.Time `json:"endDateTime,omitempty"`

	// FileName - no documentation
	FileName string `json:"fileName,omitempty"`

	// MediaType - no documentation
	MediaType string `json:"mediaType,omitempty"`
}

func (softlayer_container_network_contentdelivery_bandwidth_summary *SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary"
}
