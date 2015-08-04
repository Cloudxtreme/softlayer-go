package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Occurrence_Update - <nil>
type SoftLayer_Notification_Occurrence_Update struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// Contents - <nil>
	Contents string `json:"contents,omitempty"`
}

func (softlayer_notification_occurrence_update *SoftLayer_Notification_Occurrence_Update) String() string {
	return "SoftLayer_Notification_Occurrence_Update"
}

// SoftLayer_Notification_Occurrence_Update_Extended is SoftLayer_Notification_Occurrence_Update with all maskable types.
type SoftLayer_Notification_Occurrence_Update_Extended struct {
	SoftLayer_Notification_Occurrence_Update

	// Employee - <nil>
	Employee *SoftLayer_User_Employee `json:"employee,omitempty"`

	// NotificationOccurrenceEvent - <nil>
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
}

func (softlayer_notification_occurrence_update *SoftLayer_Notification_Occurrence_Update_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Update"
}
