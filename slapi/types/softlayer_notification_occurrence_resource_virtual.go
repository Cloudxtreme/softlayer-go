package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Notification_Occurrence_Resource_Virtual - This type contains general information related
// to a [[SoftLayer_Virtual_Guest]] resource that is impacted by a
// [[SoftLayer_Notification_Occurrence_Event]].
type SoftLayer_Notification_Occurrence_Resource_Virtual struct {

	// Hostname - <nil>
	Hostname string `json:"hostname"`

	// PrivateIp - <nil>
	PrivateIp string `json:"privateIp"`

	// PublicIp - <nil>
	PublicIp string `json:"publicIp"`

	// ResourceType - <nil>
	ResourceType string `json:"resourceType"`
}

func (softlayer_notification_occurrence_resource_virtual *SoftLayer_Notification_Occurrence_Resource_Virtual) String() string {
	return "SoftLayer_Notification_Occurrence_Resource_Virtual"
}
