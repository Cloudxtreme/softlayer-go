package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Notification_Emergency - A SoftLayer_Auxiliary_Notification_Emergency data
// object represents a notification event being broadcast to the SoftLayer customer base. It is used to
// provide information regarding outages or current known issues.
type SoftLayer_Auxiliary_Notification_Emergency struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Duration - no documentation
	Duration string `json:"duration"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// Location - no documentation
	Location string `json:"location"`

	// Message - no documentation
	Message string `json:"message"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// ServicesAffected - no documentation
	ServicesAffected string `json:"servicesAffected"`

	// Device - no documentation
	Device string `json:"device"`

	// Id - no documentation
	Id int `json:"id"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}

// SoftLayer_Auxiliary_Notification_Emergency_Extended is SoftLayer_Auxiliary_Notification_Emergency with all maskable types.
type SoftLayer_Auxiliary_Notification_Emergency_Extended struct {
	SoftLayer_Auxiliary_Notification_Emergency

	// Signature - The signature of the SoftLayer employee department associated with this notification.
	Signature *SoftLayer_Auxiliary_Notification_Emergency_Signature `json:"signature"`

	// Status - no documentation
	Status *SoftLayer_Auxiliary_Notification_Emergency_Status `json:"status"`
}

func (softlayer_auxiliary_notification_emergency *SoftLayer_Auxiliary_Notification_Emergency) String() string {
	return "SoftLayer_Auxiliary_Notification_Emergency"
}
