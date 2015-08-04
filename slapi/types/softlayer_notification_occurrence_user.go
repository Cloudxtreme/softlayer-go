package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_User - This type contains general information relating to a user
// that may be impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_User struct {

	// AcknowledgedFlag - <nil>
	AcknowledgedFlag int `json:"acknowledgedFlag,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`
}

func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) String() string {
	return "SoftLayer_Notification_Occurrence_User"
}

// SoftLayer_Notification_Occurrence_User_Extended is SoftLayer_Notification_Occurrence_User with all maskable types.
type SoftLayer_Notification_Occurrence_User_Extended struct {
	SoftLayer_Notification_Occurrence_User

	// ImpactedResources - A collection of resources impacted by the associated event.
	ImpactedResources []*SoftLayer_Notification_Occurrence_Resource `json:"impactedResources,omitempty"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// ImpactedResourceCount - A count of a collection of resources impacted by the associated event.
	ImpactedResourceCount uint64 `json:"impactedResourceCount,omitempty"`
}

func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_User"
}
