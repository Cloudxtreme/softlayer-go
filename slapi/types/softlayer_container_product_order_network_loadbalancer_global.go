package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Network_LoadBalancer_Global - This is the datatype that needs to
// be populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required
// to place a global load balancer order with SoftLayer.
type SoftLayer_Container_Product_Order_Network_LoadBalancer_Global struct {

	// Domain - no documentation
	Domain string `json:"domain"`

	// Hostname - no documentation
	Hostname string `json:"hostname"`
}

func (softlayer_container_product_order_network_loadbalancer_global *SoftLayer_Container_Product_Order_Network_LoadBalancer_Global) String() string {
	return "SoftLayer_Container_Product_Order_Network_LoadBalancer_Global"
}
