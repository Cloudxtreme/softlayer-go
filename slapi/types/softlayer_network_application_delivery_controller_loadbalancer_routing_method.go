package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method struct {

	// Id - <nil>
	Id int `json:"id"`

	// Keyname - <nil>
	Keyname string `json:"keyname"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_routing_method *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method"
}

// GetAllObjects - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_routing_method *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error) {
	var returnValue []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_routing_method *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method
	return returnValue, nil
}
