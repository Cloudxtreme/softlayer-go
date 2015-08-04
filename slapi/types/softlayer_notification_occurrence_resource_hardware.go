package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Hardware - This type contains general information related
// to a [[SoftLayer_Hardware]] resource that is impacted by a
// [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Hardware struct {

	// PublicIp - <nil>
	PublicIp string `json:"publicIp,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`
}

func (softlayer_notification_occurrence_resource_hardware *SoftLayer_Notification_Occurrence_Resource_Hardware) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Hardware"
}

// SoftLayer_Notification_Occurrence_Resource_Hardware_Extended is SoftLayer_Notification_Occurrence_Resource_Hardware with all maskable types.
type SoftLayer_Notification_Occurrence_Resource_Hardware_Extended struct {
	SoftLayer_Notification_Occurrence_Resource_Hardware

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
}

func (softlayer_notification_occurrence_resource_hardware *SoftLayer_Notification_Occurrence_Resource_Hardware_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Hardware"
}
