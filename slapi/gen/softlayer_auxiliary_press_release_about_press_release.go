package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_About_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_About_Press_Release struct {

	// AboutParagraphCount - no documentation
	AboutParagraphCount uint64 `json:"aboutParagraphCount"`

	// AboutParagraphs - <nil>
	AboutParagraphs []*SoftLayer_Auxiliary_Press_Release_About `json:"aboutParagraphs"`

	// Id - no documentation
	Id int `json:"id"`

	// PressReleaseAboutId - no documentation
	PressReleaseAboutId int `json:"pressReleaseAboutId"`

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_About_Press_Release object
// whose contact id number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Auxiliary_Press_Release service.
func (softlayer_auxiliary_press_release_about_press_release *SoftLayer_Auxiliary_Press_Release_About_Press_Release) GetObject() (*SoftLayer_Auxiliary_Press_Release_About_Press_Release, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_About_Press_Release
	return returnValue, nil
}
