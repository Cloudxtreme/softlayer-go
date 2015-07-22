package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Member_Virtual_Guest - <nil>
type SoftLayer_Scale_Member_Virtual_Guest struct {

	// VirtualGuest - no documentation
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId"`
}

// DeleteObject - <nil>
func (softlayer_scale_member_virtual_guest *SoftLayer_Scale_Member_Virtual_Guest) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_member_virtual_guest *SoftLayer_Scale_Member_Virtual_Guest) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Member_Virtual_Guest, error) {
	var returnValue *SoftLayer_Scale_Member_Virtual_Guest
	return returnValue, nil
}
