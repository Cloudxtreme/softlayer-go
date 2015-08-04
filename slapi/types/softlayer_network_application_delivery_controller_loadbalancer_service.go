package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Status - <nil>
	Status string `json:"status,omitempty"`

	// Port - <nil>
	Port int `json:"port,omitempty"`

	// Enabled - <nil>
	Enabled int `json:"enabled,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// IpAddressId - <nil>
	IpAddressId int `json:"ipAddressId,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service

	// HealthChecks - <nil>
	HealthChecks []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthChecks,omitempty"`

	// Groups - <nil>
	Groups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"groups,omitempty"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress,omitempty"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount,omitempty"`

	// GroupReferenceCount - no documentation
	GroupReferenceCount uint64 `json:"groupReferenceCount,omitempty"`

	// GroupReferences - <nil>
	GroupReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"groupReferences,omitempty"`

	// HealthCheckCount - no documentation
	HealthCheckCount uint64 `json:"healthCheckCount,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service"
}
