package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Status_Code - <nil>
type SoftLayer_Notification_Occurrence_Status_Code struct {

	// Description - <nil>
	Description string `json:"description"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// SoftLayer_Notification_Occurrence_Status_Code_Extended is SoftLayer_Notification_Occurrence_Status_Code with all maskable types.
type SoftLayer_Notification_Occurrence_Status_Code_Extended struct {
	SoftLayer_Notification_Occurrence_Status_Code

	// NotificationOccurrenceEventType - <nil>
	NotificationOccurrenceEventType *SoftLayer_Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType"`
}

func (softlayer_notification_occurrence_status_code *SoftLayer_Notification_Occurrence_Status_Code) String() string {
	return "SoftLayer_Notification_Occurrence_Status_Code"
}
