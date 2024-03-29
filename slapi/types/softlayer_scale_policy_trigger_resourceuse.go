package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy_Trigger_ResourceUse - <nil>
type SoftLayer_Scale_Policy_Trigger_ResourceUse struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// ScalePolicyId - no documentation
	ScalePolicyId int `json:"scalePolicyId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// TypeId - no documentation
	TypeId int `json:"typeId,omitempty"`

	// Watches - no documentation
	Watches []*SoftLayer_Scale_Policy_Trigger_ResourceUse_Watch `json:"watches,omitempty"`

	// WatchCount - no documentation
	WatchCount uint64 `json:"watchCount,omitempty"`

	// ScalePolicy - no documentation
	ScalePolicy *SoftLayer_Scale_Policy `json:"scalePolicy,omitempty"`

	// Type - no documentation
	Type *SoftLayer_Scale_Policy_Trigger_Type `json:"type,omitempty"`
}

func (softlayer_scale_policy_trigger_resourceuse *SoftLayer_Scale_Policy_Trigger_ResourceUse) String() string {
	return "SoftLayer_Scale_Policy_Trigger_ResourceUse"
}
