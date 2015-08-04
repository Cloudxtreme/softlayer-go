package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {

	// Allocation - <nil>
	Allocation int `json:"allocation,omitempty"`

	// VirtualIpAddressId - <nil>
	VirtualIpAddressId int `json:"virtualIpAddressId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// Port - <nil>
	Port int `json:"port,omitempty"`

	// RoutingMethodId - <nil>
	RoutingMethodId int `json:"routingMethodId,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer

	// ScaleLoadBalancerCount - A count of collection of scale load balancers this virtual server applies
	// to.
	ScaleLoadBalancerCount uint64 `json:"scaleLoadBalancerCount,omitempty"`

	// ScaleLoadBalancers - Collection of scale load balancers this virtual server applies to.
	ScaleLoadBalancers []*SoftLayer_Scale_LoadBalancer `json:"scaleLoadBalancers,omitempty"`

	// ServiceGroupCount - no documentation
	ServiceGroupCount uint64 `json:"serviceGroupCount,omitempty"`

	// RoutingMethod - <nil>
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod,omitempty"`

	// ServiceGroups - <nil>
	ServiceGroups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"serviceGroups,omitempty"`

	// VirtualIpAddress - <nil>
	VirtualIpAddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer"
}
