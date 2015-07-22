package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// CreateObject - Create a new mobile device association for a user.
func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_User_Customer_MobileDevice) (*SoftLayer_User_Customer_MobileDevice, error) {
	var returnValue *SoftLayer_User_Customer_MobileDevice
	return returnValue, nil
}

// DeleteObject - no documentation
func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit the object by passing in a modified instance of the object
func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_User_Customer_MobileDevice) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_mobiledevice *SoftLayer_User_Customer_MobileDevice) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_Customer_MobileDevice, error) {
	var returnValue *SoftLayer_User_Customer_MobileDevice
	return returnValue, nil
}
