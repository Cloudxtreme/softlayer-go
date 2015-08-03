package types

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

func (softlayer_auxiliary_press_release_about_press_release *SoftLayer_Auxiliary_Press_Release_About_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_About_Press_Release"
}
