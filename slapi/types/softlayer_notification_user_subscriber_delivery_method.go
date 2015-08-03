package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_User_Subscriber_Delivery_Method - Provides mapping details of how the
// subscriber's notification will be delivered. This maps the subscriber's id with all the delivery
// method ids used to delivery the notification.
type SoftLayer_Notification_User_Subscriber_Delivery_Method struct {

	// Active - Determines if the delivery method is active for the user.
	Active int `json:"active"`

	// NotificationMethodId - Unique identifier of the method used to deliver notification.
	NotificationMethodId int `json:"notificationMethodId"`

	// NotificationUserSubscriberId - Unique identifier of the subscriber tied to the delivery method.
	NotificationUserSubscriberId int `json:"notificationUserSubscriberId"`
}

func (softlayer_notification_user_subscriber_delivery_method *SoftLayer_Notification_User_Subscriber_Delivery_Method) String() string {
	return "SoftLayer_Notification_User_Subscriber_Delivery_Method"
}

// SoftLayer_Notification_User_Subscriber_Delivery_Method_Extended is SoftLayer_Notification_User_Subscriber_Delivery_Method with all maskable types.
type SoftLayer_Notification_User_Subscriber_Delivery_Method_Extended struct {
	SoftLayer_Notification_User_Subscriber_Delivery_Method

	// DeliveryMethod - Provides details for the method used to deliver the notification (email, sms,
	// ticket).
	DeliveryMethod *SoftLayer_Notification_Delivery_Method `json:"deliveryMethod"`

	// NotificationUserSubscriber - The Subscriber information tied to the delivery method.
	NotificationUserSubscriber *SoftLayer_Notification_User_Subscriber `json:"notificationUserSubscriber"`
}

func (softlayer_notification_user_subscriber_delivery_method *SoftLayer_Notification_User_Subscriber_Delivery_Method_Extended) String() string {
	return "SoftLayer_Notification_User_Subscriber_Delivery_Method"
}
