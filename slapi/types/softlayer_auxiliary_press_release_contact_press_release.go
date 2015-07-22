package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

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

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_Contact object whose contact
// id number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Auxiliary_Press_Release service.
func (softlayer_auxiliary_press_release_contact_press_release *SoftLayer_Auxiliary_Press_Release_Contact_Press_Release) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Auxiliary_Press_Release_Contact_Press_Release, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_Contact_Press_Release
	return returnValue, nil
}
