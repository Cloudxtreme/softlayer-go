package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

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

// GetGraphImage - Get the graph image for a load balancer service group based on the supplied graph
// type and metric. The only available graph type currently is: 'connections', and the available
// metrics are: 'day', 'week' and 'month'. This method returns the raw binary image data.
func (softlayer_network_application_delivery_controller_loadbalancer_service_group *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetGraphImage(ctx *slapi.RequestContext, graphType string, metric string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_service_group *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group
	return returnValue, nil
}

// KickAllConnections - <nil>
func (softlayer_network_application_delivery_controller_loadbalancer_service_group *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group) KickAllConnections(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
