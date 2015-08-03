package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Group_Log - <nil>
type SoftLayer_Scale_Group_Log struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Description - A textual description of what happened during this action.
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId"`
}

func (softlayer_scale_group_log *SoftLayer_Scale_Group_Log) String() string {
	return "SoftLayer_Scale_Group_Log"
}
