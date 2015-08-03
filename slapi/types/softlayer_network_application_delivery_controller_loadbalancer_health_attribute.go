package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {

	// HealthAttributeTypeId - <nil>
	HealthAttributeTypeId int `json:"healthAttributeTypeId"`

	// HealthCheck - <nil>
	HealthCheck *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthCheck"`

	// HealthCheckId - <nil>
	HealthCheckId int `json:"healthCheckId"`

	// Id - <nil>
	Id int `json:"id"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type `json:"type"`

	// Value - <nil>
	Value string `json:"value"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_attribute *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute"
}
