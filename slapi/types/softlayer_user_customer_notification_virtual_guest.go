package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Customer_Notification_Virtual_Guest - The
// SoftLayer_User_Customer_Notification_Virtual_Guest object stores links between customers and the
// virtual guests they wish to monitor. This link is not enough, the user must be sure to also create
// SoftLayer_Network_Monitor_Version1_Query_Host instance with the response action set to "notify
// users" in order for the users linked to that hardware object to be notified on failure.
type SoftLayer_User_Customer_Notification_Virtual_Guest struct {

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// GuestId - The ID of the virtual guest object that is to be monitored.
	GuestId int `json:"guestId"`

	// Id - no documentation
	Id int `json:"id"`

	// User - The user that will be notified when the associated virtual guest object fails a monitoring
	// instance.
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - The ID of the SoftLayer_User_Customer object that represents the user to be notified on
	// monitoring failure.
	UserId int `json:"userId"`
}

// CreateObject - Passing in an unsaved instance of a SoftLayer_Customer_Notification_Virtual_Guest
// object into this function will create the object and return the results to the user.
func (softlayer_user_customer_notification_virtual_guest *SoftLayer_User_Customer_Notification_Virtual_Guest) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_User_Customer_Notification_Virtual_Guest) (*SoftLayer_User_Customer_Notification_Virtual_Guest, error) {
	var returnValue *SoftLayer_User_Customer_Notification_Virtual_Guest
	return returnValue, nil
}

// CreateObjects - Passing in a collection of unsaved instances of
// SoftLayer_Customer_Notification_Virtual_Guest objects into this function will create all objects and
// return the results to the user.
func (softlayer_user_customer_notification_virtual_guest *SoftLayer_User_Customer_Notification_Virtual_Guest) CreateObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_User_Customer_Notification_Virtual_Guest) ([]*SoftLayer_User_Customer_Notification_Virtual_Guest, error) {
	var returnValue []*SoftLayer_User_Customer_Notification_Virtual_Guest
	return returnValue, nil
}

// DeleteObjects - Like any other API object, the customer notification objects can be deleted by
// passing an instance of them into this function. The ID on the object must be set.
func (softlayer_user_customer_notification_virtual_guest *SoftLayer_User_Customer_Notification_Virtual_Guest) DeleteObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_User_Customer_Notification_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindByGuestId - This method returns all SoftLayer_User_Customer_Notification_Virtual_Guest objects
// associated with the passed in ID as long as that hardware ID is owned by the current user's account.
// This behavior can also be accomplished by simply tapping monitoringUserNotification on the
// Virtual_Guest object.
func (softlayer_user_customer_notification_virtual_guest *SoftLayer_User_Customer_Notification_Virtual_Guest) FindByGuestId(commonOptions *slapi.CommonOptions, id int) ([]*SoftLayer_User_Customer_Notification_Virtual_Guest, error) {
	var returnValue []*SoftLayer_User_Customer_Notification_Virtual_Guest
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_User_Customer_Notification_Virtual_Guest object whose
// ID number corresponds to the ID number of the init parameter passed to the
// SoftLayer_User_Customer_Notification_Virtual_Guest service. You can only retrieve guest
// notifications attached to virtual guests and users that belong to your account
func (softlayer_user_customer_notification_virtual_guest *SoftLayer_User_Customer_Notification_Virtual_Guest) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_Customer_Notification_Virtual_Guest, error) {
	var returnValue *SoftLayer_User_Customer_Notification_Virtual_Guest
	return returnValue, nil
}