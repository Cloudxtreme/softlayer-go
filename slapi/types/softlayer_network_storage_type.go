package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Type - The SoftLayer_Network_Storage_Type contains a description of the
// associated SoftLayer_Network_Storage object.
type SoftLayer_Network_Storage_Type struct {

	// Description - Human readable description for the associated SoftLayer_Network_Storage object.
	Description string `json:"description,omitempty"`

	// Id - ID which corresponds with storageTypeId on storage objects.
	Id int `json:"id,omitempty"`

	// KeyName - Machine readable description code for the associated SoftLayer_Network_Storage object.
	KeyName string `json:"keyName,omitempty"`

	// VolumeCount - A count of the SoftLayer_Network_Storage object that uses this type.
	VolumeCount uint64 `json:"volumeCount,omitempty"`

	// Volumes - The SoftLayer_Network_Storage object that uses this type.
	Volumes []*SoftLayer_Network_Storage `json:"volumes,omitempty"`
}

func (softlayer_network_storage_type *SoftLayer_Network_Storage_Type) String() string {
	return "SoftLayer_Network_Storage_Type"
}
