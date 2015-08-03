package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Allowed_Host_Hardware - <nil>
type SoftLayer_Network_Storage_Allowed_Host_Hardware struct {
}

// SoftLayer_Network_Storage_Allowed_Host_Hardware_Extended is SoftLayer_Network_Storage_Allowed_Host_Hardware with all maskable types.
type SoftLayer_Network_Storage_Allowed_Host_Hardware_Extended struct {
	SoftLayer_Network_Storage_Allowed_Host_Hardware

	// Resource - The SoftLayer_Hardware object which this SoftLayer_Network_Storage_Allowed_Host is
	// referencing.
	Resource *SoftLayer_Hardware `json:"resource"`
}

func (softlayer_network_storage_allowed_host_hardware *SoftLayer_Network_Storage_Allowed_Host_Hardware) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_Hardware"
}
