package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Asset_Hardware - <nil>
type SoftLayer_Scale_Asset_Hardware struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`
}

func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) String() string {
	return "SoftLayer_Scale_Asset_Hardware"
}
