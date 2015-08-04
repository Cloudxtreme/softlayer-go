package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp - This type contains general
// information related to a [[SoftLayer_Network_Storage_Iscsi_NetApp]] resource that is impacted by a
// [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_storage_iscsi_netapp *SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp"
}

// SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp_Extended is SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp with all maskable types.
type SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp_Extended struct {
	SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_storage_iscsi_netapp *SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp"
}
