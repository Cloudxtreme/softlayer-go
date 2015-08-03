package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Firewall - The SoftLayer_Billing_Item_Network_Firewall data type
// contains general information relating to a single SoftLayer billing item whose item category code is
// 'firewall'
type SoftLayer_Billing_Item_Network_Firewall struct {

	// Resource - The firewall that a firewall billing item is associated with.
	Resource *SoftLayer_Network_Component_Firewall `json:"resource"`
}

func (softlayer_billing_item_network_firewall *SoftLayer_Billing_Item_Network_Firewall) String() string {
	return "SoftLayer_Billing_Item_Network_Firewall"
}
