package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// RoutingMethod - <nil>
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod"`

	// RoutingMethodId - <nil>
	RoutingMethodId int `json:"routingMethodId"`

	// RoutingType - <nil>
	RoutingType *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type `json:"routingType"`

	// RoutingTypeId - <nil>
	RoutingTypeId int `json:"routingTypeId"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount"`

	// ServiceReferenceCount - no documentation
	ServiceReferenceCount uint64 `json:"serviceReferenceCount"`

	// ServiceReferences - <nil>
	ServiceReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"serviceReferences"`

	// Services - <nil>
	Services []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service `json:"services"`

	// Timeout - The timeout value for connections from remote clients to the load balancer. Timeout values
	// are only valid for service groups.
	Timeout int `json:"timeout"`

	// VirtualServerCount - no documentation
	VirtualServerCount uint64 `json:"virtualServerCount"`

	// VirtualServers - <nil>
	VirtualServers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServers"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service_group *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group"
}
