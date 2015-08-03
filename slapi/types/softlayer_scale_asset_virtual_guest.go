package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Asset_Virtual_Guest - <nil>
type SoftLayer_Scale_Asset_Virtual_Guest struct {

	// VirtualGuest - no documentation
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId"`
}

func (softlayer_scale_asset_virtual_guest *SoftLayer_Scale_Asset_Virtual_Guest) String() string {
	return "SoftLayer_Scale_Asset_Virtual_Guest"
}

// CreateObject - <nil>
func (softlayer_scale_asset_virtual_guest *SoftLayer_Scale_Asset_Virtual_Guest) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Asset_Virtual_Guest) (*SoftLayer_Scale_Asset_Virtual_Guest, error) {
	var returnValue *SoftLayer_Scale_Asset_Virtual_Guest
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_asset_virtual_guest *SoftLayer_Scale_Asset_Virtual_Guest) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Asset_Virtual_Guest, error) {
	var returnValue *SoftLayer_Scale_Asset_Virtual_Guest
	return returnValue, nil
}
