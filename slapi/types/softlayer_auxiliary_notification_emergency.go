package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Notification_Emergency - A SoftLayer_Auxiliary_Notification_Emergency data
// object represents a notification event being broadcast to the SoftLayer customer base. It is used to
// provide information regarding outages or current known issues.
type SoftLayer_Auxiliary_Notification_Emergency struct {

	// Duration - no documentation
	Duration string `json:"duration"`

	// Id - no documentation
	Id int `json:"id"`

	// Message - no documentation
	Message string `json:"message"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// Device - no documentation
	Device string `json:"device"`

	// Location - no documentation
	Location string `json:"location"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// ServicesAffected - no documentation
	ServicesAffected string `json:"servicesAffected"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_auxiliary_notification_emergency *SoftLayer_Auxiliary_Notification_Emergency) String() string {
	return "SoftLayer_Auxiliary_Notification_Emergency"
}

// SoftLayer_Auxiliary_Notification_Emergency_Extended is SoftLayer_Auxiliary_Notification_Emergency with all maskable types.
type SoftLayer_Auxiliary_Notification_Emergency_Extended struct {
	SoftLayer_Auxiliary_Notification_Emergency

	// Status - no documentation
	Status *SoftLayer_Auxiliary_Notification_Emergency_Status `json:"status"`

	// Signature - The signature of the SoftLayer employee department associated with this notification.
	Signature *SoftLayer_Auxiliary_Notification_Emergency_Signature `json:"signature"`
}

func (softlayer_auxiliary_notification_emergency *SoftLayer_Auxiliary_Notification_Emergency_Extended) String() string {
	return "SoftLayer_Auxiliary_Notification_Emergency"
}
