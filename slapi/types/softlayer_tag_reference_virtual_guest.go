package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Tag_Reference_Virtual_Guest - <nil>
type SoftLayer_Tag_Reference_Virtual_Guest struct {
}

func (softlayer_tag_reference_virtual_guest *SoftLayer_Tag_Reference_Virtual_Guest) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest"
}

// SoftLayer_Tag_Reference_Virtual_Guest_Extended is SoftLayer_Tag_Reference_Virtual_Guest with all maskable types.
type SoftLayer_Tag_Reference_Virtual_Guest_Extended struct {
	SoftLayer_Tag_Reference_Virtual_Guest

	// Resource - <nil>
	Resource *SoftLayer_Virtual_Guest `json:"resource"`
}

func (softlayer_tag_reference_virtual_guest *SoftLayer_Tag_Reference_Virtual_Guest_Extended) String() string {
	return "SoftLayer_Tag_Reference_Virtual_Guest"
}
