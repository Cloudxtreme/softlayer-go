package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus - The
// SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus will contain the timestamp of the last
// backup performed by the EVault agent. The agent status will also be returned.
type SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus struct {

	// LastBackup - Timestamp of last backup performed by the EVault backup agent
	LastBackup *time.Time `json:"lastBackup"`

	// Status - Status indicating the accumulative status result of all jobs performed by the evault agent.
	// For example, if one job out three jobs failed agent status will by "Failed Backup(s)".
	Status string `json:"status"`
}
