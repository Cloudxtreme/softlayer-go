package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_MobileDevice - This class represents a mobile device belonging to a user.
// The device can be a phone, tablet, or possibly even some Android based net books. The purpose is to
// tie just enough info with the device and the user to enable push notifications through non-softlayer
// entities (Google, Apple,
type SoftLayer_User_Customer_MobileDevice struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MobileOperatingSystemId - no documentation
	MobileOperatingSystemId int `json:"mobileOperatingSystemId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber,omitempty"`

	// ModelNumber - no documentation
	ModelNumber string `json:"modelNumber,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Token - no documentation
	Token string `json:"token,omitempty"`

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DisplayResolutionXxY - no documentation
	DisplayResolutionXxY string `json:"displayResolutionXxY,omitempty"`

	// MobileDeviceTypeId - no documentation
	MobileDeviceTypeId int `json:"mobileDeviceTypeId,omitempty"`

	// AvailablePushNotificationSubscriptions - Notification subscriptions available to a mobile device.
	AvailablePushNotificationSubscriptions []*SoftLayer_Notification `json:"availablePushNotificationSubscriptions,omitempty"`

	// AvailablePushNotificationSubscriptionCount - A count of notification subscriptions available to a
	// mobile device.
	AvailablePushNotificationSubscriptionCount uint64 `json:"availablePushNotificationSubscriptionCount,omitempty"`

	// PushNotificationSubscriptionCount - A count of notification subscriptions attached to a mobile
	// device.
	PushNotificationSubscriptionCount uint64 `json:"pushNotificationSubscriptionCount,omitempty"`

	// PushNotificationSubscriptions - Notification subscriptions attached to a mobile device.
	PushNotificationSubscriptions []*SoftLayer_Notification_User_Subscriber `json:"pushNotificationSubscriptions,omitempty"`

	// OperatingSystem - no documentation
	OperatingSystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem `json:"operatingSystem,omitempty"`

	// Type - no documentation
	Type *SoftLayer_User_Customer_MobileDevice_Type `json:"type,omitempty"`

	// Customer - no documentation
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`
}

func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) String() string {
	return "SoftLayer_User_Customer_MobileDevice"
}
