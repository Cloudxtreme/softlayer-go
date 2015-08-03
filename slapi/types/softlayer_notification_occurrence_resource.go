package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource - This type contains general information relating to any
// hardware or services that may be impacted by a SoftLayer_Notification_Occurrence_Event.
type SoftLayer_Notification_Occurrence_Resource struct {

	// Active - <nil>
	Active int `json:"active"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId"`

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_notification_occurrence_resource *SoftLayer_Notification_Occurrence_Resource) String() string {
	return "SoftLayer_Notification_Occurrence_Resource"
}
