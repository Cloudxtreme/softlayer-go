package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Status_Code - <nil>
type SoftLayer_Notification_Occurrence_Status_Code struct {

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// NotificationOccurrenceEventType - <nil>
	NotificationOccurrenceEventType *SoftLayer_Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType,omitempty"`
}

func (softlayer_notification_occurrence_status_code *SoftLayer_Notification_Occurrence_Status_Code) String() string {
	return "SoftLayer_Notification_Occurrence_Status_Code"
}
