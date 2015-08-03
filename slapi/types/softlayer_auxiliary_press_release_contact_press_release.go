package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_Contact_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_Contact_Press_Release struct {

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact `json:"contacts"`

	// Id - no documentation
	Id int `json:"id"`

	// PressReleaseContactId - no documentation
	PressReleaseContactId int `json:"pressReleaseContactId"`

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`
}

func (softlayer_auxiliary_press_release_contact_press_release *SoftLayer_Auxiliary_Press_Release_Contact_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Contact_Press_Release"
}
