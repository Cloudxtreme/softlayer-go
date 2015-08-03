package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_User_Subscriber_Resource - Retrieve identifier cross-reference information.
// SoftLayer_Notification_User_Subscriber_Resource provides the resource table id and subscriber id
// relation. The resource table id is the id of the service the subscriber receives alerts for. This
// resource table id could be the unique identifier for a Storage Evault service, Global Load Balancer
// or CDN service.
type SoftLayer_Notification_User_Subscriber_Resource struct {

	// NotificationUserSubscriberId - Unique identifier of the subscriber that will receive the alerts for
	// the resource subscribed to a notification.
	NotificationUserSubscriberId int `json:"notificationUserSubscriberId"`

	// ResourceTableId - Unique identifier for a SoftLayer service that is subscribed to a notification.
	// Currently, the SoftLayer services that can be subscribed to notifications are: Storage EVault CDN
	// Global Load Balancer
	ResourceTableId int `json:"resourceTableId"`
}

func (softlayer_notification_user_subscriber_resource *SoftLayer_Notification_User_Subscriber_Resource) String() string {
	return "SoftLayer_Notification_User_Subscriber_Resource"
}

// SoftLayer_Notification_User_Subscriber_Resource_Extended is SoftLayer_Notification_User_Subscriber_Resource with all maskable types.
type SoftLayer_Notification_User_Subscriber_Resource_Extended struct {
	SoftLayer_Notification_User_Subscriber_Resource

	// NotificationUserSubscriber - The Subscriber information tied to the resource service.
	NotificationUserSubscriber *SoftLayer_Notification_User_Subscriber `json:"notificationUserSubscriber"`
}

func (softlayer_notification_user_subscriber_resource *SoftLayer_Notification_User_Subscriber_Resource_Extended) String() string {
	return "SoftLayer_Notification_User_Subscriber_Resource"
}
