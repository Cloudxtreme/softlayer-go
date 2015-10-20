package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Preference - Retrieve details for preferences. Preferences are used to allow
// the subscriber to modify their subscription in various ways. Details such as friendly name, keyname
// maximum and minimum values can be retrieved. These provide details to help configure subscriber
// preferences correctly.
type SoftLayer_Notification_Preference struct {

	// Id - Unique identifier for the notification preference.
	Id int `json:"id,omitempty"`

	// KeyName - Name that can be used by external systems to refer to preference.
	KeyName string `json:"keyName,omitempty"`

	// MaximumValue - no documentation
	MaximumValue string `json:"maximumValue,omitempty"`

	// MinimumValue - no documentation
	MinimumValue string `json:"minimumValue,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Units - The unit of measure used for the preference's value, minimum and maximum as well.
	Units string `json:"units,omitempty"`

	// Value - Default value used when setting up preferences for a new subscriber.
	Value string `json:"value,omitempty"`

	// Description - no documentation
	Description string `json:"description,omitempty"`
}

func (softlayer_notification_preference *SoftLayer_Notification_Preference) String() string {
	return "SoftLayer_Notification_Preference"
}
