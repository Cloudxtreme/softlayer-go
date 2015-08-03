package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference - <nil>
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference struct {

	// ServiceGroupId - <nil>
	ServiceGroupId int `json:"serviceGroupId"`

	// ServiceId - <nil>
	ServiceId int `json:"serviceId"`

	// Weight - <nil>
	Weight int `json:"weight"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service_group_crossreference *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference"
}

// SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference_Extended is SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference

	// Service - <nil>
	Service *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service `json:"service"`

	// ServiceGroup - <nil>
	ServiceGroup *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group `json:"serviceGroup"`
}

func (softlayer_network_application_delivery_controller_loadbalancer_service_group_crossreference *SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group_CrossReference"
}
