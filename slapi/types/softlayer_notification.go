package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification - Details provided for the notification are basic. Details such as the
// related preferences, name and keyname for the notification can be retrieved. The keyname property
// for the notification can be used to refer to a notification when integrating into the SoftLayer
// Notification system. The name property can used more for display purposes.
type SoftLayer_Notification struct {

	// KeyName - Name that can be used by external systems to refer to a notification.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`

	// Id - no documentation
	Id int `json:"id"`
}

// SoftLayer_Notification_Extended is SoftLayer_Notification with all maskable types.
type SoftLayer_Notification_Extended struct {
	SoftLayer_Notification

	// PreferenceCount - A count of the preferences related to the notification. These are preferences are
	// configurable and optional for subscribers to use.
	PreferenceCount uint64 `json:"preferenceCount"`

	// RequiredPreferenceCount - A count of the required preferences related to the notification. While
	// configurable, the subscriber does not have the option whether to use the preference.
	RequiredPreferenceCount uint64 `json:"requiredPreferenceCount"`

	// Preferences - The preferences related to the notification. These are preferences are configurable
	// and optional for subscribers to use.
	Preferences []*SoftLayer_Notification_Preference `json:"preferences"`

	// RequiredPreferences - The required preferences related to the notification. While configurable, the
	// subscriber does not have the option whether to use the preference.
	RequiredPreferences []*SoftLayer_Notification_Preference `json:"requiredPreferences"`
}

func (softlayer_notification *SoftLayer_Notification) String() string {
	return "SoftLayer_Notification"
}
