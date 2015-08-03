package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_About_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_About_Press_Release struct {

	// Id - no documentation
	Id int `json:"id"`

	// PressReleaseAboutId - no documentation
	PressReleaseAboutId int `json:"pressReleaseAboutId"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder"`
}

func (softlayer_auxiliary_press_release_about_press_release *SoftLayer_Auxiliary_Press_Release_About_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_About_Press_Release"
}

// SoftLayer_Auxiliary_Press_Release_About_Press_Release_Extended is SoftLayer_Auxiliary_Press_Release_About_Press_Release with all maskable types.
type SoftLayer_Auxiliary_Press_Release_About_Press_Release_Extended struct {
	SoftLayer_Auxiliary_Press_Release_About_Press_Release

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount"`

	// AboutParagraphs - <nil>
	AboutParagraphs []*SoftLayer_Auxiliary_Press_Release_About `json:"aboutParagraphs"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases"`

	// AboutParagraphCount - no documentation
	AboutParagraphCount uint64 `json:"aboutParagraphCount"`
}

func (softlayer_auxiliary_press_release_about_press_release *SoftLayer_Auxiliary_Press_Release_About_Press_Release_Extended) String() string {
	return "SoftLayer_Auxiliary_Press_Release_About_Press_Release"
}
