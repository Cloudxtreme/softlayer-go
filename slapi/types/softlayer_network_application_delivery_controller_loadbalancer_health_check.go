package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check struct {

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute `json:"attributes"`

	// HealthCheckTypeId - <nil>
	HealthCheckTypeId int `json:"healthCheckTypeId"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// ScaleLoadBalancerCount - A count of collection of scale load balancers that use this health check.
	ScaleLoadBalancerCount uint64 `json:"scaleLoadBalancerCount"`

	// ScaleLoadBalancers - Collection of scale load balancers that use this health check.
	ScaleLoadBalancers []*SoftLayer_Scale_LoadBalancer `json:"scaleLoadBalancers"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount"`

	// Services - <nil>
	Services []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service `json:"services"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type `json:"type"`
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_health_check *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check
	return returnValue, nil
}
