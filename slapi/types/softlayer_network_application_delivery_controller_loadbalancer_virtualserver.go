package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {

	// Id - <nil>
	Id int `json:"id"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// RoutingMethodId - <nil>
	RoutingMethodId int `json:"routingMethodId"`

	// Allocation - <nil>
	Allocation int `json:"allocation"`

	// Name - <nil>
	Name string `json:"name"`

	// Port - <nil>
	Port int `json:"port"`

	// VirtualIpAddressId - <nil>
	VirtualIpAddressId int `json:"virtualIpAddressId"`
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer

	// RoutingMethod - <nil>
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod"`

	// ScaleLoadBalancers - Collection of scale load balancers this virtual server applies to.
	ScaleLoadBalancers []*SoftLayer_Scale_LoadBalancer `json:"scaleLoadBalancers"`

	// VirtualIpAddress - <nil>
	VirtualIpAddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress"`

	// ScaleLoadBalancerCount - A count of collection of scale load balancers this virtual server applies
	// to.
	ScaleLoadBalancerCount uint64 `json:"scaleLoadBalancerCount"`

	// ServiceGroupCount - no documentation
	ServiceGroupCount uint64 `json:"serviceGroupCount"`

	// ServiceGroups - <nil>
	ServiceGroups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"serviceGroups"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer"
}
