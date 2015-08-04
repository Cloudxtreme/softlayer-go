package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Virtual - This type contains general information related
// to a [[SoftLayer_Virtual_Guest]] resource that is impacted by a
// [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Virtual struct {

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp,omitempty"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// PublicIp - <nil>
	PublicIp string `json:"publicIp,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`
}

func (softlayer_notification_occurrence_resource_virtual *SoftLayer_Notification_Occurrence_Resource_Virtual) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Virtual"
}

// SoftLayer_Notification_Occurrence_Resource_Virtual_Extended is SoftLayer_Notification_Occurrence_Resource_Virtual with all maskable types.
type SoftLayer_Notification_Occurrence_Resource_Virtual_Extended struct {
	SoftLayer_Notification_Occurrence_Resource_Virtual

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`
}

func (softlayer_notification_occurrence_resource_virtual *SoftLayer_Notification_Occurrence_Resource_Virtual_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Virtual"
}
