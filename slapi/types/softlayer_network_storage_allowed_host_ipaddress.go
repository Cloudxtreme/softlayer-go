package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Allowed_Host_IpAddress - <nil>
type SoftLayer_Network_Storage_Allowed_Host_IpAddress struct {
}

// SoftLayer_Network_Storage_Allowed_Host_IpAddress_Extended is SoftLayer_Network_Storage_Allowed_Host_IpAddress with all maskable types.
type SoftLayer_Network_Storage_Allowed_Host_IpAddress_Extended struct {
	SoftLayer_Network_Storage_Allowed_Host_IpAddress

	// Resource - The SoftLayer_Network_Subnet_IpAddress object which this
	// SoftLayer_Network_Storage_Allowed_Host is referencing.
	Resource *SoftLayer_Network_Subnet_IpAddress `json:"resource"`
}

func (softlayer_network_storage_allowed_host_ipaddress *SoftLayer_Network_Storage_Allowed_Host_IpAddress) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_IpAddress"
}
