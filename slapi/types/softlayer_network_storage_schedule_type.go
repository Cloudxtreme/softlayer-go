package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Schedule_Type - A schedule type is used to define what a schedule was
// created to do. When creating a schedule to take snapshots of a volume, the 'Snapshot' schedule type
// would be used.
type SoftLayer_Network_Storage_Schedule_Type struct {

	// Keyname - no documentation
	Keyname string `json:"keyname,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_network_storage_schedule_type *SoftLayer_Network_Storage_Schedule_Type) String() string {
	return "SoftLayer_Network_Storage_Schedule_Type"
}
