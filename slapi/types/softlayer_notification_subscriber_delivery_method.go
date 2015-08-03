package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Subscriber_Delivery_Method - Provides details for the subscriber's delivery
// methods.
type SoftLayer_Notification_Subscriber_Delivery_Method struct {

	// Active - Indicates the subscriber's delivery method availability for notifications.
	Active int `json:"active"`

	// CreateDate - Date the subscriber's delivery method was created.
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - Date the subscriber's delivery method was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// NotificationDeliveryMethod - <nil>
	NotificationDeliveryMethod *SoftLayer_Notification_Delivery_Method `json:"notificationDeliveryMethod"`

	// NotificationDeliveryMethodId - no documentation
	NotificationDeliveryMethodId int `json:"notificationDeliveryMethodId"`

	// NotificationSubscriber - <nil>
	NotificationSubscriber *SoftLayer_Notification_Subscriber `json:"notificationSubscriber"`

	// NotificationSubscriberId - no documentation
	NotificationSubscriberId int `json:"notificationSubscriberId"`
}

func (softlayer_notification_subscriber_delivery_method *SoftLayer_Notification_Subscriber_Delivery_Method) String() string {
	return "SoftLayer_Notification_Subscriber_Delivery_Method"
}
