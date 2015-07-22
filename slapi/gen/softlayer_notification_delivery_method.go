package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Delivery_Method - Provides details for the delivery methods available.
type SoftLayer_Notification_Delivery_Method struct {

	// Active - Determines if the delivery method is still used by the system.
	Active int `json:"active"`

	// Description - no documentation
	Description string `json:"description"`

	// Id - Unique identifier for the various notification delivery methods.
	Id int `json:"id"`

	// KeyName - Name that can be used by external systems to refer to delivery method.
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}
