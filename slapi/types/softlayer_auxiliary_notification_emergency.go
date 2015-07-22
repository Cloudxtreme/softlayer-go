package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Notification_Emergency - A SoftLayer_Auxiliary_Notification_Emergency data
// object represents a notification event being broadcast to the SoftLayer customer base. It is used to
// provide information regarding outages or current known issues.
type SoftLayer_Auxiliary_Notification_Emergency struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Device - no documentation
	Device string `json:"device"`

	// Duration - no documentation
	Duration string `json:"duration"`

	// Id - no documentation
	Id int `json:"id"`

	// Location - no documentation
	Location string `json:"location"`

	// Message - no documentation
	Message string `json:"message"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// ServicesAffected - no documentation
	ServicesAffected string `json:"servicesAffected"`

	// Signature - The signature of the SoftLayer employee department associated with this notification.
	Signature *SoftLayer_Auxiliary_Notification_Emergency_Signature `json:"signature"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// Status - no documentation
	Status *SoftLayer_Auxiliary_Notification_Emergency_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`
}

// GetAllObjects - Retrieve an array of SoftLayer_Auxiliary_Notification_Emergency data types, which
// contain all notification events regardless of status.
func (softlayer_auxiliary_notification_emergency *SoftLayer_Auxiliary_Notification_Emergency) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Notification_Emergency, error) {
	var returnValue []*SoftLayer_Auxiliary_Notification_Emergency
	return returnValue, nil
}

// GetCurrentNotifications - Retrieve an array of SoftLayer_Auxiliary_Notification_Emergency data
// types, which contain all current notification events.
func (softlayer_auxiliary_notification_emergency *SoftLayer_Auxiliary_Notification_Emergency) GetCurrentNotifications(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Notification_Emergency, error) {
	var returnValue []*SoftLayer_Auxiliary_Notification_Emergency
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Notification_Emergency object, it can be
// used to check for current notifications being broadcast by SoftLayer.
func (softlayer_auxiliary_notification_emergency *SoftLayer_Auxiliary_Notification_Emergency) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Notification_Emergency, error) {
	var returnValue *SoftLayer_Auxiliary_Notification_Emergency
	return returnValue, nil
}
