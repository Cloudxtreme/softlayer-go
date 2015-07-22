package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Notification_Hardware - The Customer_Notification_Hardware object stores
// links between customers and the hardware devices they wish to monitor. This link is not enough, the
// user must be sure to also create SoftLayer_Network_Monitor_Version1_Query_Host instance with the
// response action set to "notify users" in order for the users linked to that hardware object to be
// notified on failure.
type SoftLayer_User_Customer_Notification_Hardware struct {

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The ID of the Hardware object that is to be monitored.
	HardwareId int `json:"hardwareId"`

	// Id - no documentation
	Id int `json:"id"`

	// User - The user that will be notified when the associated hardware object fails a monitoring
	// instance.
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - The ID of the SoftLayer_User_Customer object that represents the user to be notified on
	// monitoring failure.
	UserId int `json:"userId"`
}

// CreateObject - Passing in an unsaved instances of a Customer_Notification_Hardware object into this
// function will create the object and return the results to the user.
func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware) CreateObject(templateObject SoftLayer_User_Customer_Notification_Hardware) (*SoftLayer_User_Customer_Notification_Hardware, error) {
	var returnValue *SoftLayer_User_Customer_Notification_Hardware
	return returnValue, nil
}

// CreateObjects - Passing in a collection of unsaved instances of Customer_Notification_Hardware
// objects into this function will create all objects and return the results to the user.
func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware) CreateObjects(templateObjects []SoftLayer_User_Customer_Notification_Hardware) ([]*SoftLayer_Dns_Domain, error) {
	var returnValue []*SoftLayer_Dns_Domain
	return returnValue, nil
}

// DeleteObjects - Like any other API object, the customer notification objects can be deleted by
// passing an instance of them into this function. The ID on the object must be set.
func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware) DeleteObjects(templateObjects []SoftLayer_User_Customer_Notification_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindByHardwareId - This method returns all Customer_Notification_Hardware objects associated with
// the passed in hardware ID as long as that hardware ID is owned by the current user's account. This
// behavior can also be accomplished by simply tapping monitoringUserNotification on the
// Hardware_Server object.
func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware) FindByHardwareId(hardwareId int) ([]*SoftLayer_User_Customer_Notification_Hardware, error) {
	var returnValue []*SoftLayer_User_Customer_Notification_Hardware
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_User_Customer_Notification_Hardware object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_User_Customer_Notification_Hardware service. You can only retrieve hardware notifications
// attached to hardware and users that belong to your account
func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware) GetObject() (*SoftLayer_User_Customer_Notification_Hardware, error) {
	var returnValue *SoftLayer_User_Customer_Notification_Hardware
	return returnValue, nil
}
