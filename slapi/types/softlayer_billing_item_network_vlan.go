package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Vlan - The SoftLayer_Billing_Item_Network_Vlant data type contains
// general information relating to a single SoftLayer billing item whose item category code is one of
// the following: * network_vlan These item categories denote that the billing item has network vlan
// information attached.
type SoftLayer_Billing_Item_Network_Vlan struct {

	// Resource - The resource for a network vlan related billing item.
	Resource *SoftLayer_Network_Vlan `json:"resource"`
}

func (softlayer_billing_item_network_vlan *SoftLayer_Billing_Item_Network_Vlan) String() string {
	return "SoftLayer_Billing_Item_Network_Vlan"
}
