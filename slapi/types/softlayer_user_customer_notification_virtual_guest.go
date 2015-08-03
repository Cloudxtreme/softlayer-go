package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Notification_Virtual_Guest - The
// SoftLayer_User_Customer_Notification_Virtual_Guest object stores links between customers and the
// virtual guests they wish to monitor. This link is not enough, the user must be sure to also create
// SoftLayer_Network_Monitor_Version1_Query_Host instance with the response action set to "notify
// users" in order for the users linked to that hardware object to be notified on failure.
type SoftLayer_User_Customer_Notification_Virtual_Guest struct {

	// GuestId - The ID of the virtual guest object that is to be monitored.
	GuestId int `json:"guestId"`

	// Id - no documentation
	Id int `json:"id"`

	// UserId - The ID of the SoftLayer_User_Customer object that represents the user to be notified on
	// monitoring failure.
	UserId int `json:"userId"`
}

// SoftLayer_User_Customer_Notification_Virtual_Guest_Extended is SoftLayer_User_Customer_Notification_Virtual_Guest with all maskable types.
type SoftLayer_User_Customer_Notification_Virtual_Guest_Extended struct {
	SoftLayer_User_Customer_Notification_Virtual_Guest

	// Guest - no documentation
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// User - The user that will be notified when the associated virtual guest object fails a monitoring
	// instance.
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_notification_virtual_guest *SoftLayer_User_Customer_Notification_Virtual_Guest) String() string {
	return "SoftLayer_User_Customer_Notification_Virtual_Guest"
}
