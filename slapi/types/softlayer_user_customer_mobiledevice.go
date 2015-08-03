package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_MobileDevice - This class represents a mobile device belonging to a user.
// The device can be a phone, tablet, or possibly even some Android based netbooks. The purpose is to
// tie just enough info with the device and the user to enable push notifications through non-softlayer
// entities (Google, Apple,
type SoftLayer_User_Customer_MobileDevice struct {

	// AvailablePushNotificationSubscriptionCount - A count of notification subscriptions available to a
	// mobile device.
	AvailablePushNotificationSubscriptionCount uint64 `json:"availablePushNotificationSubscriptionCount"`

	// AvailablePushNotificationSubscriptions - Notification subscriptions available to a mobile device.
	AvailablePushNotificationSubscriptions []*SoftLayer_Notification `json:"availablePushNotificationSubscriptions"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`

	// DisplayResolutionXxY - no documentation
	DisplayResolutionXxY string `json:"displayResolutionXxY"`

	// Id - no documentation
	Id int `json:"id"`

	// MobileDeviceTypeId - no documentation
	MobileDeviceTypeId int `json:"mobileDeviceTypeId"`

	// MobileOperatingSystemId - no documentation
	MobileOperatingSystemId int `json:"mobileOperatingSystemId"`

	// ModelNumber - no documentation
	ModelNumber string `json:"modelNumber"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// OperatingSystem - <nil>
	OperatingSystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem `json:"operatingSystem"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber"`

	// PushNotificationSubscriptionCount - A count of notification subscriptions attached to a mobile
	// device.
	PushNotificationSubscriptionCount uint64 `json:"pushNotificationSubscriptionCount"`

	// PushNotificationSubscriptions - Notification subscriptions attached to a mobile device.
	PushNotificationSubscriptions []*SoftLayer_Notification_User_Subscriber `json:"pushNotificationSubscriptions"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber"`

	// Token - no documentation
	Token string `json:"token"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_MobileDevice_Type `json:"type"`

	// UserId - no documentation
	UserId int `json:"userId"`
}

func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) String() string {
	return "SoftLayer_User_Customer_MobileDevice"
}
