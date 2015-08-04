package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// HealthCheckTypeId - <nil>
	HealthCheckTypeId int `json:"healthCheckTypeId,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_check *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// ScaleLoadBalancerCount - A count of collection of scale load balancers that use this health check.
	ScaleLoadBalancerCount uint64 `json:"scaleLoadBalancerCount,omitempty"`

	// ScaleLoadBalancers - Collection of scale load balancers that use this health check.
	ScaleLoadBalancers []*SoftLayer_Scale_LoadBalancer `json:"scaleLoadBalancers,omitempty"`

	// Services - <nil>
	Services []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service `json:"services,omitempty"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Type `json:"type,omitempty"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount,omitempty"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute `json:"attributes,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_check *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check"
}
