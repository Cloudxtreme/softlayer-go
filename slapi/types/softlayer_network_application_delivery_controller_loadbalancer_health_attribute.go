package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute struct {

	// HealthAttributeTypeId - <nil>
	HealthAttributeTypeId int `json:"healthAttributeTypeId"`

	// HealthCheckId - <nil>
	HealthCheckId int `json:"healthCheckId"`

	// Id - <nil>
	Id int `json:"id"`

	// Value - <nil>
	Value string `json:"value"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_attribute *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute

	// HealthCheck - <nil>
	HealthCheck *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthCheck"`

	// Type - <nil>
	Type *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Type `json:"type"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_health_attribute *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Attribute"
}
