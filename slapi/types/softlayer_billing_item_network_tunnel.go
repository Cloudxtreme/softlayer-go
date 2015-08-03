package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Tunnel - The SoftLayer_Billing_Item_Network_Storage data type
// describes the billing items related to StorageLayer accounts.
type SoftLayer_Billing_Item_Network_Tunnel struct {
}

// SoftLayer_Billing_Item_Network_Tunnel_Extended is SoftLayer_Billing_Item_Network_Tunnel with all maskable types.
type SoftLayer_Billing_Item_Network_Tunnel_Extended struct {
	SoftLayer_Billing_Item_Network_Tunnel

	// Resource - The IPsec VPN that a network tunnel billing item is associated with.
	Resource *SoftLayer_Network_Tunnel_Module_Context `json:"resource"`
}

func (softlayer_billing_item_network_tunnel *SoftLayer_Billing_Item_Network_Tunnel) String() string {
	return "SoftLayer_Billing_Item_Network_Tunnel"
}
