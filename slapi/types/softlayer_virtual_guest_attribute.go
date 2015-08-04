package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Attribute - <nil>
type SoftLayer_Virtual_Guest_Attribute struct {

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// Guest - <nil>
	Guest *SoftLayer_Virtual_Guest `json:"guest,omitempty"`

	// Type - <nil>
	Type *SoftLayer_Virtual_Guest_Attribute_Type `json:"type,omitempty"`
}

func (softlayer_virtual_guest_attribute *SoftLayer_Virtual_Guest_Attribute) String() string {
	return "SoftLayer_Virtual_Guest_Attribute"
}
