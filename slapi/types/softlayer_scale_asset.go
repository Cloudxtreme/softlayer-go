package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Asset - <nil>
type SoftLayer_Scale_Asset struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId"`
}

func (softlayer_scale_asset *SoftLayer_Scale_Asset) String() string {
	return "SoftLayer_Scale_Asset"
}

// DeleteObject - <nil>
func (softlayer_scale_asset *SoftLayer_Scale_Asset) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_asset *SoftLayer_Scale_Asset) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Asset, error) {
	var returnValue *SoftLayer_Scale_Asset
	return returnValue, nil
}
