package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Account - <nil>
type SoftLayer_Notification_Occurrence_Account struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Active - <nil>
	Active int `json:"active"`

	// LastNotificationUpdate - <nil>
	LastNotificationUpdate *SoftLayer_Notification_Occurrence_Update `json:"lastNotificationUpdate"`

	// NotificationOccurrenceEvent - <nil>
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`
}

func (softlayer_notification_occurrence_account *SoftLayer_Notification_Occurrence_Account) String() string {
	return "SoftLayer_Notification_Occurrence_Account"
}
