package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Subscriber - <nil>
type SoftLayer_Notification_Subscriber struct {

	// Active - <nil>
	Active int `json:"active"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// DeliveryMethodCount - no documentation
	DeliveryMethodCount uint64 `json:"deliveryMethodCount"`

	// DeliveryMethods - <nil>
	DeliveryMethods []*SoftLayer_Notification_Subscriber_Delivery_Method `json:"deliveryMethods"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Notification - <nil>
	Notification *SoftLayer_Notification `json:"notification"`

	// NotificationId - <nil>
	NotificationId int `json:"notificationId"`

	// NotificationSubscriberTypeId - <nil>
	NotificationSubscriberTypeId int `json:"notificationSubscriberTypeId"`

	// NotificationSubscriberTypeResourceId - <nil>
	NotificationSubscriberTypeResourceId int `json:"notificationSubscriberTypeResourceId"`
}
