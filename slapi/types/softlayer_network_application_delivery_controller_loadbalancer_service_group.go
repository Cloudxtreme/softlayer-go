package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group struct {

	// RoutingTypeId - <nil>
	RoutingTypeId int `json:"routingTypeId"`

	// Name - <nil>
	Name string `json:"name"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Timeout - The timeout value for connections from remote clients to the load balancer. Timeout values
	// are only valid for service groups.
	Timeout int `json:"timeout"`

	// Id - <nil>
	Id int `json:"id"`

	// RoutingMethodId - <nil>
	RoutingMethodId int `json:"routingMethodId"`
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group

	// ServiceReferences - <nil>
	ServiceReferences []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference `json:"serviceReferences"`

	// Services - <nil>
	Services []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service `json:"services"`

	// ServiceCount - no documentation
	ServiceCount uint64 `json:"serviceCount"`

	// ServiceReferenceCount - no documentation
	ServiceReferenceCount uint64 `json:"serviceReferenceCount"`

	// VirtualServerCount - no documentation
	VirtualServerCount uint64 `json:"virtualServerCount"`

	// VirtualServers - <nil>
	VirtualServers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualServer `json:"virtualServers"`

	// RoutingType - <nil>
	RoutingType *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Type `json:"routingType"`

	// RoutingMethod - <nil>
	RoutingMethod *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Routing_Method `json:"routingMethod"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service_group *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group"
}
