package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Allowed_Host_VirtualGuest - <nil>
type SoftLayer_Network_Storage_Allowed_Host_VirtualGuest struct {
}

// SoftLayer_Network_Storage_Allowed_Host_VirtualGuest_Extended is SoftLayer_Network_Storage_Allowed_Host_VirtualGuest with all maskable types.
type SoftLayer_Network_Storage_Allowed_Host_VirtualGuest_Extended struct {
	SoftLayer_Network_Storage_Allowed_Host_VirtualGuest

	// Resource - The SoftLayer_Virtual_Guest object which this SoftLayer_Network_Storage_Allowed_Host is
	// referencing.
	Resource *SoftLayer_Virtual_Guest `json:"resource"`
}

func (softlayer_network_storage_allowed_host_virtualguest *SoftLayer_Network_Storage_Allowed_Host_VirtualGuest) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_VirtualGuest"
}
