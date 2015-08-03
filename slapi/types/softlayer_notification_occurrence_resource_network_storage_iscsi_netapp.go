package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp - This type contains general
// information related to a [[SoftLayer_Network_Storage_Iscsi_NetApp]] resource that is impacted by a
// [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp struct {

	// Hostname - <nil>
	Hostname string `json:"hostname"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType"`
}

func (softlayer_notification_occurrence_resource_network_storage_iscsi_netapp *SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Network_Storage_Iscsi_NetApp"
}
