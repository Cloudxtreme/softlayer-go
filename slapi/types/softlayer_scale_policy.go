package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy - <nil>
type SoftLayer_Scale_Policy struct {

	// Cooldown - The number of seconds this policy will wait after lastActionDate on group before
	// performing another action. If not present, the group's cooldown value is used.
	Cooldown int `json:"cooldown,omitempty"`

	// ScaleGroupId - The identifier of the group this member belongs to.
	ScaleGroupId int `json:"scaleGroupId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag,omitempty"`

	// Name - The name of this policy. It must be unique within the group.
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// ResourceUseTriggerCount - A count of the resource-use triggers to check for this group.
	ResourceUseTriggerCount uint64 `json:"resourceUseTriggerCount,omitempty"`

	// ScaleActionCount - A count of the scale actions to perform upon any trigger hit. Currently this must
	// be a single value.
	ScaleActionCount uint64 `json:"scaleActionCount,omitempty"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup,omitempty"`

	// RepeatingTriggerCount - A count of the repeating triggers to check for this group.
	RepeatingTriggerCount uint64 `json:"repeatingTriggerCount,omitempty"`

	// TriggerCount - no documentation
	TriggerCount uint64 `json:"triggerCount,omitempty"`

	// Actions - The actions to perform upon any trigger hit. Currently this must be a single value.
	Actions []*SoftLayer_Scale_Policy_Action `json:"actions,omitempty"`

	// OneTimeTriggers - no documentation
	OneTimeTriggers []*SoftLayer_Scale_Policy_Trigger_OneTime `json:"oneTimeTriggers,omitempty"`

	// ResourceUseTriggers - no documentation
	ResourceUseTriggers []*SoftLayer_Scale_Policy_Trigger_ResourceUse `json:"resourceUseTriggers,omitempty"`

	// ScaleActions - The scale actions to perform upon any trigger hit. Currently this must be a single
	// value.
	ScaleActions []*SoftLayer_Scale_Policy_Action_Scale `json:"scaleActions,omitempty"`

	// OneTimeTriggerCount - A count of the one-time triggers to check for this group.
	OneTimeTriggerCount uint64 `json:"oneTimeTriggerCount,omitempty"`

	// RepeatingTriggers - no documentation
	RepeatingTriggers []*SoftLayer_Scale_Policy_Trigger_Repeating `json:"repeatingTriggers,omitempty"`

	// Triggers - no documentation
	Triggers []*SoftLayer_Scale_Policy_Trigger `json:"triggers,omitempty"`

	// ActionCount - A count of the actions to perform upon any trigger hit. Currently this must be a
	// single value.
	ActionCount uint64 `json:"actionCount,omitempty"`
}

func (softlayer_scale_policy *SoftLayer_Scale_Policy) String() string {
	return "SoftLayer_Scale_Policy"
}
