package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy_Action_Scale - <nil>
type SoftLayer_Scale_Policy_Action_Scale struct {

	// Amount - The number to scale by. This number has different meanings based on type.
	Amount int `json:"amount,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ScaleType - The type of scale to perform. Possible values: * - Force the group to be set at a
	// specific number of group members. This may include scaling up or down or not at all. If the amount
	// is outside of the min/max range of the group, an error occurs. * - Scale the group up or down based
	// on the positive or negative percentage given in amount. The number is a percent of the current group
	// member count. Any extra percent after the decimal point is always ignored. If the resulting amount
	// is zero, -1 or 1 is used depending upon whether the percentage was negative or positive
	// respectively. * - Scale the group up or down by the positive or negative value given in amount.
	ScaleType string `json:"scaleType,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ScalePolicyId - no documentation
	ScalePolicyId int `json:"scalePolicyId,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`
}

func (softlayer_scale_policy_action_scale *SoftLayer_Scale_Policy_Action_Scale) String() string {
	return "SoftLayer_Scale_Policy_Action_Scale"
}

// SoftLayer_Scale_Policy_Action_Scale_Extended is SoftLayer_Scale_Policy_Action_Scale with all maskable types.
type SoftLayer_Scale_Policy_Action_Scale_Extended struct {
	SoftLayer_Scale_Policy_Action_Scale

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Action_Type `json:"type,omitempty"`
}

func (softlayer_scale_policy_action_scale *SoftLayer_Scale_Policy_Action_Scale_Extended) String() string {
	return "SoftLayer_Scale_Policy_Action_Scale"
}
