package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails - The
// SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails will contain basic details for all
// backup and restore jobs performed by the StorageLayer EVault service offering.
type SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails struct {

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId,omitempty"`

	// LastRunDate - no documentation
	LastRunDate *time.Time `json:"lastRunDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Result - no documentation
	Result string `json:"result,omitempty"`

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId,omitempty"`

	// BytesUsed - The number of bytes currently used by the backup job. (provided only for backup jobs)
	BytesUsed uint64 `json:"bytesUsed,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// OriginalSize - Size of backup job when it was first run. (provided only for backup jobs)
	OriginalSize uint64 `json:"originalSize,omitempty"`

	// PercentageOfTotalUsage - Percentage of overall used space allocated by the job. (provided only for
	// backup jobs)
	PercentageOfTotalUsage int `json:"percentageOfTotalUsage,omitempty"`
}

func (softlayer_container_network_storage_evault_webcc_jobdetails *SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails) String() string {
	return "SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails"
}
