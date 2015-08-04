package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile_Customer - <nil>
type SoftLayer_Layout_Profile_Customer struct {

	// ActiveFlag - no documentation
	ActiveFlag int `json:"activeFlag,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// UserRecordId - The [[SoftLayer_User_Customer]] owning this layout profile
	UserRecordId int `json:"userRecordId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ModifyDate - Timestamp of when the layout profile was last updated
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// LayoutPreferenceCount - no documentation
	LayoutPreferenceCount uint64 `json:"layoutPreferenceCount,omitempty"`

	// LayoutContainers - <nil>
	LayoutContainers []*SoftLayer_Layout_Container `json:"layoutContainers,omitempty"`

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Customer `json:"userRecord,omitempty"`

	// LayoutContainerCount - no documentation
	LayoutContainerCount uint64 `json:"layoutContainerCount,omitempty"`

	// LayoutPreferences - <nil>
	LayoutPreferences []*SoftLayer_Layout_Profile_Preference `json:"layoutPreferences,omitempty"`
}

func (softlayer_layout_profile_customer *SoftLayer_Layout_Profile_Customer) String() string {
	return "SoftLayer_Layout_Profile_Customer"
}
