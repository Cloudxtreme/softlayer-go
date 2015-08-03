package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Notification_Mobile - This is an extension of the SoftLayer_Notification class. These are
// implementation details specific to those notifications which can be subscribed to and received on a
// mobile device.
type SoftLayer_Notification_Mobile struct {
}

func (softlayer_notification_mobile *SoftLayer_Notification_Mobile) String() string {
	return "SoftLayer_Notification_Mobile"
}

// CreateSubscriberForMobileDevice - no documentation
func (softlayer_notification_mobile *SoftLayer_Notification_Mobile) CreateSubscriberForMobileDevice(ctx *slapi.RequestContext, keyName string, resourceTableId int, userRecordId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_notification_mobile *SoftLayer_Notification_Mobile) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Notification_Mobile, error) {
	var returnValue *SoftLayer_Notification_Mobile
	return returnValue, nil
}
