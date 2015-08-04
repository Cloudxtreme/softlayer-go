package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Layout_Profile_Customer - <nil>
type SoftLayer_Layout_Profile_Customer struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// UserRecordId - The [[SoftLayer_User_Customer]] owning this layout profile
	UserRecordId int `json:"userRecordId,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ActiveFlag - no documentation
	ActiveFlag int `json:"activeFlag,omitempty"`

	// ModifyDate - Timestamp of when the layout profile was last updated
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

func (softlayer_layout_profile_customer *SoftLayer_Layout_Profile_Customer) String() string {
	return "SoftLayer_Layout_Profile_Customer"
}

// SoftLayer_Layout_Profile_Customer_Extended is SoftLayer_Layout_Profile_Customer with all maskable types.
type SoftLayer_Layout_Profile_Customer_Extended struct {
	SoftLayer_Layout_Profile_Customer

	// LayoutContainerCount - no documentation
	LayoutContainerCount uint64 `json:"layoutContainerCount,omitempty"`

	// LayoutPreferenceCount - no documentation
	LayoutPreferenceCount uint64 `json:"layoutPreferenceCount,omitempty"`

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Customer `json:"userRecord,omitempty"`

	// LayoutContainers - <nil>
	LayoutContainers []*SoftLayer_Layout_Container `json:"layoutContainers,omitempty"`

	// LayoutPreferences - <nil>
	LayoutPreferences []*SoftLayer_Layout_Profile_Preference `json:"layoutPreferences,omitempty"`
}

func (softlayer_layout_profile_customer *SoftLayer_Layout_Profile_Customer_Extended) String() string {
	return "SoftLayer_Layout_Profile_Customer"
}
