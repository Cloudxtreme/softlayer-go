package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Hub_Swift - <nil>
type SoftLayer_Network_Storage_Hub_Swift struct {
}

// SoftLayer_Network_Storage_Hub_Swift_Extended is SoftLayer_Network_Storage_Hub_Swift with all maskable types.
type SoftLayer_Network_Storage_Hub_Swift_Extended struct {
	SoftLayer_Network_Storage_Hub_Swift

	// StorageNodes - <nil>
	StorageNodes []*SoftLayer_Network_Service_Resource `json:"storageNodes"`

	// StorageNodeCount - no documentation
	StorageNodeCount uint64 `json:"storageNodeCount"`
}

func (softlayer_network_storage_hub_swift *SoftLayer_Network_Storage_Hub_Swift) String() string {
	return "SoftLayer_Network_Storage_Hub_Swift"
}
