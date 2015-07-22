package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Notification_Occurrence_Event - <nil>
type SoftLayer_Notification_Occurrence_Event struct {

	// AttachmentCount - A count of a collection of attachments for this event which provide supplementary
	// information to impacted users some examples are RFO (Reason For Outage) and root cause analysis
	// documents.
	AttachmentCount uint64 `json:"attachmentCount"`

	// Attachments - A collection of attachments for this event which provide supplementary information to
	// impacted users some examples are RFO (Reason For Outage) and root cause analysis documents.
	Attachments []*SoftLayer_Notification_Occurrence_Event_Attachment `json:"attachments"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// FirstUpdate - no documentation
	FirstUpdate *SoftLayer_Notification_Occurrence_Update `json:"firstUpdate"`

	// Id - <nil>
	Id int `json:"id"`

	// ImpactedAccountCount - A count of a collection of accounts impacted by this event. Each impacted
	// account record relates directly to a [[SoftLayer_Account]].
	ImpactedAccountCount uint64 `json:"impactedAccountCount"`

	// ImpactedAccounts - A collection of accounts impacted by this event. Each impacted account record
	// relates directly to a [[SoftLayer_Account]].
	ImpactedAccounts []*SoftLayer_Notification_Occurrence_Account `json:"impactedAccounts"`

	// ImpactedResourceCount - A count of a collection of resources impacted by this event. Each record
	// will relate to some physical resource that the user has access to such as [[SoftLayer_Hardware]] or
	// [[SoftLayer_Virtual_Guest]].
	ImpactedResourceCount uint64 `json:"impactedResourceCount"`

	// ImpactedResources - A collection of resources impacted by this event. Each record will relate to
	// some physical resource that the user has access to such as [[SoftLayer_Hardware]] or
	// [[SoftLayer_Virtual_Guest]].
	ImpactedResources []*SoftLayer_Notification_Occurrence_Resource `json:"impactedResources"`

	// ImpactedUserCount - A count of a collection of users impacted by this event. Each impacted user
	// record relates directly to a [[SoftLayer_User_Customer]].
	ImpactedUserCount uint64 `json:"impactedUserCount"`

	// ImpactedUsers - A collection of users impacted by this event. Each impacted user record relates
	// directly to a [[SoftLayer_User_Customer]].
	ImpactedUsers []*SoftLayer_Notification_Occurrence_User `json:"impactedUsers"`

	// LastImpactedUserCount - <nil>
	LastImpactedUserCount int `json:"lastImpactedUserCount"`

	// LastUpdate - no documentation
	LastUpdate *SoftLayer_Notification_Occurrence_Update `json:"lastUpdate"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// NotificationOccurrenceEventType - The type of event such as planned or unplanned maintenance.
	NotificationOccurrenceEventType *SoftLayer_Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType"`

	// RecoveryTime - <nil>
	RecoveryTime int `json:"recoveryTime"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`

	// StatusCode - <nil>
	StatusCode *SoftLayer_Notification_Occurrence_Status_Code `json:"statusCode"`

	// Subject - <nil>
	Subject string `json:"subject"`

	// Summary - <nil>
	Summary string `json:"summary"`

	// SystemTicketId - <nil>
	SystemTicketId int `json:"systemTicketId"`

	// UpdateCount - no documentation
	UpdateCount uint64 `json:"updateCount"`

	// Updates - no documentation
	Updates []*SoftLayer_Notification_Occurrence_Update `json:"updates"`
}

// AcknowledgeNotification - no documentation
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) AcknowledgeNotification(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllObjects - <nil>
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Notification_Occurrence_Event, error) {
	var returnValue []*SoftLayer_Notification_Occurrence_Event
	return returnValue, nil
}

// GetAttachedFile - Retrieve the contents of the file attached to a SoftLayer event by it's given
// identifier.
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) GetAttachedFile(commonOptions *slapi.CommonOptions, attachmentId int) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetImpactedAccountCount - This method will return the number of impacted owned accounts associated
// with this event for the current user.
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) GetImpactedAccountCount(commonOptions *slapi.CommonOptions) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetImpactedDeviceCount - This method will return the number of impacted devices associated with this
// event for the current user.
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) GetImpactedDeviceCount(commonOptions *slapi.CommonOptions) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetImpactedDevices - This method will return a collection of
// SoftLayer_Notification_Occurrence_Resource objects which is a listing of the current users' impacted
// devices that are associated with this event.
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) GetImpactedDevices(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Notification_Occurrence_Resource, error) {
	var returnValue []*SoftLayer_Notification_Occurrence_Resource
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Notification_Occurrence_Event, error) {
	var returnValue *SoftLayer_Notification_Occurrence_Event
	return returnValue, nil
}
