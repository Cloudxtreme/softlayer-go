package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy - <nil>
type SoftLayer_Scale_Policy struct {

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// ScaleGroupId - The identifier of the group this member belongs to.
	ScaleGroupId int `json:"scaleGroupId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - The name of this policy. It must be unique within the group.
	Name string `json:"name"`

	// Cooldown - The number of seconds this policy will wait after lastActionDate on group before
	// performing another action. If not present, the group's cooldown value is used.
	Cooldown int `json:"cooldown"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`
}

func (softlayer_scale_policy *SoftLayer_Scale_Policy) String() string {
	return "SoftLayer_Scale_Policy"
}

// SoftLayer_Scale_Policy_Extended is SoftLayer_Scale_Policy with all maskable types.
type SoftLayer_Scale_Policy_Extended struct {
	SoftLayer_Scale_Policy

	// OneTimeTriggers - no documentation
	OneTimeTriggers []*SoftLayer_Scale_Policy_Trigger_OneTime `json:"oneTimeTriggers"`

	// RepeatingTriggers - no documentation
	RepeatingTriggers []*SoftLayer_Scale_Policy_Trigger_Repeating `json:"repeatingTriggers"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// Triggers - no documentation
	Triggers []*SoftLayer_Scale_Policy_Trigger `json:"triggers"`

	// RepeatingTriggerCount - A count of the repeating triggers to check for this group.
	RepeatingTriggerCount uint64 `json:"repeatingTriggerCount"`

	// OneTimeTriggerCount - A count of the one-time triggers to check for this group.
	OneTimeTriggerCount uint64 `json:"oneTimeTriggerCount"`

	// ResourceUseTriggerCount - A count of the resource-use triggers to check for this group.
	ResourceUseTriggerCount uint64 `json:"resourceUseTriggerCount"`

	// ScaleActionCount - A count of the scale actions to perform upon any trigger hit. Currently this must
	// be a single value.
	ScaleActionCount uint64 `json:"scaleActionCount"`

	// Actions - The actions to perform upon any trigger hit. Currently this must be a single value.
	Actions []*SoftLayer_Scale_Policy_Action `json:"actions"`

	// ResourceUseTriggers - no documentation
	ResourceUseTriggers []*SoftLayer_Scale_Policy_Trigger_ResourceUse `json:"resourceUseTriggers"`

	// ScaleActions - The scale actions to perform upon any trigger hit. Currently this must be a single
	// value.
	ScaleActions []*SoftLayer_Scale_Policy_Action_Scale `json:"scaleActions"`

	// TriggerCount - no documentation
	TriggerCount uint64 `json:"triggerCount"`

	// ActionCount - A count of the actions to perform upon any trigger hit. Currently this must be a
	// single value.
	ActionCount uint64 `json:"actionCount"`
}

func (softlayer_scale_policy *SoftLayer_Scale_Policy_Extended) String() string {
	return "SoftLayer_Scale_Policy"
}
