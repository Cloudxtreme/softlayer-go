package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type struct {

	// Id - <nil>
	Id int `json:"id"`

	// Keyname - <nil>
	Keyname string `json:"keyname"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_check_type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type"
}

// GetAllObjects - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_health_check_type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error) {
	var returnValue []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_health_check_type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type
	return returnValue, nil
}
