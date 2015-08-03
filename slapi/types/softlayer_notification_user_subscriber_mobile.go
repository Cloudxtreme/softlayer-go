package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Notification_User_Subscriber_Mobile - A notification subscriber will have details
// pertaining to the subscriber's notification subscription. You can receive details such as
// preferences, details of the preferences, delivery methods and the delivery methods for the
// subscriber. There are preferences and delivery methods that cannot be modified. Also, there are some
// subscriptions that are required.
type SoftLayer_Notification_User_Subscriber_Mobile struct {
}

func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) String() string {
	return "SoftLayer_Notification_User_Subscriber_Mobile"
}

// ClearSnoozeTimer - <nil>
func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) ClearSnoozeTimer(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Notification_User_Subscriber_Mobile, error) {
	var returnValue *SoftLayer_Notification_User_Subscriber_Mobile
	return returnValue, nil
}

// SetSnoozeTimer - <nil>
func (softlayer_notification_user_subscriber_mobile *SoftLayer_Notification_User_Subscriber_Mobile) SetSnoozeTimer(ctx *slapi.RequestContext, start int, end int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
