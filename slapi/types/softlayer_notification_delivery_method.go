package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Delivery_Method - Provides details for the delivery methods available.
type SoftLayer_Notification_Delivery_Method struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Id - Unique identifier for the various notification delivery methods.
	Id int `json:"id,omitempty"`

	// KeyName - Name that can be used by external systems to refer to delivery method.
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Active - Determines if the delivery method is still used by the system.
	Active int `json:"active,omitempty"`
}

func (softlayer_notification_delivery_method *SoftLayer_Notification_Delivery_Method) String() string {
	return "SoftLayer_Notification_Delivery_Method"
}
