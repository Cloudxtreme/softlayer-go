package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_LoadBalancer_Global - The
// SoftLayer_Billing_Item_Network_LoadBalancer_Global data type contains general information relating
// to a single SoftLayer billing item whose item category code is 'global_load_balancer'
type SoftLayer_Billing_Item_Network_LoadBalancer_Global struct {
}

// SoftLayer_Billing_Item_Network_LoadBalancer_Global_Extended is SoftLayer_Billing_Item_Network_LoadBalancer_Global with all maskable types.
type SoftLayer_Billing_Item_Network_LoadBalancer_Global_Extended struct {
	SoftLayer_Billing_Item_Network_LoadBalancer_Global

	// Resource - The resource for a global load balancer billing item.
	Resource *SoftLayer_Network_LoadBalancer_Global_Account `json:"resource"`
}

func (softlayer_billing_item_network_loadbalancer_global *SoftLayer_Billing_Item_Network_LoadBalancer_Global) String() string {
	return "SoftLayer_Billing_Item_Network_LoadBalancer_Global"
}
