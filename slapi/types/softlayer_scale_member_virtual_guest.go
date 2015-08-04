package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Member_Virtual_Guest - <nil>
type SoftLayer_Scale_Member_Virtual_Guest struct {

	// ScaleGroupId - The identifier of the group this member belongs to.
	ScaleGroupId int `json:"scaleGroupId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_scale_member_virtual_guest *SoftLayer_Scale_Member_Virtual_Guest) String() string {
	return "SoftLayer_Scale_Member_Virtual_Guest"
}

// SoftLayer_Scale_Member_Virtual_Guest_Extended is SoftLayer_Scale_Member_Virtual_Guest with all maskable types.
type SoftLayer_Scale_Member_Virtual_Guest_Extended struct {
	SoftLayer_Scale_Member_Virtual_Guest

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`

	// VirtualGuest - no documentation
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`
}

func (softlayer_scale_member_virtual_guest *SoftLayer_Scale_Member_Virtual_Guest_Extended) String() string {
	return "SoftLayer_Scale_Member_Virtual_Guest"
}
