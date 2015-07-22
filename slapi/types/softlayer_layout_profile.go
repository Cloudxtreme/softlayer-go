package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// CreateObject - no documentation
func (softlayer_layout_profile *SoftLayer_Layout_Profile) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Layout_Profile) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - This method deletes an existing layout profile and associated custom preferences
func (softlayer_layout_profile *SoftLayer_Layout_Profile) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - This method edits an existing layout profile object by passing in a modified instance
// of the object.
func (softlayer_layout_profile *SoftLayer_Layout_Profile) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Layout_Profile) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_layout_profile *SoftLayer_Layout_Profile) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Layout_Profile, error) {
	var returnValue *SoftLayer_Layout_Profile
	return returnValue, nil
}

// ModifyPreference - This method modifies an existing associated
// [[SoftLayer_Layout_Profile_Preference]] object. If the preference object being modified is a default
// value object, a new record is created to override the default value. Only preferences that are
// assigned to a profile may be updated. Attempts to update a non-existent preference object will
// result in an exception being thrown.
func (softlayer_layout_profile *SoftLayer_Layout_Profile) ModifyPreference(ctx *slapi.RequestContext, templateObject SoftLayer_Layout_Profile_Preference) (*SoftLayer_Layout_Profile_Preference, error) {
	var returnValue *SoftLayer_Layout_Profile_Preference
	return returnValue, nil
}

// ModifyPreferences - Using this method, multiple [[SoftLayer_Layout_Profile_Preference]] objects may
// be updated at once. Refer to [[SoftLayer_Layout_Profile::modifyPreference()]] for more information.
func (softlayer_layout_profile *SoftLayer_Layout_Profile) ModifyPreferences(ctx *slapi.RequestContext, layoutPreferenceObjects []SoftLayer_Layout_Profile_Preference) ([]*SoftLayer_Layout_Profile_Preference, error) {
	var returnValue []*SoftLayer_Layout_Profile_Preference
	return returnValue, nil
}
