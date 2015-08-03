package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Snapshot - <nil>
type SoftLayer_Network_Storage_Snapshot struct {
}

// SoftLayer_Network_Storage_Snapshot_Extended is SoftLayer_Network_Storage_Snapshot with all maskable types.
type SoftLayer_Network_Storage_Snapshot_Extended struct {
	SoftLayer_Network_Storage_Snapshot

	// VolumeName - no documentation
	VolumeName string `json:"volumeName"`

	// CreationSchedule - If applicable, the schedule which was executed to create a snapshot.
	CreationSchedule *SoftLayer_Network_Storage_Schedule `json:"creationSchedule"`
}

func (softlayer_network_storage_snapshot *SoftLayer_Network_Storage_Snapshot) String() string {
	return "SoftLayer_Network_Storage_Snapshot"
}
