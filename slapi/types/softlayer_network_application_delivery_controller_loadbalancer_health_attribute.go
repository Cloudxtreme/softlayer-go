package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {

	// HealthAttributeTypeId - <nil>
	HealthAttributeTypeId int `json:"healthAttributeTypeId,omitempty"`

	// HealthCheckId - <nil>
	HealthCheckId int `json:"healthCheckId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Value - <nil>
	Value string `json:"value,omitempty"`

	// HealthCheck - <nil>
	HealthCheck *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthCheck,omitempty"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type `json:"type,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_attribute *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute"
}
