package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_Contact_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_Contact_Press_Release struct {

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`

	// Id - no documentation
	Id int `json:"id"`

	// PressReleaseContactId - no documentation
	PressReleaseContactId int `json:"pressReleaseContactId"`
}

func (softlayer_auxiliary_press_release_contact_press_release *SoftLayer_Auxiliary_Press_Release_Contact_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Contact_Press_Release"
}

// SoftLayer_Auxiliary_Press_Release_Contact_Press_Release_Extended is SoftLayer_Auxiliary_Press_Release_Contact_Press_Release with all maskable types.
type SoftLayer_Auxiliary_Press_Release_Contact_Press_Release_Extended struct {
	SoftLayer_Auxiliary_Press_Release_Contact_Press_Release

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact `json:"contacts"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`
}

func (softlayer_auxiliary_press_release_contact_press_release *SoftLayer_Auxiliary_Press_Release_Contact_Press_Release_Extended) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Contact_Press_Release"
}
