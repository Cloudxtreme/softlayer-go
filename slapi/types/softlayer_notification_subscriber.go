package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Subscriber - <nil>
type SoftLayer_Notification_Subscriber struct {

	// Id - <nil>
	Id int `json:"id"`

	// NotificationId - <nil>
	NotificationId int `json:"notificationId"`

	// NotificationSubscriberTypeId - <nil>
	NotificationSubscriberTypeId int `json:"notificationSubscriberTypeId"`

	// NotificationSubscriberTypeResourceId - <nil>
	NotificationSubscriberTypeResourceId int `json:"notificationSubscriberTypeResourceId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Active - <nil>
	Active int `json:"active"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_notification_subscriber *SoftLayer_Notification_Subscriber) String() string {
	return "SoftLayer_Notification_Subscriber"
}

// SoftLayer_Notification_Subscriber_Extended is SoftLayer_Notification_Subscriber with all maskable types.
type SoftLayer_Notification_Subscriber_Extended struct {
	SoftLayer_Notification_Subscriber

	// DeliveryMethods - <nil>
	DeliveryMethods []*SoftLayer_Notification_Subscriber_Delivery_Method `json:"deliveryMethods"`

	// Notification - <nil>
	Notification *SoftLayer_Notification `json:"notification"`

	// DeliveryMethodCount - no documentation
	DeliveryMethodCount uint64 `json:"deliveryMethodCount"`
}

func (softlayer_notification_subscriber *SoftLayer_Notification_Subscriber_Extended) String() string {
	return "SoftLayer_Notification_Subscriber"
}
