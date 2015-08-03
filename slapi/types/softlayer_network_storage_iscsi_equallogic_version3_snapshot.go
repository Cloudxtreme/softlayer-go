package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Snapshot - An iscsi snapshot is a point-in-time
// view of the data on an associated iscsi volume. Iscsi snapshots use a copy-on-write technology to
// minimize the amount of snapshot space used. When a snapshot is initially created it will use no
// snapshot space. At the time data changes on a volume which existed when a snapshot was created the
// original data will be saved in the associated volume's snapshot reserve space. As a snapshot is
// created offline it must be set mountable in order to mount it via an iscsi initiator service.
type SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Snapshot struct {

	// CreationSchedule - If applicable, the schedule which was executed to create a snapshot.
	CreationSchedule *SoftLayer_Network_Storage_Schedule `json:"creationSchedule"`

	// VolumeName - no documentation
	VolumeName string `json:"volumeName"`
}

func (softlayer_network_storage_iscsi_equallogic_version3_snapshot *SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Snapshot) String() string {
	return "SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Snapshot"
}
