package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

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

// CreateObject - Use the method to create a new subscription for a notification. This method is the
// entry method to the notification system. Certain properties are required to create a subscription
// while others are optional. The required property is the resourceRecord property which is type
// SoftLayer_Notification_User_Subscriber_Resource. For the resourceRecord property, the only property
// that needs to be populated is the resourceTableId. The resourceTableId is the unique identifier of a
// SoftLayer service to create the subscription for. For example, the unique identifier of the Storage
// Evault service to create the subscription on. Optional properties that can be set is the preferences
// property. The preference property is an array SoftLayer_Notification_User_Subscriber_Preference. By
// default, the system will populate the preferences with the default values if no preferences are
// passed in. The preferences passed in must be the preferences related to the notification subscribing
// to. The notification preferences and preference details (such as minimum and maximum values) can be
// retrieved using the SoftLayer_Notification service. The properties that need to be populated for
// preferences are the notificationPreferenceId and value. For example to create a subscriber for a
// Storage EVault service to be notified 15 times during a billing cycle and to be notified when the
// vault usage reaches 85% of its allowed capacity use the following structure: *userRecordId = 1111
// *notificationId = 3 *resourceRecord **resourceTableId = 1234 *preferences[1]
// **notificationPreferenceId = 2 **value = 85 *preference[2] **notificationPreferenceId = 3 **value =
// 15
func (softlayer_notification_user_subscriber *SoftLayer_Notification_User_Subscriber) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Notification_User_Subscriber) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - The subscriber's subscription status can be "turned off" or "turned on" if the
// subscription is not required. Subscriber preferences may also be edited. To edit the preferences,
// you must pass in the id off the preferences to edit. Here is an example of structure to pass in. In
// this example, the structure will set the subscriber status to active and the threshold preference to
// 90 and the limit preference to 20 *id = 1111 *active = 1 *preferences[1] **id = 11 **value = 90
// *preference[2] **id = 12 **value = 20
func (softlayer_notification_user_subscriber *SoftLayer_Notification_User_Subscriber) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Notification_User_Subscriber) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_user_subscriber *SoftLayer_Notification_User_Subscriber) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Notification_User_Subscriber, error) {
	var returnValue *SoftLayer_Notification_User_Subscriber
	return returnValue, nil
}
