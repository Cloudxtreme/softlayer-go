package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Bandwidth_Version1_Allotment_Detail - The
// SoftLayer_Network_Bandwidth_Version1_Allotment_Detail data type contains specific information
// relating to a single bandwidth allotment record.
type SoftLayer_Network_Bandwidth_Version1_Allotment_Detail struct {

	// Allocation - no documentation
	Allocation *SoftLayer_Network_Bandwidth_Version1_Allocation `json:"allocation"`

	// AllocationId - no documentation
	AllocationId int `json:"allocationId"`

	// BandwidthAllotment - no documentation
	BandwidthAllotment *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotment"`

	// BandwidthAllotmentId - no documentation
	BandwidthAllotmentId int `json:"bandwidthAllotmentId"`

	// BandwidthUsage - no documentation
	BandwidthUsage []*SoftLayer_Network_Bandwidth_Version1_Usage `json:"bandwidthUsage"`

	// BandwidthUsageCount - no documentation
	BandwidthUsageCount uint64 `json:"bandwidthUsageCount"`

	// EffectiveDate - Beginning this date the bandwidth allotment is active.
	EffectiveDate *time.Time `json:"effectiveDate"`

	// EndEffectiveDate - From this date the bandwidth allotment is no longer active.
	EndEffectiveDate *time.Time `json:"endEffectiveDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId"`
}
