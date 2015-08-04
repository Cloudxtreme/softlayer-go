package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile - The SoftLayer_Layout_Profile contains the definition of the layout
// profile
type SoftLayer_Layout_Profile struct {

	// ModifyDate - Timestamp of when the layout profile was last updated
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ActiveFlag - no documentation
	ActiveFlag int `json:"activeFlag,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// UserRecordId - The [[SoftLayer_User_Customer]] owning this layout profile
	UserRecordId int `json:"userRecordId,omitempty"`
}

func (softlayer_layout_profile *SoftLayer_Layout_Profile) String() string {
	return "SoftLayer_Layout_Profile"
}

// SoftLayer_Layout_Profile_Extended is SoftLayer_Layout_Profile with all maskable types.
type SoftLayer_Layout_Profile_Extended struct {
	SoftLayer_Layout_Profile

	// LayoutContainerCount - no documentation
	LayoutContainerCount uint64 `json:"layoutContainerCount,omitempty"`

	// LayoutPreferenceCount - no documentation
	LayoutPreferenceCount uint64 `json:"layoutPreferenceCount,omitempty"`

	// LayoutPreferences - <nil>
	LayoutPreferences []*SoftLayer_Layout_Profile_Preference `json:"layoutPreferences,omitempty"`

	// LayoutContainers - <nil>
	LayoutContainers []*SoftLayer_Layout_Container `json:"layoutContainers,omitempty"`
}

func (softlayer_layout_profile *SoftLayer_Layout_Profile_Extended) String() string {
	return "SoftLayer_Layout_Profile"
}
