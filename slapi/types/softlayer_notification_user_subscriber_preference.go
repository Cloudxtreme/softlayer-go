package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_User_Subscriber_Preference - Preferences are settings that can be modified to
// change the behavior of the subscription. For example, modify the limit preference to only receive
// notifications 10 times instead of 1 during a billing cycle. Some preferences have certain
// restrictions on values that can be set.
type SoftLayer_Notification_User_Subscriber_Preference struct {

	// Id - Unique identifier for the subscriber's preferences.
	Id int `json:"id"`

	// NotificationPreferenceId - Unique identifier of the default preference for which the subscriber
	// preference is based on. For example, if no preferences are supplied during the creation of a
	// subscriber. The default values are pulled using this property.
	NotificationPreferenceId int `json:"notificationPreferenceId"`

	// NotificationUserSubscriberId - Unique identifier of the subscriber tied to the subscriber
	// preference.
	NotificationUserSubscriberId int `json:"notificationUserSubscriberId"`

	// Value - The user supplied value to "override" the "default" preference's value.
	Value string `json:"value"`
}

func (softlayer_notification_user_subscriber_preference *SoftLayer_Notification_User_Subscriber_Preference) String() string {
	return "SoftLayer_Notification_User_Subscriber_Preference"
}

// SoftLayer_Notification_User_Subscriber_Preference_Extended is SoftLayer_Notification_User_Subscriber_Preference with all maskable types.
type SoftLayer_Notification_User_Subscriber_Preference_Extended struct {
	SoftLayer_Notification_User_Subscriber_Preference

	// DefaultPreference - Details such name, keyname, minimum and maximum values for the preference.
	DefaultPreference *SoftLayer_Notification_Preference `json:"defaultPreference"`

	// NotificationUserSubscriber - no documentation
	NotificationUserSubscriber *SoftLayer_Notification_User_Subscriber `json:"notificationUserSubscriber"`
}

func (softlayer_notification_user_subscriber_preference *SoftLayer_Notification_User_Subscriber_Preference_Extended) String() string {
	return "SoftLayer_Notification_User_Subscriber_Preference"
}
