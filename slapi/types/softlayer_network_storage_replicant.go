package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Replicant - <nil>
type SoftLayer_Network_Storage_Replicant struct {
}

func (softlayer_network_storage_replicant *SoftLayer_Network_Storage_Replicant) String() string {
	return "SoftLayer_Network_Storage_Replicant"
}

// SoftLayer_Network_Storage_Replicant_Extended is SoftLayer_Network_Storage_Replicant with all maskable types.
type SoftLayer_Network_Storage_Replicant_Extended struct {
	SoftLayer_Network_Storage_Replicant

	// FailbackInProgressFlag - When a replicant is in the process of synchronizing with the parent volume
	// this flag will be true.
	FailbackInProgressFlag string `json:"failbackInProgressFlag"`

	// VolumeName - no documentation
	VolumeName string `json:"volumeName"`
}

func (softlayer_network_storage_replicant *SoftLayer_Network_Storage_Replicant_Extended) String() string {
	return "SoftLayer_Network_Storage_Replicant"
}
