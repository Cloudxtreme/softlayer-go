package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy_Trigger_Repeating - <nil>
type SoftLayer_Scale_Policy_Trigger_Repeating struct {

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// ScalePolicyId - no documentation
	ScalePolicyId int `json:"scalePolicyId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Schedule - The cron-formatted schedule. This is run in the UTC timezone.
	Schedule string `json:"schedule,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Trigger_Type `json:"type,omitempty"`
}

func (softlayer_scale_policy_trigger_repeating *SoftLayer_Scale_Policy_Trigger_Repeating) String() string {
	return "SoftLayer_Scale_Policy_Trigger_Repeating"
}
