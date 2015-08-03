package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Occurrence_Event - <nil>
type SoftLayer_Notification_Occurrence_Event struct {

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate"`

	// RecoveryTime - <nil>
	RecoveryTime int `json:"recoveryTime"`

	// Subject - <nil>
	Subject string `json:"subject"`

	// SystemTicketId - <nil>
	SystemTicketId int `json:"systemTicketId"`

	// LastImpactedUserCount - <nil>
	LastImpactedUserCount int `json:"lastImpactedUserCount"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Summary - <nil>
	Summary string `json:"summary"`

	// Id - <nil>
	Id int `json:"id"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) String() string {
	return "SoftLayer_Notification_Occurrence_Event"
}

// SoftLayer_Notification_Occurrence_Event_Extended is SoftLayer_Notification_Occurrence_Event with all maskable types.
type SoftLayer_Notification_Occurrence_Event_Extended struct {
	SoftLayer_Notification_Occurrence_Event

	// ImpactedUserCount - A count of a collection of users impacted by this event. Each impacted user
	// record relates directly to a [[SoftLayer_User_Customer]].
	ImpactedUserCount uint64 `json:"impactedUserCount"`

	// Attachments - A collection of attachments for this event which provide supplementary information to
	// impacted users some examples are RFO (Reason For Outage) and root cause analysis documents.
	Attachments []*SoftLayer_Notification_Occurrence_Event_Attachment `json:"attachments"`

	// NotificationOccurrenceEventType - The type of event such as planned or unplanned maintenance.
	NotificationOccurrenceEventType *SoftLayer_Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// ImpactedAccounts - A collection of accounts impacted by this event. Each impacted account record
	// relates directly to a [[SoftLayer_Account]].
	ImpactedAccounts []*SoftLayer_Notification_Occurrence_Account `json:"impactedAccounts"`

	// ImpactedResources - A collection of resources impacted by this event. Each record will relate to
	// some physical resource that the user has access to such as [[SoftLayer_Hardware]] or
	// [[SoftLayer_Virtual_Guest]].
	ImpactedResources []*SoftLayer_Notification_Occurrence_Resource `json:"impactedResources"`

	// Updates - no documentation
	Updates []*SoftLayer_Notification_Occurrence_Update `json:"updates"`

	// ImpactedUsers - A collection of users impacted by this event. Each impacted user record relates
	// directly to a [[SoftLayer_User_Customer]].
	ImpactedUsers []*SoftLayer_Notification_Occurrence_User `json:"impactedUsers"`

	// ImpactedAccountCount - A count of a collection of accounts impacted by this event. Each impacted
	// account record relates directly to a [[SoftLayer_Account]].
	ImpactedAccountCount uint64 `json:"impactedAccountCount"`

	// AttachmentCount - A count of a collection of attachments for this event which provide supplementary
	// information to impacted users some examples are RFO (Reason For Outage) and root cause analysis
	// documents.
	AttachmentCount uint64 `json:"attachmentCount"`

	// ImpactedResourceCount - A count of a collection of resources impacted by this event. Each record
	// will relate to some physical resource that the user has access to such as [[SoftLayer_Hardware]] or
	// [[SoftLayer_Virtual_Guest]].
	ImpactedResourceCount uint64 `json:"impactedResourceCount"`

	// UpdateCount - no documentation
	UpdateCount uint64 `json:"updateCount"`

	// FirstUpdate - no documentation
	FirstUpdate *SoftLayer_Notification_Occurrence_Update `json:"firstUpdate"`

	// LastUpdate - no documentation
	LastUpdate *SoftLayer_Notification_Occurrence_Update `json:"lastUpdate"`

	// StatusCode - <nil>
	StatusCode *SoftLayer_Notification_Occurrence_Status_Code `json:"statusCode"`
}

func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Event"
}
