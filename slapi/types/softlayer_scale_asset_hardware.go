package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Asset_Hardware - <nil>
type SoftLayer_Scale_Asset_Hardware struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId,omitempty"`
}

func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) String() string {
	return "SoftLayer_Scale_Asset_Hardware"
}

// SoftLayer_Scale_Asset_Hardware_Extended is SoftLayer_Scale_Asset_Hardware with all maskable types.
type SoftLayer_Scale_Asset_Hardware_Extended struct {
	SoftLayer_Scale_Asset_Hardware

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`
}

func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware_Extended) String() string {
	return "SoftLayer_Scale_Asset_Hardware"
}
