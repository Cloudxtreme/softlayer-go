package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Allowed_Host_Subnet - <nil>
type SoftLayer_Network_Storage_Allowed_Host_Subnet struct {
}

// SoftLayer_Network_Storage_Allowed_Host_Subnet_Extended is SoftLayer_Network_Storage_Allowed_Host_Subnet with all maskable types.
type SoftLayer_Network_Storage_Allowed_Host_Subnet_Extended struct {
	SoftLayer_Network_Storage_Allowed_Host_Subnet

	// Resource - The SoftLayer_Network_Subnet object which this SoftLayer_Network_Storage_Allowed_Host is
	// referencing.
	Resource *SoftLayer_Network_Subnet `json:"resource"`
}

func (softlayer_network_storage_allowed_host_subnet *SoftLayer_Network_Storage_Allowed_Host_Subnet) String() string {
	return "SoftLayer_Network_Storage_Allowed_Host_Subnet"
}
