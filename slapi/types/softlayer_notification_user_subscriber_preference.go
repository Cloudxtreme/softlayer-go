package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Notification_User_Subscriber_Preference - Preferences are settings that can be modified to
// change the behavior of the subscription. For example, modify the limit preference to only receive
// notifications 10 times instead of 1 during a billing cycle. Some preferences have certain
// restrictions on values that can be set.
type SoftLayer_Notification_User_Subscriber_Preference struct {

	// DefaultPreference - Details such name, keyname, minimum and maximum values for the preference.
	DefaultPreference *SoftLayer_Notification_Preference `json:"defaultPreference"`

	// Id - Unique identifier for the subscriber's preferences.
	Id int `json:"id"`

	// NotificationPreferenceId - Unique identifier of the default preference for which the subscriber
	// preference is based on. For example, if no preferences are supplied during the creation of a
	// subscriber. The default values are pulled using this property.
	NotificationPreferenceId int `json:"notificationPreferenceId"`

	// NotificationUserSubscriber - no documentation
	NotificationUserSubscriber *SoftLayer_Notification_User_Subscriber `json:"notificationUserSubscriber"`

	// NotificationUserSubscriberId - Unique identifier of the subscriber tied to the subscriber
	// preference.
	NotificationUserSubscriberId int `json:"notificationUserSubscriberId"`

	// Value - The user supplied value to "override" the "default" preference's value.
	Value string `json:"value"`
}

func (softlayer_notification_user_subscriber_preference *SoftLayer_Notification_User_Subscriber_Preference) String() string {
	return "SoftLayer_Notification_User_Subscriber_Preference"
}

// CreateObject - Use the method to create a new notification preference for a subscriber
func (softlayer_notification_user_subscriber_preference *SoftLayer_Notification_User_Subscriber_Preference) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Notification_User_Subscriber_Preference) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - <nil>
func (softlayer_notification_user_subscriber_preference *SoftLayer_Notification_User_Subscriber_Preference) EditObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Notification_User_Subscriber_Preference) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_user_subscriber_preference *SoftLayer_Notification_User_Subscriber_Preference) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Notification_User_Subscriber_Preference, error) {
	var returnValue *SoftLayer_Notification_User_Subscriber_Preference
	return returnValue, nil
}
