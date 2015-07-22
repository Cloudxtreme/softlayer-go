package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Policy_Trigger_ResourceUse - <nil>
type SoftLayer_Scale_Policy_Trigger_ResourceUse struct {

	// WatchCount - no documentation
	WatchCount uint64 `json:"watchCount"`

	// Watches - no documentation
	Watches []*SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch `json:"watches"`
}

// CreateObject - <nil>
func (softlayer_scale_policy_trigger_resourceuse *SoftLayer_Scale_Policy_Trigger_ResourceUse) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Policy_Trigger_ResourceUse) (*SoftLayer_Scale_Policy_Trigger_ResourceUse, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_ResourceUse
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_trigger_resourceuse *SoftLayer_Scale_Policy_Trigger_ResourceUse) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Policy_Trigger_ResourceUse, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_ResourceUse
	return returnValue, nil
}
