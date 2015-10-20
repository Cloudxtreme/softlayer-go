package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Notification_Subscriber_Customer - <nil>
type SoftLayer_Notification_Subscriber_Customer struct {

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// NotificationId - <nil>
	NotificationId int `json:"notificationId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// NotificationSubscriberTypeId - <nil>
	NotificationSubscriberTypeId int `json:"notificationSubscriberTypeId,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// NotificationSubscriberTypeResourceId - <nil>
	NotificationSubscriberTypeResourceId int `json:"notificationSubscriberTypeResourceId,omitempty"`

	// Notification - <nil>
	Notification *SoftLayer_Notification `json:"notification,omitempty"`

	// DeliveryMethodCount - no documentation
	DeliveryMethodCount uint64 `json:"deliveryMethodCount,omitempty"`

	// SubscriberRecord - <nil>
	SubscriberRecord *SoftLayer_User_Customer `json:"subscriberRecord,omitempty"`

	// DeliveryMethods - <nil>
	DeliveryMethods []*SoftLayer_Notification_Subscriber_Delivery_Method `json:"deliveryMethods,omitempty"`
}

func (softlayer_notification_subscriber_customer *SoftLayer_Notification_Subscriber_Customer) String() string {
	return "SoftLayer_Notification_Subscriber_Customer"
}
