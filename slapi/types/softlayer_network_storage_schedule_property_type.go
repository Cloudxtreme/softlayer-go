package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Schedule_Property_Type - A schedule property type is used to allow for a
// standardized method of defining network storage schedules.
type SoftLayer_Network_Storage_Schedule_Property_Type struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Keyname - no documentation
	Keyname string `json:"keyname,omitempty"`

	// Name - A schedule property type's name, for example 'Start Date'.
	Name string `json:"name,omitempty"`

	// NasType - The type of Storage volume type which a property type may be associated with.
	NasType string `json:"nasType,omitempty"`

	// Description - A type's description, for example 'Date for the schedule to start.'.
	Description string `json:"description,omitempty"`
}

func (softlayer_network_storage_schedule_property_type *SoftLayer_Network_Storage_Schedule_Property_Type) String() string {
	return "SoftLayer_Network_Storage_Schedule_Property_Type"
}
