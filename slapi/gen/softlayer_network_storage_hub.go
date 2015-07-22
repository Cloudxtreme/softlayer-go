package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Hub - The SoftLayer_Network_Storage_Hub data type models Virtual Server
// type Storage storage offerings.
type SoftLayer_Network_Storage_Hub struct {

	// BandwidthBillingItemCount - A count of the billing items tied to a Storage service's bandwidth
	// usage.
	BandwidthBillingItemCount uint64 `json:"bandwidthBillingItemCount"`

	// BandwidthBillingItems - The billing items tied to a Storage service's bandwidth usage.
	BandwidthBillingItems []*SoftLayer_Billing_Item `json:"bandwidthBillingItems"`
}
