package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Hub - The SoftLayer_Network_Storage_Hub data type models Virtual Server
// type Storage storage offerings.
type SoftLayer_Network_Storage_Hub struct {
}

func (softlayer_network_storage_hub *SoftLayer_Network_Storage_Hub) String() string {
	return "SoftLayer_Network_Storage_Hub"
}

// SoftLayer_Network_Storage_Hub_Extended is SoftLayer_Network_Storage_Hub with all maskable types.
type SoftLayer_Network_Storage_Hub_Extended struct {
	SoftLayer_Network_Storage_Hub

	// BandwidthBillingItems - The billing items tied to a Storage service's bandwidth usage.
	BandwidthBillingItems []*SoftLayer_Billing_Item `json:"bandwidthBillingItems"`

	// BandwidthBillingItemCount - A count of the billing items tied to a Storage service's bandwidth
	// usage.
	BandwidthBillingItemCount uint64 `json:"bandwidthBillingItemCount"`
}

func (softlayer_network_storage_hub *SoftLayer_Network_Storage_Hub_Extended) String() string {
	return "SoftLayer_Network_Storage_Hub"
}
