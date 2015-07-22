package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer struct {

	// Allocation - <nil>
	Allocation int `json:"allocation"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Port - <nil>
	Port int `json:"port"`

	// RoutingMethod - <nil>
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod"`

	// RoutingMethodId - <nil>
	RoutingMethodId int `json:"routingMethodId"`

	// ScaleLoadBalancerCount - A count of collection of scale load balancers this virtual server applies
	// to.
	ScaleLoadBalancerCount uint64 `json:"scaleLoadBalancerCount"`

	// ScaleLoadBalancers - Collection of scale load balancers this virtual server applies to.
	ScaleLoadBalancers []*SoftLayer_Scale_LoadBalancer `json:"scaleLoadBalancers"`

	// ServiceGroupCount - no documentation
	ServiceGroupCount uint64 `json:"serviceGroupCount"`

	// ServiceGroups - <nil>
	ServiceGroups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"serviceGroups"`

	// VirtualIpAddress - <nil>
	VirtualIpAddress *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"virtualIpAddress"`

	// VirtualIpAddressId - <nil>
	VirtualIpAddressId int `json:"virtualIpAddressId"`
}

// DeleteObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer
	return returnValue, nil
}

// StartSsl - Start SSL acceleration on all SSL virtual services (those with a type of This action
// should be taken only after configuring an SSL certificate for the virtual IP.
func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) StartSsl(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// StopSsl - Stop SSL acceleration on all SSL virtual services (those with a type of
func (softlayer_network_application_delivery_controller_loadbalancer_virtualserver *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer) StopSsl(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
