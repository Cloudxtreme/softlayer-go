package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile_Preference - The SoftLayer_Layout_Profile_Preference contains definitions
// for layout preferences
type SoftLayer_Layout_Profile_Preference struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// LayoutItemId - no documentation
	LayoutItemId int `json:"layoutItemId,omitempty"`

	// LayoutProfileId - The internal identifier of the related [[SoftLayer_Layout_Profile]]
	LayoutProfileId int `json:"layoutProfileId,omitempty"`

	// DefaultValueFlag - no documentation
	DefaultValueFlag int `json:"defaultValueFlag,omitempty"`

	// LayoutContainerId - The id of the related [[SoftLayer_Layout_Container]]
	LayoutContainerId int `json:"layoutContainerId,omitempty"`

	// LayoutPreferenceId - The internal identifier of the overridden [[SoftLayer_Layout_Preference]]
	LayoutPreferenceId int `json:"layoutPreferenceId,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Value - no documentation
	Value string `json:"value,omitempty"`
}

func (softlayer_layout_profile_preference *SoftLayer_Layout_Profile_Preference) String() string {
	return "SoftLayer_Layout_Profile_Preference"
}

// SoftLayer_Layout_Profile_Preference_Extended is SoftLayer_Layout_Profile_Preference with all maskable types.
type SoftLayer_Layout_Profile_Preference_Extended struct {
	SoftLayer_Layout_Profile_Preference

	// LayoutProfile - <nil>
	LayoutProfile *SoftLayer_Layout_Profile `json:"layoutProfile,omitempty"`

	// LayoutContainer - <nil>
	LayoutContainer *SoftLayer_Layout_Container `json:"layoutContainer,omitempty"`

	// LayoutItem - <nil>
	LayoutItem *SoftLayer_Layout_Item `json:"layoutItem,omitempty"`

	// LayoutPreference - <nil>
	LayoutPreference *SoftLayer_Layout_Preference `json:"layoutPreference,omitempty"`
}

func (softlayer_layout_profile_preference *SoftLayer_Layout_Profile_Preference_Extended) String() string {
	return "SoftLayer_Layout_Profile_Preference"
}
