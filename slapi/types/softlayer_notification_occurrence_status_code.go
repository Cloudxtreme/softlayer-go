package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Status_Code - <nil>
type SoftLayer_Notification_Occurrence_Status_Code struct {

	// Description - <nil>
	Description string `json:"description,omitempty"`

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_notification_occurrence_status_code *SoftLayer_Notification_Occurrence_Status_Code) String() string {
	return "SoftLayer_Notification_Occurrence_Status_Code"
}

// SoftLayer_Notification_Occurrence_Status_Code_Extended is SoftLayer_Notification_Occurrence_Status_Code with all maskable types.
type SoftLayer_Notification_Occurrence_Status_Code_Extended struct {
	SoftLayer_Notification_Occurrence_Status_Code

	// NotificationOccurrenceEventType - <nil>
	NotificationOccurrenceEventType *SoftLayer_Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType,omitempty"`
}

func (softlayer_notification_occurrence_status_code *SoftLayer_Notification_Occurrence_Status_Code_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Status_Code"
}
