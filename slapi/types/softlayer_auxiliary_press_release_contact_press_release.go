package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_Contact_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_Contact_Press_Release struct {

	// PressReleaseContactId - no documentation
	PressReleaseContactId int `json:"pressReleaseContactId,omitempty"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId,omitempty"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount,omitempty"`

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount,omitempty"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact `json:"contacts,omitempty"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases,omitempty"`
}

func (softlayer_auxiliary_press_release_contact_press_release *SoftLayer_Auxiliary_Press_Release_Contact_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Contact_Press_Release"
}
