package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Layout_Profile_Preference - The SoftLayer_Layout_Profile_Preference contains definitions
// for layout preferences
type SoftLayer_Layout_Profile_Preference struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DefaultValueFlag - no documentation
	DefaultValueFlag int `json:"defaultValueFlag"`

	// LayoutContainer - <nil>
	LayoutContainer *SoftLayer_Layout_Container `json:"layoutContainer"`

	// LayoutContainerId - The id of the related [[SoftLayer_Layout_Container]]
	LayoutContainerId int `json:"layoutContainerId"`

	// LayoutItem - <nil>
	LayoutItem *SoftLayer_Layout_Item `json:"layoutItem"`

	// LayoutItemId - no documentation
	LayoutItemId int `json:"layoutItemId"`

	// LayoutPreference - <nil>
	LayoutPreference *SoftLayer_Layout_Preference `json:"layoutPreference"`

	// LayoutPreferenceId - The internal identifier of the overridden [[SoftLayer_Layout_Preference]]
	LayoutPreferenceId int `json:"layoutPreferenceId"`

	// LayoutProfile - <nil>
	LayoutProfile *SoftLayer_Layout_Profile `json:"layoutProfile"`

	// LayoutProfileId - The internal identifier of the related [[SoftLayer_Layout_Profile]]
	LayoutProfileId int `json:"layoutProfileId"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Value - no documentation
	Value string `json:"value"`
}

func (softlayer_layout_profile_preference *SoftLayer_Layout_Profile_Preference) String() string {
	return "SoftLayer_Layout_Profile_Preference"
}

// GetObject - <nil>
func (softlayer_layout_profile_preference *SoftLayer_Layout_Profile_Preference) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Layout_Profile_Preference, error) {
	var returnValue *SoftLayer_Layout_Profile_Preference
	return returnValue, nil
}
