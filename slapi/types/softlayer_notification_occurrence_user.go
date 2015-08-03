package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_User - This type contains general information relating to a user
// that may be impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_User struct {

	// AcknowledgedFlag - <nil>
	AcknowledgedFlag int `json:"acknowledgedFlag"`

	// Active - <nil>
	Active int `json:"active"`

	// Id - <nil>
	Id int `json:"id"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`
}

func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User) String() string {
	return "SoftLayer_Notification_Occurrence_User"
}

// SoftLayer_Notification_Occurrence_User_Extended is SoftLayer_Notification_Occurrence_User with all maskable types.
type SoftLayer_Notification_Occurrence_User_Extended struct {
	SoftLayer_Notification_Occurrence_User

	// ImpactedResourceCount - A count of a collection of resources impacted by the associated event.
	ImpactedResourceCount uint64 `json:"impactedResourceCount"`

	// ImpactedResources - A collection of resources impacted by the associated event.
	ImpactedResources []*SoftLayer_Notification_Occurrence_Resource `json:"impactedResources"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_notification_occurrence_user *SoftLayer_Notification_Occurrence_User_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_User"
}
