package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails - The
// SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails will contain basic details for all
// backup and restore jobs performed by the StorageLayer EVault service offering.
type SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails struct {

	// BytesUsed - The number of bytes currently used by the backup job. (provided only for backup jobs)
	BytesUsed uint64 `json:"bytesUsed"`

	// Description - no documentation
	Description string `json:"description"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`

	// LastRunDate - no documentation
	LastRunDate *time.Time `json:"lastRunDate"`

	// Name - no documentation
	Name string `json:"name"`

	// OriginalSize - Size of backup job when it was first run. (provided only for backup jobs)
	OriginalSize uint64 `json:"originalSize"`

	// PercentageOfTotalUsage - Percentage of overall used space allocated by the job. (provided only for
	// backup jobs)
	PercentageOfTotalUsage int `json:"percentageOfTotalUsage"`

	// Result - no documentation
	Result string `json:"result"`

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId"`
}

func (softlayer_container_network_storage_evault_webcc_jobdetails *SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails) String() string {
	return "SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails"
}
