package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile_Preference - The SoftLayer_Layout_Profile_Preference contains definitions
// for layout preferences
type SoftLayer_Layout_Profile_Preference struct {

	// LayoutPreferenceId - The internal identifier of the overridden [[SoftLayer_Layout_Preference]]
	LayoutPreferenceId int `json:"layoutPreferenceId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// LayoutContainerId - The id of the related [[SoftLayer_Layout_Container]]
	LayoutContainerId int `json:"layoutContainerId,omitempty"`

	// LayoutItemId - no documentation
	LayoutItemId int `json:"layoutItemId,omitempty"`

	// LayoutProfileId - The internal identifier of the related [[SoftLayer_Layout_Profile]]
	LayoutProfileId int `json:"layoutProfileId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`

	// DefaultValueFlag - no documentation
	DefaultValueFlag int `json:"defaultValueFlag,omitempty"`

	// LayoutProfile - <nil>
	LayoutProfile *SoftLayer_Layout_Profile `json:"layoutProfile,omitempty"`

	// LayoutContainer - <nil>
	LayoutContainer *SoftLayer_Layout_Container `json:"layoutContainer,omitempty"`

	// LayoutPreference - <nil>
	LayoutPreference *SoftLayer_Layout_Preference `json:"layoutPreference,omitempty"`

	// LayoutItem - <nil>
	LayoutItem *SoftLayer_Layout_Item `json:"layoutItem,omitempty"`
}

func (softlayer_layout_profile_preference *SoftLayer_Layout_Profile_Preference) String() string {
	return "SoftLayer_Layout_Profile_Preference"
}
