package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Asset_Hardware - <nil>
type SoftLayer_Scale_Asset_Hardware struct {

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`
}

// CreateObject - <nil>
func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Scale_Asset_Hardware) (*SoftLayer_Scale_Asset_Hardware, error) {
	var returnValue *SoftLayer_Scale_Asset_Hardware
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_asset_hardware *SoftLayer_Scale_Asset_Hardware) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Scale_Asset_Hardware, error) {
	var returnValue *SoftLayer_Scale_Asset_Hardware
	return returnValue, nil
}
