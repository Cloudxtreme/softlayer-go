package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Subscriber_Customer - <nil>
type SoftLayer_Notification_Subscriber_Customer struct {
}

func (softlayer_notification_subscriber_customer *SoftLayer_Notification_Subscriber_Customer) String() string {
	return "SoftLayer_Notification_Subscriber_Customer"
}

// SoftLayer_Notification_Subscriber_Customer_Extended is SoftLayer_Notification_Subscriber_Customer with all maskable types.
type SoftLayer_Notification_Subscriber_Customer_Extended struct {
	SoftLayer_Notification_Subscriber_Customer

	// SubscriberRecord - <nil>
	SubscriberRecord *SoftLayer_User_Customer `json:"subscriberRecord"`
}

func (softlayer_notification_subscriber_customer *SoftLayer_Notification_Subscriber_Customer_Extended) String() string {
	return "SoftLayer_Notification_Subscriber_Customer"
}
