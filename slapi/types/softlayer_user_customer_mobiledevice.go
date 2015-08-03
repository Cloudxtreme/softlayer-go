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

	// UserId - no documentation
	UserId int `json:"userId"`

	// MobileDeviceTypeId - no documentation
	MobileDeviceTypeId int `json:"mobileDeviceTypeId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber"`

	// DisplayResolutionXxY - no documentation
	DisplayResolutionXxY string `json:"displayResolutionXxY"`

	// MobileOperatingSystemId - no documentation
	MobileOperatingSystemId int `json:"mobileOperatingSystemId"`

	// ModelNumber - no documentation
	ModelNumber string `json:"modelNumber"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber"`

	// Token - no documentation
	Token string `json:"token"`
}

func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) String() string {
	return "SoftLayer_User_Customer_MobileDevice"
}

// SoftLayer_User_Customer_MobileDevice_Extended is SoftLayer_User_Customer_MobileDevice with all maskable types.
type SoftLayer_User_Customer_MobileDevice_Extended struct {
	SoftLayer_User_Customer_MobileDevice

	// Type - <nil>
	Type *SoftLayer_User_Customer_MobileDevice_Type `json:"type"`

	// PushNotificationSubscriptionCount - A count of notification subscriptions attached to a mobile
	// device.
	PushNotificationSubscriptionCount uint64 `json:"pushNotificationSubscriptionCount"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer"`

	// OperatingSystem - <nil>
	OperatingSystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem `json:"operatingSystem"`

	// AvailablePushNotificationSubscriptionCount - A count of notification subscriptions available to a
	// mobile device.
	AvailablePushNotificationSubscriptionCount uint64 `json:"availablePushNotificationSubscriptionCount"`

	// AvailablePushNotificationSubscriptions - Notification subscriptions available to a mobile device.
	AvailablePushNotificationSubscriptions []*SoftLayer_Notification `json:"availablePushNotificationSubscriptions"`

	// PushNotificationSubscriptions - Notification subscriptions attached to a mobile device.
	PushNotificationSubscriptions []*SoftLayer_Notification_User_Subscriber `json:"pushNotificationSubscriptions"`
}

func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice_Extended) String() string {
	return "SoftLayer_User_Customer_MobileDevice"
}
