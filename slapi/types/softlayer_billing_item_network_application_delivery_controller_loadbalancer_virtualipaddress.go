package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress - A
// SoftLayer_Billing_Item_Network_Application_Delivery_Controller_LoadBalancer represents the
// [[SoftLayer_Billing_Item|billing item]] related to a single
// [[SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress|load balancer]]
// instance.
type SoftLayer_Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {

	// Resource - The load balancer that a load balancer billing item is associated with.
	Resource *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"resource"`
}

func (softlayer_billing_item_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Billing_Item_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}
