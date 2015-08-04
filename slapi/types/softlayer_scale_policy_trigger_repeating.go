package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy_Trigger_Repeating - <nil>
type SoftLayer_Scale_Policy_Trigger_Repeating struct {

	// ScalePolicyId - no documentation
	ScalePolicyId int `json:"scalePolicyId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Schedule - The cron-formatted schedule. This is run in the UTC timezone.
	Schedule string `json:"schedule,omitempty"`
}

func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating) String() string {
	return "SoftLayer_Scale_Policy_Trigger_Repeating"
}

// SoftLayer_Scale_Policy_Trigger_Repeating_Extended is SoftLayer_Scale_Policy_Trigger_Repeating with all maskable types.
type SoftLayer_Scale_Policy_Trigger_Repeating_Extended struct {
	SoftLayer_Scale_Policy_Trigger_Repeating

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Trigger_Type `json:"type,omitempty"`

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy,omitempty"`
}

func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating_Extended) String() string {
	return "SoftLayer_Scale_Policy_Trigger_Repeating"
}
