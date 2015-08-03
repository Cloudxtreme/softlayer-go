package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile - The SoftLayer_Layout_Profile contains the definition of the layout
// profile
type SoftLayer_Layout_Profile struct {

	// ActiveFlag - no documentation
	ActiveFlag int `json:"activeFlag"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// LayoutContainerCount - no documentation
	LayoutContainerCount uint64 `json:"layoutContainerCount"`

	// LayoutContainers - <nil>
	LayoutContainers []*SoftLayer_Layout_Container `json:"layoutContainers"`

	// LayoutPreferenceCount - no documentation
	LayoutPreferenceCount uint64 `json:"layoutPreferenceCount"`

	// LayoutPreferences - <nil>
	LayoutPreferences []*SoftLayer_Layout_Profile_Preference `json:"layoutPreferences"`

	// ModifyDate - Timestamp of when the layout profile was last updated
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// UserRecordId - The [[SoftLayer_User_Customer]] owning this layout profile
	UserRecordId int `json:"userRecordId"`
}

func (softlayer_layout_profile *SoftLayer_Layout_Profile) String() string {
	return "SoftLayer_Layout_Profile"
}
