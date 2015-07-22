package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Policy_Trigger_Repeating - <nil>
type SoftLayer_Scale_Policy_Trigger_Repeating struct {

	// Schedule - The cron-formatted schedule. This is run in the UTC timezone.
	Schedule string `json:"schedule"`
}

// CreateObject - <nil>
func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Scale_Policy_Trigger_Repeating) (*SoftLayer_Scale_Policy_Trigger_Repeating, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_Repeating
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Scale_Policy_Trigger_Repeating, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger_Repeating
	return returnValue, nil
}

// ValidateCronExpression - <nil>
func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating) ValidateCronExpression(commonOptions *slapi.CommonOptions, expression string) error {
	return nil
}
