package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Policy_Trigger_OneTime - <nil>
type SoftLayer_Scale_Policy_Trigger_OneTime struct {

	// Date - no documentation
	Date *time.Time `json:"date"`
}

func (softlayer_scale_policy_trigger_onetime *SoftLayer_Scale_Policy_Trigger_OneTime) String() string {
	return "SoftLayer_Scale_Policy_Trigger_OneTime"
}

// CreateObject - <nil>
func (softlayer_scale_policy_trigger_onetime *SoftLayer_Scale_Policy_Trigger_OneTime) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Policy_Trigger_OneTime) (*SoftLayer_Scale_Policy_Trigger_OneTime, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_OneTime
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_trigger_onetime *SoftLayer_Scale_Policy_Trigger_OneTime) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Policy_Trigger_OneTime, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_OneTime
	return returnValue, nil
}
