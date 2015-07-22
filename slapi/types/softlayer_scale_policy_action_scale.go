package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Policy_Action_Scale - <nil>
type SoftLayer_Scale_Policy_Action_Scale struct {

	// Amount - The number to scale by. This number has different meanings based on type.
	Amount int `json:"amount"`

	// ScaleType - The type of scale to perform. Possible values: * - Force the group to be set at a
	// specific number of group members. This may include scaling up or down or not at all. If the amount
	// is outside of the min/max range of the group, an error occurs. * - Scale the group up or down based
	// on the positive or negative percentage given in amount. The number is a percent of the current group
	// member count. Any extra percent after the decimal point is always ignored. If the resulting amount
	// is zero, -1 or 1 is used depending upon whether the percentage was negative or positive
	// respectively. * - Scale the group up or down by the positive or negative value given in amount.
	ScaleType string `json:"scaleType"`
}

// CreateObject - <nil>
func (softlayer_scale_policy_action_scale *SoftLayer_Scale_Policy_Action_Scale) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Scale_Policy_Action_Scale) (*SoftLayer_Scale_Policy_Action_Scale, error) {
	var returnValue *SoftLayer_Scale_Policy_Action_Scale
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_action_scale *SoftLayer_Scale_Policy_Action_Scale) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Scale_Policy_Action_Scale, error) {
	var returnValue *SoftLayer_Scale_Policy_Action_Scale
	return returnValue, nil
}
