package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Layout_Profile_Containers - <nil>
type SoftLayer_Layout_Profile_Containers struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// LayoutContainerId - The id of the referenced [[SoftLayer_Layout_Container]]
	LayoutContainerId int `json:"layoutContainerId"`

	// LayoutContainerType - no documentation
	LayoutContainerType *SoftLayer_Layout_Container `json:"layoutContainerType"`

	// LayoutProfile - no documentation
	LayoutProfile *SoftLayer_Layout_Profile `json:"layoutProfile"`

	// LayoutProfileId - The id of the referenced [[SoftLayer_Layout_Profile]]
	LayoutProfileId int `json:"layoutProfileId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_layout_profile_containers *SoftLayer_Layout_Profile_Containers) String() string {
	return "SoftLayer_Layout_Profile_Containers"
}

// CreateObject - <nil>
func (softlayer_layout_profile_containers *SoftLayer_Layout_Profile_Containers) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Layout_Profile_Containers) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - <nil>
func (softlayer_layout_profile_containers *SoftLayer_Layout_Profile_Containers) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Layout_Profile_Containers) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_layout_profile_containers *SoftLayer_Layout_Profile_Containers) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Layout_Profile_Containers, error) {
	var returnValue *SoftLayer_Layout_Profile_Containers
	return returnValue, nil
}
