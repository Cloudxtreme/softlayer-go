package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Layout_Container_Type - The SoftLayer_Layout_Container_Type contains definitions for
// container types
type SoftLayer_Layout_Container_Type struct {

	// Keyname - The unique key name of the container type, used primarily for programmatic purposes
	Keyname string `json:"keyname,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_layout_container_type *SoftLayer_Layout_Container_Type) String() string {
	return "SoftLayer_Layout_Container_Type"
}
