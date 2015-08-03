package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
// - This type contains general information related to a
// [[SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress]] resource that is
// impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {

	// Hostname - <nil>
	Hostname string `json:"hostname"`

	// PublicIp - <nil>
	PublicIp string `json:"publicIp"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType"`
}

func (softlayer_notification_occurrence_resource_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}
