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

	// ScaleGroupId - no documentation
	ScaleGroupId int `json:"scaleGroupId"`
}

func (softlayer_scale_group_log *SoftLayer_Scale_Group_Log) String() string {
	return "SoftLayer_Scale_Group_Log"
}

// SoftLayer_Scale_Group_Log_Extended is SoftLayer_Scale_Group_Log with all maskable types.
type SoftLayer_Scale_Group_Log_Extended struct {
	SoftLayer_Scale_Group_Log

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`
}

func (softlayer_scale_group_log *SoftLayer_Scale_Group_Log_Extended) String() string {
	return "SoftLayer_Scale_Group_Log"
}
