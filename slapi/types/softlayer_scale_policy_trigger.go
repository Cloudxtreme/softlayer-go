package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Scale_Policy_Trigger - <nil>
type SoftLayer_Scale_Policy_Trigger struct {

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
	Type *SoftLayer_Scale_Policy_Trigger_Type `json:"type"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

// CreateObject - <nil>
func (softlayer_scale_policy_trigger *SoftLayer_Scale_Policy_Trigger) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Scale_Policy_Trigger) (*SoftLayer_Scale_Policy_Trigger, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_scale_policy_trigger *SoftLayer_Scale_Policy_Trigger) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_scale_policy_trigger *SoftLayer_Scale_Policy_Trigger) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Scale_Policy_Trigger) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy_trigger *SoftLayer_Scale_Policy_Trigger) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Scale_Policy_Trigger, error) {
	var returnValue *SoftLayer_Scale_Policy_Trigger
	return returnValue, nil
}
