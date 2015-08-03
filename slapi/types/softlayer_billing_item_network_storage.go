package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Storage - The SoftLayer_Billing_Item_Network_Storage data type
// describes the billing items related to StorageLayer accounts.
type SoftLayer_Billing_Item_Network_Storage struct {
}

func (softlayer_billing_item_network_storage *SoftLayer_Billing_Item_Network_Storage) String() string {
	return "SoftLayer_Billing_Item_Network_Storage"
}

// SoftLayer_Billing_Item_Network_Storage_Extended is SoftLayer_Billing_Item_Network_Storage with all maskable types.
type SoftLayer_Billing_Item_Network_Storage_Extended struct {
	SoftLayer_Billing_Item_Network_Storage

	// Resource - The StorageLayer account that a network storage billing item is associated with.
	Resource *SoftLayer_Network_Storage `json:"resource"`
}

func (softlayer_billing_item_network_storage *SoftLayer_Billing_Item_Network_Storage_Extended) String() string {
	return "SoftLayer_Billing_Item_Network_Storage"
}
