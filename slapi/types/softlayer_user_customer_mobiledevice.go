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

	// MobileOperatingSystemId - no documentation
	MobileOperatingSystemId int `json:"mobileOperatingSystemId,omitempty"`

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`

	// SerialNumber - no documentation
	SerialNumber string `json:"serialNumber,omitempty"`

	// MobileDeviceTypeId - no documentation
	MobileDeviceTypeId int `json:"mobileDeviceTypeId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// PhoneNumber - no documentation
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// ModelNumber - no documentation
	ModelNumber string `json:"modelNumber,omitempty"`

	// Token - no documentation
	Token string `json:"token,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DisplayResolutionXxY - no documentation
	DisplayResolutionXxY string `json:"displayResolutionXxY,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) String() string {
	return "SoftLayer_User_Customer_MobileDevice"
}

// SoftLayer_User_Customer_MobileDevice_Extended is SoftLayer_User_Customer_MobileDevice with all maskable types.
type SoftLayer_User_Customer_MobileDevice_Extended struct {
	SoftLayer_User_Customer_MobileDevice

	// PushNotificationSubscriptions - Notification subscriptions attached to a mobile device.
	PushNotificationSubscriptions []*SoftLayer_Notification_User_Subscriber `json:"pushNotificationSubscriptions,omitempty"`

	// AvailablePushNotificationSubscriptionCount - A count of notification subscriptions available to a
	// mobile device.
	AvailablePushNotificationSubscriptionCount uint64 `json:"availablePushNotificationSubscriptionCount,omitempty"`

	// PushNotificationSubscriptionCount - A count of notification subscriptions attached to a mobile
	// device.
	PushNotificationSubscriptionCount uint64 `json:"pushNotificationSubscriptionCount,omitempty"`

	// Customer - <nil>
	Customer *SoftLayer_User_Customer `json:"customer,omitempty"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_MobileDevice_Type `json:"type,omitempty"`

	// AvailablePushNotificationSubscriptions - Notification subscriptions available to a mobile device.
	AvailablePushNotificationSubscriptions []*SoftLayer_Notification `json:"availablePushNotificationSubscriptions,omitempty"`

	// OperatingSystem - <nil>
	OperatingSystem *SoftLayer_User_Customer_MobileDevice_OperatingSystem `json:"operatingSystem,omitempty"`
}

func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice_Extended) String() string {
	return "SoftLayer_User_Customer_MobileDevice"
}
