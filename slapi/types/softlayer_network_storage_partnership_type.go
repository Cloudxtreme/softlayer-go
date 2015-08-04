package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Partnership_Type - A network storage partnership type is used to define
// the link between two volumes.
type SoftLayer_Network_Storage_Partnership_Type struct {

	// Description - A type's description, for example snapshot partnership'.
	Description string `json:"description,omitempty"`

	// Keyname - no documentation
	Keyname string `json:"keyname,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_network_storage_partnership_type *SoftLayer_Network_Storage_Partnership_Type) String() string {
	return "SoftLayer_Network_Storage_Partnership_Type"
}
