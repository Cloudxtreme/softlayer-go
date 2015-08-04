package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Property_Type - The storage property types provide standard definitions
// for properties which can be used with any type for Storage offering. The properties provide
// additional information about a volume which they are assigned to.
type SoftLayer_Network_Storage_Property_Type struct {

	// Keyname - no documentation
	Keyname string `json:"keyname,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Description - A type's description, for example 'Determines whether the volume is currently
	// mountable'.
	Description string `json:"description,omitempty"`
}

func (softlayer_network_storage_property_type *SoftLayer_Network_Storage_Property_Type) String() string {
	return "SoftLayer_Network_Storage_Property_Type"
}
