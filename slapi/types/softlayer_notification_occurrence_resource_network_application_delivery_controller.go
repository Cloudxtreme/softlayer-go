package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller - This type
// contains general information related to a [[SoftLayer_Network_Application_Delivery_Controller]]
// resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {

	// PublicIp - <nil>
	PublicIp string `json:"publicIp,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_application_delivery_controller *SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller"
}
