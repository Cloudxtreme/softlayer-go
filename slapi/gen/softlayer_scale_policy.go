package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Policy - <nil>
type SoftLayer_Scale_Policy struct {

	// ActionCount - A count of the actions to perform upon any trigger hit. Currently this must be a
	// single value.
	ActionCount uint64 `json:"actionCount"`

	// Actions - The actions to perform upon any trigger hit. Currently this must be a single value.
	Actions []*SoftLayer_Scale_Policy_Action `json:"actions"`

	// Cooldown - The number of seconds this policy will wait after lastActionDate on group before
	// performing another action. If not present, the group's cooldown value is used.
	Cooldown int `json:"cooldown"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteFlag - When set and true any edit that happens on this object, be it calling edit on this
	// directly or setting as a child while editing a parent object, will end up being a deletion.
	DeleteFlag bool `json:"deleteFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - The name of this policy. It must be unique within the group.
	Name string `json:"name"`

	// OneTimeTriggerCount - A count of the one-time triggers to check for this group.
	OneTimeTriggerCount uint64 `json:"oneTimeTriggerCount"`

	// OneTimeTriggers - no documentation
	OneTimeTriggers []*SoftLayer_Scale_Policy_Trigger_OneTime `json:"oneTimeTriggers"`

	// RepeatingTriggerCount - A count of the repeating triggers to check for this group.
	RepeatingTriggerCount uint64 `json:"repeatingTriggerCount"`

	// RepeatingTriggers - no documentation
	RepeatingTriggers []*SoftLayer_Scale_Policy_Trigger_Repeating `json:"repeatingTriggers"`

	// ResourceUseTriggerCount - A count of the resource-use triggers to check for this group.
	ResourceUseTriggerCount uint64 `json:"resourceUseTriggerCount"`

	// ResourceUseTriggers - no documentation
	ResourceUseTriggers []*SoftLayer_Scale_Policy_Trigger_ResourceUse `json:"resourceUseTriggers"`

	// ScaleActionCount - A count of the scale actions to perform upon any trigger hit. Currently this must
	// be a single value.
	ScaleActionCount uint64 `json:"scaleActionCount"`

	// ScaleActions - The scale actions to perform upon any trigger hit. Currently this must be a single
	// value.
	ScaleActions []*SoftLayer_Scale_Policy_Action_Scale `json:"scaleActions"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// ScaleGroupId - The identifier of the group this member belongs to.
	ScaleGroupId int `json:"scaleGroupId"`

	// TriggerCount - no documentation
	TriggerCount uint64 `json:"triggerCount"`

	// Triggers - no documentation
	Triggers []*SoftLayer_Scale_Policy_Trigger `json:"triggers"`
}

// CreateObject - <nil>
func (softlayer_scale_policy *SoftLayer_Scale_Policy) CreateObject(templateObject SoftLayer_Scale_Policy) (*SoftLayer_Scale_Policy, error) {
	var returnValue *SoftLayer_Scale_Policy
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_scale_policy *SoftLayer_Scale_Policy) DeleteObject() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_scale_policy *SoftLayer_Scale_Policy) EditObject(templateObject SoftLayer_Scale_Policy) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_policy *SoftLayer_Scale_Policy) GetObject() (*SoftLayer_Scale_Policy, error) {
	var returnValue *SoftLayer_Scale_Policy
	return returnValue, nil
}

// Trigger - <nil>
func (softlayer_scale_policy *SoftLayer_Scale_Policy) Trigger() ([]*SoftLayer_Scale_Member, error) {
	var returnValue []*SoftLayer_Scale_Member
	return returnValue, nil
}
