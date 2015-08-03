package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_User_Subscriber - A notification subscriber will have details pertaining to
// the subscriber's notification subscription. You can receive details such as preferences, details of
// the preferences, delivery methods and the delivery methods for the subscriber. There are preferences
// and delivery methods that cannot be modified. Also, there are some subscriptions that are required.
type SoftLayer_Notification_User_Subscriber struct {

	// Active - no documentation
	Active int `json:"active"`

	// DeliveryMethodCount - A count of the delivery methods used to send the subscribed notification.
	DeliveryMethodCount uint64 `json:"deliveryMethodCount"`

	// DeliveryMethods - The delivery methods used to send the subscribed notification.
	DeliveryMethods []*SoftLayer_Notification_Delivery_Method `json:"deliveryMethods"`

	// Id - Unique identifier of the subscriber that will receive the alerts.
	Id int `json:"id"`

	// Notification - no documentation
	Notification *SoftLayer_Notification `json:"notification"`

	// NotificationId - Unique identifier of the notification subscribed to.
	NotificationId int `json:"notificationId"`

	// PreferenceCount - A count of associated subscriber preferences used for the notification
	// subscription. For example, preferences include number of deliveries (limit) and threshold.
	PreferenceCount uint64 `json:"preferenceCount"`

	// Preferences - Associated subscriber preferences used for the notification subscription. For example,
	// preferences include number of deliveries (limit) and threshold.
	Preferences []*SoftLayer_Notification_User_Subscriber_Preference `json:"preferences"`

	// PreferencesDetailCount - A count of preference details such as description, minimum and maximum
	// limits, default value and unit of measure.
	PreferencesDetailCount uint64 `json:"preferencesDetailCount"`

	// PreferencesDetails - Preference details such as description, minimum and maximum limits, default
	// value and unit of measure.
	PreferencesDetails []*SoftLayer_Notification_Preference `json:"preferencesDetails"`

	// ResourceRecord - no documentation
	ResourceRecord *SoftLayer_Notification_User_Subscriber_Resource `json:"resourceRecord"`

	// UserRecord - no documentation
	UserRecord *SoftLayer_User_Customer `json:"userRecord"`

	// UserRecordId - Unique identifier of the user the subscription is for.
	UserRecordId int `json:"userRecordId"`
}

func (softlayer_notification_user_subscriber *SoftLayer_Notification_User_Subscriber) String() string {
	return "SoftLayer_Notification_User_Subscriber"
}
