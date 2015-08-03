package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Storage_Hub_Mezeo_File -
// SoftLayer_Container_Network_Storage_Hub_Mezeo_File provides specific details that only apply to
// files that are sent or received from CloudLayer storage resources.
type SoftLayer_Container_Network_Storage_Hub_Mezeo_File struct {

	// LockDuration - The amount of time a file is locked for writing, measured in seconds.
	LockDuration int `json:"lockDuration"`

	// LockId - A file can be locked at anytime to prevent the file from being overwritten by another copy.
	// When a file is locked, a lockId is provided to either unlock the file or renew the lock.
	LockId string `json:"lockId"`

	// PublishUrlCount - A file that is stored on the CloudLayer resource can be shared with the public at
	// anytime. When a file is shared, a token is generated to create one or more unique URLs.
	// ''PublishUrlCount'' provides the number of URLs that have been created for this file.
	PublishUrlCount string `json:"publishUrlCount"`
}

func (softlayer_container_network_storage_hub_mezeo_file *SoftLayer_Container_Network_Storage_Hub_Mezeo_File) String() string {
	return "SoftLayer_Container_Network_Storage_Hub_Mezeo_File"
}
