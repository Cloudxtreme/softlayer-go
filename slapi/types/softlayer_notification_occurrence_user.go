package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Notification_Occurrence_User - This type contains general information relating to a user
// that may be impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_User struct {

	// AcknowledgedFlag - <nil>
	AcknowledgedFlag int `json:"acknowledgedFlag"`

	// Active - <nil>
	Active int `json:"active"`

	// Id - <nil>
	Id int `json:"id"`

	// ImpactedResourceCount - A count of a collection of resources impacted by the associated event.
	ImpactedResourceCount uint64 `json:"impactedResourceCount"`

	// ImpactedResources - A collection of resources impacted by the associated event.
	ImpactedResources []*SoftLayer_Notification_Occurrence_Resource `json:"impactedResources"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`
}

func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) String() string {
	return "SoftLayer_Notification_Occurrence_User"
}

// Acknowledge - <nil>
func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) Acknowledge(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - <nil>
func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Notification_Occurrence_User, error) {
	var returnValue []*SoftLayer_Notification_Occurrence_User
	return returnValue, nil
}

// GetImpactedDeviceCount - <nil>
func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) GetImpactedDeviceCount(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Notification_Occurrence_User, error) {
	var returnValue *SoftLayer_Notification_Occurrence_User
	return returnValue, nil
}
