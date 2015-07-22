package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service struct {

	// Enabled - <nil>
	Enabled int `json:"enabled"`

	// GroupCount - no documentation
	GroupCount uint64 `json:"groupCount"`

	// GroupReferenceCount - no documentation
	GroupReferenceCount uint64 `json:"groupReferenceCount"`

	// GroupReferences - <nil>
	GroupReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"groupReferences"`

	// Groups - <nil>
	Groups []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"groups"`

	// HealthCheckCount - no documentation
	HealthCheckCount uint64 `json:"healthCheckCount"`

	// HealthChecks - <nil>
	HealthChecks []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Health_Check `json:"healthChecks"`

	// Id - <nil>
	Id int `json:"id"`

	// IpAddress - <nil>
	IpAddress *SoftLayer_Network_Subnet_IpAddress `json:"ipAddress"`

	// IpAddressId - <nil>
	IpAddressId int `json:"ipAddressId"`

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Port - <nil>
	Port int `json:"port"`

	// Status - <nil>
	Status string `json:"status"`
}

// DeleteObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetGraphImage - Get the graph image for a load balancer service based on the supplied graph type and
// metric. The available graph types are: 'connections' and 'status', and the available metrics are:
// 'day', 'week' and 'month'. This method returns the raw binary image data.
func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) GetGraphImage(graphType string, metric string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) GetObject() (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service
	return returnValue, nil
}

// ToggleStatus - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service) ToggleStatus() (bool, error) {
	var returnValue bool
	return returnValue, nil
}
