package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release struct {

	// ReleaseLocation - no documentation
	ReleaseLocation string `json:"releaseLocation"`

	// Title - no documentation
	Title string `json:"title"`

	// PublishDate - no documentation
	PublishDate *time.Time `json:"publishDate"`

	// SubTitle - no documentation
	SubTitle string `json:"subTitle"`

	// WebsiteHighlightFlag - Whether or not a press release is highlighted on the SoftLayer Website.
	WebsiteHighlightFlag bool `json:"websiteHighlightFlag"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release"
}

// SoftLayer_Auxiliary_Press_Release_Extended is SoftLayer_Auxiliary_Press_Release with all maskable types.
type SoftLayer_Auxiliary_Press_Release_Extended struct {
	SoftLayer_Auxiliary_Press_Release

	// AboutCount - no documentation
	AboutCount uint64 `json:"aboutCount"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact_Press_Release `json:"contacts"`

	// MediaPartners - <nil>
	MediaPartners []*SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`

	// MediaPartnerCount - no documentation
	MediaPartnerCount uint64 `json:"mediaPartnerCount"`

	// About - <nil>
	About []*SoftLayer_Auxiliary_Press_Release_About_Press_Release `json:"about"`

	// PressReleaseContent - <nil>
	PressReleaseContent *SoftLayer_Auxiliary_Press_Release_Content `json:"pressReleaseContent"`
}

func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release_Extended) String() string {
	return "SoftLayer_Auxiliary_Press_Release"
}
