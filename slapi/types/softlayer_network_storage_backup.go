package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Backup - The SoftLayer_Network_Storage_Backup contains general information
// regarding a Storage backup service such as account id, username, maximum capacity, password,
// Storage's product type and the server id.
type SoftLayer_Network_Storage_Backup struct {
}

// SoftLayer_Network_Storage_Backup_Extended is SoftLayer_Network_Storage_Backup with all maskable types.
type SoftLayer_Network_Storage_Backup_Extended struct {
	SoftLayer_Network_Storage_Backup

	// CurrentCyclePeakUsage - Peak number of bytes used in the vault for the current billing cycle.
	CurrentCyclePeakUsage uint `json:"currentCyclePeakUsage"`

	// PreviousCyclePeakUsage - Peak number of bytes used in the vault for the previous billing cycle.
	PreviousCyclePeakUsage uint `json:"previousCyclePeakUsage"`
}

func (softlayer_network_storage_backup *SoftLayer_Network_Storage_Backup) String() string {
	return "SoftLayer_Network_Storage_Backup"
}
