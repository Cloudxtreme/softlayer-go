package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Scale_Member - <nil>
type SoftLayer_Scale_Member struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ScaleGroupId - The identifier of the group this member belongs to.
	ScaleGroupId int `json:"scaleGroupId"`
}

func (softlayer_scale_member *SoftLayer_Scale_Member) String() string {
	return "SoftLayer_Scale_Member"
}

// SoftLayer_Scale_Member_Extended is SoftLayer_Scale_Member with all maskable types.
type SoftLayer_Scale_Member_Extended struct {
	SoftLayer_Scale_Member

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`
}

func (softlayer_scale_member *SoftLayer_Scale_Member_Extended) String() string {
	return "SoftLayer_Scale_Member"
}
