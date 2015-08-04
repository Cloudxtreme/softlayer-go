package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service struct {

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Status - <nil>
	Status string `json:"status,omitempty"`

	// Enabled - <nil>
	Enabled int `json:"enabled,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// IpAddressId - <nil>
	IpAddressId int `json:"ipAddressId,omitempty"`

	// Port - <nil>
	Port int `json:"port,omitempty"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount,omitempty"`

	// Groups - <nil>
	Groups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"groups,omitempty"`

	// GroupReferenceCount - no documentation
	GroupReferenceCount uint64 `json:"groupReferenceCount,omitempty"`

	// HealthChecks - <nil>
	HealthChecks []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthChecks,omitempty"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress,omitempty"`

	// HealthCheckCount - no documentation
	HealthCheckCount uint64 `json:"healthCheckCount,omitempty"`

	// GroupReferences - <nil>
	GroupReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"groupReferences,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service"
}
