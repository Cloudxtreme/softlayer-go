package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile_Containers - <nil>
type SoftLayer_Layout_Profile_Containers struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// LayoutContainerId - The id of the referenced [[SoftLayer_Layout_Container]]
	LayoutContainerId int `json:"layoutContainerId,omitempty"`

	// LayoutProfileId - The id of the referenced [[SoftLayer_Layout_Profile]]
	LayoutProfileId int `json:"layoutProfileId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// LayoutProfile - no documentation
	LayoutProfile *SoftLayer_Layout_Profile `json:"layoutProfile,omitempty"`

	// LayoutContainerType - no documentation
	LayoutContainerType *SoftLayer_Layout_Container `json:"layoutContainerType,omitempty"`
}

func (softlayer_layout_profile_containers *SoftLayer_Layout_Profile_Containers) String() string {
	return "SoftLayer_Layout_Profile_Containers"
}
