package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Account - <nil>
type SoftLayer_Notification_Occurrence_Account struct {

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// LastNotificationUpdate - <nil>
	LastNotificationUpdate *SoftLayer_Notification_Occurrence_Update `json:"lastNotificationUpdate,omitempty"`

	// NotificationOccurrenceEvent - <nil>
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
}

func (softlayer_notification_occurrence_account *SoftLayer_Notification_Occurrence_Account) String() string {
	return "SoftLayer_Notification_Occurrence_Account"
}
