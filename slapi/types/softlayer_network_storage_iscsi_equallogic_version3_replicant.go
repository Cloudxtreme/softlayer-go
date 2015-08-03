package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant - An iscsi replicant receives incoming
// data from an associated iscsi volume. While the replicant is not in failover mode it will not be
// mountable. Upon failover the replicant can be mounted and used as a normal volume. It is suggested
// to only do this as part of a disaster recovery plan.
type SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant struct {

	// FailbackInProgressFlag - When a replicant is in the process of synchronizing with the parent volume
	// this flag will be true.
	FailbackInProgressFlag bool `json:"failbackInProgressFlag"`

	// VolumeName - no documentation
	VolumeName string `json:"volumeName"`
}

func (softlayer_network_storage_iscsi_equallogic_version3_replicant *SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant) String() string {
	return "SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant"
}
