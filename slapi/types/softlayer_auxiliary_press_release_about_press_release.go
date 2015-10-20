package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_About_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_About_Press_Release struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// PressReleaseAboutId - no documentation
	PressReleaseAboutId int `json:"pressReleaseAboutId,omitempty"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId,omitempty"`

	// SortOrder - no documentation
	SortOrder int `json:"sortOrder,omitempty"`

	// AboutParagraphCount - no documentation
	AboutParagraphCount uint64 `json:"aboutParagraphCount,omitempty"`

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount,omitempty"`

	// AboutParagraphs - <nil>
	AboutParagraphs []*SoftLayer_Auxiliary_Press_Release_About `json:"aboutParagraphs,omitempty"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases,omitempty"`
}

func (softlayer_auxiliary_press_release_about_press_release *SoftLayer_Auxiliary_Press_Release_About_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_About_Press_Release"
}
