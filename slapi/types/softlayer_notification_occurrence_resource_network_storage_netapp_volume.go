package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume - This type contains
// general information related to a [[SoftLayer_Network_Storage_NetApp_Volume]] resource that is
// impacted by a [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume struct {

	// Hostname - <nil>
	Hostname string `json:"hostname"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType"`
}

func (softlayer_notification_occurrence_resource_network_storage_netapp_volume *SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Storage_NetApp_Volume"
}
