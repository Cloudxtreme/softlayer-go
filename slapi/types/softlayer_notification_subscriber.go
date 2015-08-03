package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Subscriber - <nil>
type SoftLayer_Notification_Subscriber struct {

	// NotificationSubscriberTypeResourceId - <nil>
	NotificationSubscriberTypeResourceId int `json:"notificationSubscriberTypeResourceId"`

	// Active - <nil>
	Active int `json:"active"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// NotificationSubscriberTypeId - <nil>
	NotificationSubscriberTypeId int `json:"notificationSubscriberTypeId"`

	// Id - <nil>
	Id int `json:"id"`

	// NotificationId - <nil>
	NotificationId int `json:"notificationId"`
}

// SoftLayer_Notification_Subscriber_Extended is SoftLayer_Notification_Subscriber with all maskable types.
type SoftLayer_Notification_Subscriber_Extended struct {
	SoftLayer_Notification_Subscriber

	// DeliveryMethodCount - no documentation
	DeliveryMethodCount uint64 `json:"deliveryMethodCount"`

	// DeliveryMethods - <nil>
	DeliveryMethods []*SoftLayer_Notification_Subscriber_Delivery_Method `json:"deliveryMethods"`

	// Notification - <nil>
	Notification *SoftLayer_Notification `json:"notification"`
}

func (softlayer_notification_subscriber *SoftLayer_Notification_Subscriber) String() string {
	return "SoftLayer_Notification_Subscriber"
}
