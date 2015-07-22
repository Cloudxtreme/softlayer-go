package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_LoadBalancer_VirtualIpAddress - A
// SoftLayer_Billing_Item_Network_LoadBalancer_VirtualIpAddress represents the
// [[SoftLayer_Billing_Item|billing item]] related to a single
// [[SoftLayer_Network_LoadBalancer_VirtualIpAddress|load balancer]] instance.
type SoftLayer_Billing_Item_Network_LoadBalancer_VirtualIpAddress struct {

	// Resource - The load balancer's virtual IP address that the billing item is associated with.
	Resource *SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"resource"`
}
