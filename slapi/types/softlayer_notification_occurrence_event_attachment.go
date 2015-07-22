package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Occurrence_Event_Attachment - SoftLayer events can have have files attached
// to them by a SoftLayer employee. Attaching a file to a event is a way to provide supplementary
// information such as a RFO (reason for outage) document or root cause analysis. The
// SoftLayer_Notification_Occurrence_Event_Attachment data type models a single file attached to a
// event.
type SoftLayer_Notification_Occurrence_Event_Attachment struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// FileName - no documentation
	FileName string `json:"fileName"`

	// FileSize - no documentation
	FileSize string `json:"fileSize"`

	// Id - no documentation
	Id int `json:"id"`

	// NotificationOccurrenceEvent - <nil>
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`

	// NotificationOccurrenceEventId - The unique event identifier that the file is attached to.
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId"`
}