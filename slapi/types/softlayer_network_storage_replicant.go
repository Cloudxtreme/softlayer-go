package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Replicant - <nil>
type SoftLayer_Network_Storage_Replicant struct {

	// FailbackInProgressFlag - When a replicant is in the process of synchronizing with the parent volume
	// this flag will be true.
	FailbackInProgressFlag string `json:"failbackInProgressFlag"`

	// VolumeName - no documentation
	VolumeName string `json:"volumeName"`
}

func (softlayer_network_storage_replicant *SoftLayer_Network_Storage_Replicant) String() string {
	return "SoftLayer_Network_Storage_Replicant"
}
