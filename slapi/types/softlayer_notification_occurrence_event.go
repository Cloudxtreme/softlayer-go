package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Occurrence_Event - <nil>
type SoftLayer_Notification_Occurrence_Event struct {

	// EndDate - <nil>
	EndDate *time.Time `json:"endDate,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Summary - <nil>
	Summary string `json:"summary,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// RecoveryTime - <nil>
	RecoveryTime int `json:"recoveryTime,omitempty"`

	// LastImpactedUserCount - <nil>
	LastImpactedUserCount int `json:"lastImpactedUserCount,omitempty"`

	// SystemTicketId - <nil>
	SystemTicketId int `json:"systemTicketId,omitempty"`

	// Subject - <nil>
	Subject string `json:"subject,omitempty"`

	// StartDate - <nil>
	StartDate *time.Time `json:"startDate,omitempty"`

	// UpdateCount - no documentation
	UpdateCount uint64 `json:"updateCount,omitempty"`

	// ImpactedResources - A collection of resources impacted by this event. Each record will relate to
	// some physical resource that the user has access to such as [[SoftLayer_Hardware]] or
	// [[SoftLayer_Virtual_Guest]].
	ImpactedResources []*SoftLayer_Notification_Occurrence_Resource `json:"impactedResources,omitempty"`

	// ImpactedUsers - A collection of users impacted by this event. Each impacted user record relates
	// directly to a [[SoftLayer_User_Customer]].
	ImpactedUsers []*SoftLayer_Notification_Occurrence_User `json:"impactedUsers,omitempty"`

	// LastUpdate - no documentation
	LastUpdate *SoftLayer_Notification_Occurrence_Update `json:"lastUpdate,omitempty"`

	// Updates - no documentation
	Updates []*SoftLayer_Notification_Occurrence_Update `json:"updates,omitempty"`

	// StatusCode - <nil>
	StatusCode *SoftLayer_Notification_Occurrence_Status_Code `json:"statusCode,omitempty"`

	// ImpactedUserCount - A count of a collection of users impacted by this event. Each impacted user
	// record relates directly to a [[SoftLayer_User_Customer]].
	ImpactedUserCount uint64 `json:"impactedUserCount,omitempty"`

	// Attachments - A collection of attachments for this event which provide supplementary information to
	// impacted users some examples are RFO (Reason For Outage) and root cause analysis documents.
	Attachments []*SoftLayer_Notification_Occurrence_Event_Attachment `json:"attachments,omitempty"`

	// FirstUpdate - no documentation
	FirstUpdate *SoftLayer_Notification_Occurrence_Update `json:"firstUpdate,omitempty"`

	// ImpactedAccounts - A collection of accounts impacted by this event. Each impacted account record
	// relates directly to a [[SoftLayer_Account]].
	ImpactedAccounts []*SoftLayer_Notification_Occurrence_Account `json:"impactedAccounts,omitempty"`

	// NotificationOccurrenceEventType - The type of event such as planned or unplanned maintenance.
	NotificationOccurrenceEventType *SoftLayer_Notification_Occurrence_Event_Type `json:"notificationOccurrenceEventType,omitempty"`

	// ImpactedResourceCount - A count of a collection of resources impacted by this event. Each record
	// will relate to some physical resource that the user has access to such as [[SoftLayer_Hardware]] or
	// [[SoftLayer_Virtual_Guest]].
	ImpactedResourceCount uint64 `json:"impactedResourceCount,omitempty"`

	// AttachmentCount - A count of a collection of attachments for this event which provide supplementary
	// information to impacted users some examples are RFO (Reason For Outage) and root cause analysis
	// documents.
	AttachmentCount uint64 `json:"attachmentCount,omitempty"`

	// ImpactedAccountCount - A count of a collection of accounts impacted by this event. Each impacted
	// account record relates directly to a [[SoftLayer_Account]].
	ImpactedAccountCount uint64 `json:"impactedAccountCount,omitempty"`
}

func (softlayer_notification_occurrence_event *SoftLayer_Notification_Occurrence_Event) String() string {
	return "SoftLayer_Notification_Occurrence_Event"
}
