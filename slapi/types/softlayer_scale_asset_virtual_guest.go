package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Asset_Virtual_Guest - <nil>
type SoftLayer_Scale_Asset_Virtual_Guest struct {

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`

	// VirtualGuest - no documentation
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// VirtualGuestId - no documentation
	VirtualGuestId int `json:"virtualGuestId,omitempty"`
}

func (softlayer_scale_asset_virtual_guest *SoftLayer_Scale_Asset_Virtual_Guest) String() string {
	return "SoftLayer_Scale_Asset_Virtual_Guest"
}
