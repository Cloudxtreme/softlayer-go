package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Scale_Member_Virtual_Guest - <nil>
type SoftLayer_Scale_Member_Virtual_Guest struct {
}

// SoftLayer_Scale_Member_Virtual_Guest_Extended is SoftLayer_Scale_Member_Virtual_Guest with all maskable types.
type SoftLayer_Scale_Member_Virtual_Guest_Extended struct {
	SoftLayer_Scale_Member_Virtual_Guest

	// VirtualGuest - no documentation
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId"`
}

func (softlayer_scale_member_virtual_guest *SoftLayer_Scale_Member_Virtual_Guest) String() string {
	return "SoftLayer_Scale_Member_Virtual_Guest"
}
