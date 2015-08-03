package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Nas - The SoftLayer_Network_Storage_Nas contains general information
// regarding a NAS Storage service such as account id, username, password, maximum capacity, Storage's
// product type and capacity.
type SoftLayer_Network_Storage_Nas struct {
}

func (softlayer_network_storage_nas *SoftLayer_Network_Storage_Nas) String() string {
	return "SoftLayer_Network_Storage_Nas"
}

// SoftLayer_Network_Storage_Nas_Extended is SoftLayer_Network_Storage_Nas with all maskable types.
type SoftLayer_Network_Storage_Nas_Extended struct {
	SoftLayer_Network_Storage_Nas

	// RecentBytesUsed - <nil>
	RecentBytesUsed *SoftLayer_Network_Storage_Daily_Usage `json:"recentBytesUsed"`
}

func (softlayer_network_storage_nas *SoftLayer_Network_Storage_Nas_Extended) String() string {
	return "SoftLayer_Network_Storage_Nas"
}
