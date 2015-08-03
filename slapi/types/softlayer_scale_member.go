package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Scale_Member - <nil>
type SoftLayer_Scale_Member struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ScaleGroup - no documentation
	ScaleGroup *SoftLayer_Scale_Group `json:"scaleGroup"`

	// ScaleGroupId - The identifier of the group this member belongs to.
	ScaleGroupId int `json:"scaleGroupId"`
}

func (softlayer_scale_member *SoftLayer_Scale_Member) String() string {
	return "SoftLayer_Scale_Member"
}

// DeleteObject - <nil>
func (softlayer_scale_member *SoftLayer_Scale_Member) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_scale_member *SoftLayer_Scale_Member) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Scale_Member, error) {
	var returnValue *SoftLayer_Scale_Member
	return returnValue, nil
}
