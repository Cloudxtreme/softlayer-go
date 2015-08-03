package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Occurrence_Update - <nil>
type SoftLayer_Notification_Occurrence_Update struct {

	// Contents - <nil>
	Contents string `json:"contents"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// NotificationOccurrenceEvent - <nil>
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_notification_occurrence_update *SoftLayer_Notification_Occurrence_Update) String() string {
	return "SoftLayer_Notification_Occurrence_Update"
}
