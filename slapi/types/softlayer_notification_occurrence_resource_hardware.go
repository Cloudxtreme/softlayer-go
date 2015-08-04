package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Hardware - This type contains general information related
// to a [[SoftLayer_Hardware]] resource that is impacted by a
// [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Hardware struct {

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp,omitempty"`

	// PublicIp - <nil>
	PublicIp string `json:"publicIp,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`
}

func (softlayer_notification_occurrence_resource_hardware *SoftLayer_Notification_Occurrence_Resource_Hardware) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Hardware"
}

// SoftLayer_Notification_Occurrence_Resource_Hardware_Extended is SoftLayer_Notification_Occurrence_Resource_Hardware with all maskable types.
type SoftLayer_Notification_Occurrence_Resource_Hardware_Extended struct {
	SoftLayer_Notification_Occurrence_Resource_Hardware

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`
}

func (softlayer_notification_occurrence_resource_hardware *SoftLayer_Notification_Occurrence_Resource_Hardware_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Hardware"
}
