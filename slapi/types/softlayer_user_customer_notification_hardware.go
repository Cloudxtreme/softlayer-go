package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Notification_Hardware - The Customer_Notification_Hardware object stores
// links between customers and the hardware devices they wish to monitor. This link is not enough, the
// user must be sure to also create SoftLayer_Network_Monitor_Version1_Query_Host instance with the
// response action set to "notify users" in order for the users linked to that hardware object to be
// notified on failure.
type SoftLayer_User_Customer_Notification_Hardware struct {

	// HardwareId - The ID of the Hardware object that is to be monitored.
	HardwareId int `json:"hardwareId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// UserId - The ID of the SoftLayer_User_Customer object that represents the user to be notified on
	// monitoring failure.
	UserId int `json:"userId,omitempty"`
}

func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware) String() string {
	return "SoftLayer_User_Customer_Notification_Hardware"
}

// SoftLayer_User_Customer_Notification_Hardware_Extended is SoftLayer_User_Customer_Notification_Hardware with all maskable types.
type SoftLayer_User_Customer_Notification_Hardware_Extended struct {
	SoftLayer_User_Customer_Notification_Hardware

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// User - The user that will be notified when the associated hardware object fails a monitoring
	// instance.
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_user_customer_notification_hardware *SoftLayer_User_Customer_Notification_Hardware_Extended) String() string {
	return "SoftLayer_User_Customer_Notification_Hardware"
}
