package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Storage_Evault_WebCc_BackupResults - The
// SoftLayer_Container_Network_Storage_Evault_WebCc_BackupResults will contain the timeframe of backups
// and the results will also be returned.
type SoftLayer_Container_Network_Storage_Evault_WebCc_BackupResults struct {

	// Failed - no documentation
	Failed string `json:"failed"`

	// Success - no documentation
	Success string `json:"success"`

	// BeginTime - no documentation
	BeginTime *time.Time `json:"beginTime"`

	// Conflict - no documentation
	Conflict string `json:"conflict"`

	// EndTime - no documentation
	EndTime *time.Time `json:"endTime"`
}

func (softlayer_container_network_storage_evault_webcc_backupresults *SoftLayer_Container_Network_Storage_Evault_WebCc_BackupResults) String() string {
	return "SoftLayer_Container_Network_Storage_Evault_WebCc_BackupResults"
}
