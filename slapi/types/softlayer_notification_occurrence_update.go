package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Occurrence_Update - <nil>
type SoftLayer_Notification_Occurrence_Update struct {

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`

	// Contents - <nil>
	Contents string `json:"contents"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`
}

// SoftLayer_Notification_Occurrence_Update_Extended is SoftLayer_Notification_Occurrence_Update with all maskable types.
type SoftLayer_Notification_Occurrence_Update_Extended struct {
	SoftLayer_Notification_Occurrence_Update

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee"`

	// NotificationOccurrenceEvent - <nil>
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`
}

func (softlayer_notification_occurrence_update *SoftLayer_Notification_Occurrence_Update) String() string {
	return "SoftLayer_Notification_Occurrence_Update"
}
