package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Attribute - <nil>
type SoftLayer_Virtual_Guest_Attribute struct {

	// Guest - <nil>
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// Type - <nil>
	Type *SoftLayer_Virtual_Guest_Attribute_Type `json:"type"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_virtual_guest_attribute *SoftLayer_Virtual_Guest_Attribute) String() string {
	return "SoftLayer_Virtual_Guest_Attribute"
}
