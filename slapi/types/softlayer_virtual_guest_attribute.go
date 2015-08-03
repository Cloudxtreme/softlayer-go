package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_Attribute - <nil>
type SoftLayer_Virtual_Guest_Attribute struct {

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_virtual_guest_attribute *SoftLayer_Virtual_Guest_Attribute) String() string {
	return "SoftLayer_Virtual_Guest_Attribute"
}

// SoftLayer_Virtual_Guest_Attribute_Extended is SoftLayer_Virtual_Guest_Attribute with all maskable types.
type SoftLayer_Virtual_Guest_Attribute_Extended struct {
	SoftLayer_Virtual_Guest_Attribute

	// Guest - <nil>
	Guest *SoftLayer_Virtual_Guest `json:"guest"`

	// Type - <nil>
	Type *SoftLayer_Virtual_Guest_Attribute_Type `json:"type"`
}

func (softlayer_virtual_guest_attribute *SoftLayer_Virtual_Guest_Attribute_Extended) String() string {
	return "SoftLayer_Virtual_Guest_Attribute"
}
