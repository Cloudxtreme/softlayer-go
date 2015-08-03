package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy_Action - <nil>
type SoftLayer_Scale_Policy_Action struct {

	// TypeId - no documentation
	TypeId int `json:"typeId"`

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
}

// SoftLayer_Scale_Policy_Action_Extended is SoftLayer_Scale_Policy_Action with all maskable types.
type SoftLayer_Scale_Policy_Action_Extended struct {
	SoftLayer_Scale_Policy_Action

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy"`

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Action_Type `json:"type"`
}

func (softlayer_scale_policy_action *SoftLayer_Scale_Policy_Action) String() string {
	return "SoftLayer_Scale_Policy_Action"
}
