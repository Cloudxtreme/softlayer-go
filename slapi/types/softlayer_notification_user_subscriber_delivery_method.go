package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_User_Subscriber_Delivery_Method - Provides mapping details of how the
// subscriber's notification will be delivered. This maps the subscriber's id with all the delivery
// method ids used to delivery the notification.
type SoftLayer_Notification_User_Subscriber_Delivery_Method struct {

	// NotificationMethodId - Unique identifier of the method used to deliver notification.
	NotificationMethodId int `json:"notificationMethodId,omitempty"`

	// NotificationUserSubscriberId - Unique identifier of the subscriber tied to the delivery method.
	NotificationUserSubscriberId int `json:"notificationUserSubscriberId,omitempty"`

	// Active - Determines if the delivery method is active for the user.
	Active int `json:"active,omitempty"`

	// DeliveryMethod - Provides details for the method used to deliver the notification (email, sms,
	// ticket).
	DeliveryMethod *SoftLayer_Notification_Delivery_Method `json:"deliveryMethod,omitempty"`

	// NotificationUserSubscriber - The Subscriber information tied to the delivery method.
	NotificationUserSubscriber *SoftLayer_Notification_User_Subscriber `json:"notificationUserSubscriber,omitempty"`
}

func (softlayer_notification_user_subscriber_delivery_method *SoftLayer_Notification_User_Subscriber_Delivery_Method) String() string {
	return "SoftLayer_Notification_User_Subscriber_Delivery_Method"
}
