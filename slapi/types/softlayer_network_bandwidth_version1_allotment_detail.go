package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Bandwidth_Version1_Allotment_Detail - The
// SoftLayer_Network_Bandwidth_Version1_Allotment_Detail data type contains specific information
// relating to a single bandwidth allotment record.
type SoftLayer_Network_Bandwidth_Version1_Allotment_Detail struct {

	// AllocationId - no documentation
	AllocationId int `json:"allocationId,omitempty"`

	// BandwidthAllotmentId - no documentation
	BandwidthAllotmentId int `json:"bandwidthAllotmentId,omitempty"`

	// EffectiveDate - Beginning this date the bandwidth allotment is active.
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// EndEffectiveDate - From this date the bandwidth allotment is no longer active.
	EndEffectiveDate *time.Time `json:"endEffectiveDate,omitempty"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_network_bandwidth_version1_allotment_detail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Allotment_Detail"
}

// SoftLayer_Network_Bandwidth_Version1_Allotment_Detail_Extended is SoftLayer_Network_Bandwidth_Version1_Allotment_Detail with all maskable types.
type SoftLayer_Network_Bandwidth_Version1_Allotment_Detail_Extended struct {
	SoftLayer_Network_Bandwidth_Version1_Allotment_Detail

	// BandwidthUsageCount - no documentation
	BandwidthUsageCount uint64 `json:"bandwidthUsageCount,omitempty"`

	// BandwidthAllotment - no documentation
	BandwidthAllotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotment,omitempty"`

	// Allocation - no documentation
	Allocation *SoftLayer_Network_Bandwidth_Version1_Allocation `json:"allocation,omitempty"`

	// BandwidthUsage - no documentation
	BandwidthUsage []*SoftLayer_Network_Bandwidth_Version1_Usage `json:"bandwidthUsage,omitempty"`
}

func (softlayer_network_bandwidth_version1_allotment_detail *SoftLayer_Network_Bandwidth_Version1_Allotment_Detail_Extended) String() string {
	return "SoftLayer_Network_Bandwidth_Version1_Allotment_Detail"
}
