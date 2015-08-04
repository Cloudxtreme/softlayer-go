package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Subscriber - <nil>
type SoftLayer_Notification_Subscriber struct {

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// NotificationSubscriberTypeId - <nil>
	NotificationSubscriberTypeId int `json:"notificationSubscriberTypeId,omitempty"`

	// NotificationSubscriberTypeResourceId - <nil>
	NotificationSubscriberTypeResourceId int `json:"notificationSubscriberTypeResourceId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// NotificationId - <nil>
	NotificationId int `json:"notificationId,omitempty"`
}

func (softlayer_notification_subscriber *SoftLayer_Notification_Subscriber) String() string {
	return "SoftLayer_Notification_Subscriber"
}

// SoftLayer_Notification_Subscriber_Extended is SoftLayer_Notification_Subscriber with all maskable types.
type SoftLayer_Notification_Subscriber_Extended struct {
	SoftLayer_Notification_Subscriber

	// Notification - <nil>
	Notification *SoftLayer_Notification `json:"notification,omitempty"`

	// DeliveryMethodCount - no documentation
	DeliveryMethodCount uint64 `json:"deliveryMethodCount,omitempty"`

	// DeliveryMethods - <nil>
	DeliveryMethods []*SoftLayer_Notification_Subscriber_Delivery_Method `json:"deliveryMethods,omitempty"`
}

func (softlayer_notification_subscriber *SoftLayer_Notification_Subscriber_Extended) String() string {
	return "SoftLayer_Notification_Subscriber"
}
