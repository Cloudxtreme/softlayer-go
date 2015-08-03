package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status - SoftLayer customer servers
// that are purchased with the Microsoft Windows operating system are configured by default to retrieve
// updates from SoftLayer's local Windows Server Update Services server. Periodically, these servers
// synchronize and check for new updates from their local server.
// SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status models the results of a server's
// last synchronization attempt as queried from SoftLayer's servers.
type SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status struct {

	// LastRebootDate - The last time a server rebooted due to a Windows Update.
	LastRebootDate *time.Time `json:"lastRebootDate"`

	// LastStatusDate - The last time that SoftLayer's local server received a status update from a
	// customer server.
	LastStatusDate *time.Time `json:"lastStatusDate"`

	// LastSyncDate - The last time a server synchronized with SoftLayer's local server.
	LastSyncDate *time.Time `json:"lastSyncDate"`

	// PrivateIPAddress - no documentation
	PrivateIPAddress string `json:"privateIPAddress"`

	// SyncStatus - The status message returned from a server's last synchronization with SoftLayer's local
	// server.
	SyncStatus string `json:"syncStatus"`

	// UpdateStatus - A server's update status, as retrieved form SoftLayer's local server.
	UpdateStatus string `json:"updateStatus"`
}

func (softlayer_container_utility_microsoft_windows_updateservices_status *SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status) String() string {
	return "SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status"
}
