package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
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

	// ScalePolicyId - no documentation
	ScalePolicyId int `json:"scalePolicyId"`

	// TypeId - no documentation
	TypeId int `json:"typeId"`
}

func (softlayer_scale_policy_trigger *SoftLayer_Scale_Policy_Trigger) String() string {
	return "SoftLayer_Scale_Policy_Trigger"
}

// SoftLayer_Scale_Policy_Trigger_Extended is SoftLayer_Scale_Policy_Trigger with all maskable types.
type SoftLayer_Scale_Policy_Trigger_Extended struct {
	SoftLayer_Scale_Policy_Trigger

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy"`

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Trigger_Type `json:"type"`
}

func (softlayer_scale_policy_trigger *SoftLayer_Scale_Policy_Trigger_Extended) String() string {
	return "SoftLayer_Scale_Policy_Trigger"
}
