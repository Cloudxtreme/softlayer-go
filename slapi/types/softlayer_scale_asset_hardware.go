package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Asset_Hardware - <nil>
type SoftLayer_Scale_Asset_Hardware struct {

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`
}

func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) String() string {
	return "SoftLayer_Scale_Asset_Hardware"
}

// CreateObject - <nil>
func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Asset_Hardware) (*SoftLayer_Scale_Asset_Hardware, error) {
	var returnValue *SoftLayer_Scale_Asset_Hardware
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Asset_Hardware, error) {
	var returnValue *SoftLayer_Scale_Asset_Hardware
	return returnValue, nil
}
