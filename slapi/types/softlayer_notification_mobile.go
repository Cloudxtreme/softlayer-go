package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Notification_Mobile - This is an extension of the SoftLayer_Notification class. These are
// implementation details specific to those notifications which can be subscribed to and received on a
// mobile device.
type SoftLayer_Notification_Mobile struct {
}

// CreateSubscriberForMobileDevice - no documentation
func (softlayer_notification_mobile *SoftLayer_Notification_Mobile) CreateSubscriberForMobileDevice(commonOptions *slapi.CommonOptions, keyName string, resourceTableId int, userRecordId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_mobile *SoftLayer_Notification_Mobile) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Notification_Mobile, error) {
	var returnValue *SoftLayer_Notification_Mobile
	return returnValue, nil
}
