package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {

	// RoutingTypeId - <nil>
	RoutingTypeId int `json:"routingTypeId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// Timeout - The timeout value for connections from remote clients to the load balancer. Timeout values
	// are only valid for service groups.
	Timeout int `json:"timeout,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// RoutingMethodId - <nil>
	RoutingMethodId int `json:"routingMethodId,omitempty"`

	// Services - <nil>
	Services []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service `json:"services,omitempty"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount,omitempty"`

	// RoutingType - <nil>
	RoutingType *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type `json:"routingType,omitempty"`

	// ServiceReferences - <nil>
	ServiceReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"serviceReferences,omitempty"`

	// VirtualServers - <nil>
	VirtualServers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServers,omitempty"`

	// VirtualServerCount - no documentation
	VirtualServerCount uint64 `json:"virtualServerCount,omitempty"`

	// RoutingMethod - <nil>
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod,omitempty"`

	// ServiceReferenceCount - no documentation
	ServiceReferenceCount uint64 `json:"serviceReferenceCount,omitempty"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service_group *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group"
}
