package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Policy_Action - <nil>
type SoftLayer_Scale_Policy_Action struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy"`

	// ScalePolicyId - no documentation
	ScalePolicyId int `json:"scalePolicyId"`

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Action_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

// DeleteObject - <nil>
func (softlayer_scale_policy_action *SoftLayer_Scale_Policy_Action) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_scale_policy_action *SoftLayer_Scale_Policy_Action) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Scale_Policy_Action) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_action *SoftLayer_Scale_Policy_Action) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Policy_Action, error) {
	var returnValue *SoftLayer_Scale_Policy_Action
	return returnValue, nil
}
