package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume - This type contains
// general information related to a [[SoftLayer_Network_Storage_NetApp_Volume]] resource that is
// impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp,omitempty"`

	// ResourceAccountId - <<< EOT The unique identifier for the [[SoftLayer_Account]] associated with
	ResourceAccountId int `json:"resourceAccountId,omitempty"`

	// Active - <nil>
	Active int `json:"active,omitempty"`

	// FilterLabel - <<< EOT A label which gives some background as to what piece of
	FilterLabel string `json:"filterLabel,omitempty"`

	// NotificationOccurrenceEventId - no documentation
	NotificationOccurrenceEventId int `json:"notificationOccurrenceEventId,omitempty"`

	// Hostname - <nil>
	Hostname string `json:"hostname,omitempty"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType,omitempty"`

	// ResourceName - <nil>
	ResourceName string `json:"resourceName,omitempty"`

	// ResourceTableId - <<< EOT The unique identifier for the physical resource that is associated
	ResourceTableId int `json:"resourceTableId,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_storage_netapp_volume *SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume"
}

// SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume_Extended is SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume with all maskable types.
type SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume_Extended struct {
	SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume

	// NotificationOccurrenceEvent - no documentation
	NotificationOccurrenceEvent *SoftLayer_Notification_Occurrence_Event `json:"notificationOccurrenceEvent,omitempty"`

	// Resource - no documentation
	Resource *SoftLayer_Entity `json:"resource,omitempty"`
}

func (softlayer_notification_occurrence_resource_network_storage_netapp_volume *SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume_Extended) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume"
}
