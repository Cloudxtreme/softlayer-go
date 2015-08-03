package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Asset - <nil>
type SoftLayer_Scale_Asset struct {

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_scale_asset *SoftLayer_Scale_Asset) String() string {
	return "SoftLayer_Scale_Asset"
}

// SoftLayer_Scale_Asset_Extended is SoftLayer_Scale_Asset with all maskable types.
type SoftLayer_Scale_Asset_Extended struct {
	SoftLayer_Scale_Asset

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`
}

func (softlayer_scale_asset *SoftLayer_Scale_Asset_Extended) String() string {
	return "SoftLayer_Scale_Asset"
}
