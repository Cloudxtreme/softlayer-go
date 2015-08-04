package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress
// - This type contains general information related to a
// [[SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress]] resource that is
// impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress struct {

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// PublicIp - <nil>
	PublicIp string `json:"publicIp,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}

// SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended is SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress with all maskable types.
type SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended struct {
	SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_application_delivery_controller_loadbalancer_virtualipaddress *SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress"
}
