package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service struct {

	// Status - <nil>
	Status string `json:"status"`

	// Enabled - <nil>
	Enabled int `json:"enabled"`

	// Id - <nil>
	Id int `json:"id"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Port - <nil>
	Port int `json:"port"`

	// IpAddressId - <nil>
	IpAddressId int `json:"ipAddressId"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service

	// GroupReferences - <nil>
	GroupReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"groupReferences"`

	// Groups - <nil>
	Groups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"groups"`

	// HealthChecks - <nil>
	HealthChecks []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthChecks"`

	// HealthCheckCount - no documentation
	HealthCheckCount uint64 `json:"healthCheckCount"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount"`

	// GroupReferenceCount - no documentation
	GroupReferenceCount uint64 `json:"groupReferenceCount"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service"
}
