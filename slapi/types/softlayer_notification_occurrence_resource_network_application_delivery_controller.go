package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller - This type
// contains general information related to a [[SoftLayer_Network_Application_Delivery_Controller]]
// resource that is impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller struct {

	// Hostname - <nil>
	Hostname string `json:"hostname"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp"`

	// PublicIp - <nil>
	PublicIp string `json:"publicIp"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType"`
}

func (softlayer_notification_occurrence_resource_network_application_delivery_controller *SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Application_Delivery_Controller"
}
