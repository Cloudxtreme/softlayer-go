package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Notification_User_Subscriber_Mobile - A notification subscriber will have details
// pertaining to the subscriber's notification subscription. You can receive details such as
// preferences, details of the preferences, delivery methods and the delivery methods for the
// subscriber. There are preferences and delivery methods that cannot be modified. Also, there are some
// subscriptions that are required.
type SoftLayer_Notification_User_Subscriber_Mobile struct {
}

// ClearSnoozeTimer - <nil>
func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) ClearSnoozeTimer(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Notification_User_Subscriber_Mobile, error) {
	var returnValue *SoftLayer_Notification_User_Subscriber_Mobile
	return returnValue, nil
}

// SetSnoozeTimer - <nil>
func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) SetSnoozeTimer(commonOptions *slapi.CommonOptions, start int, end int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
