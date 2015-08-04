package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// PublishDate - no documentation
	PublishDate *time.Time `json:"publishDate,omitempty"`

	// SubTitle - no documentation
	SubTitle string `json:"subTitle,omitempty"`

	// ReleaseLocation - no documentation
	ReleaseLocation string `json:"releaseLocation,omitempty"`

	// Title - no documentation
	Title string `json:"title,omitempty"`

	// WebsiteHighlightFlag - Whether or not a press release is highlighted on the SoftLayer Website.
	WebsiteHighlightFlag bool `json:"websiteHighlightFlag,omitempty"`
}

func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release"
}

// SoftLayer_Auxiliary_Press_Release_Extended is SoftLayer_Auxiliary_Press_Release with all maskable types.
type SoftLayer_Auxiliary_Press_Release_Extended struct {
	SoftLayer_Auxiliary_Press_Release

	// AboutCount - no documentation
	AboutCount uint64 `json:"aboutCount,omitempty"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount,omitempty"`

	// MediaPartnerCount - no documentation
	MediaPartnerCount uint64 `json:"mediaPartnerCount,omitempty"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact_Press_Release `json:"contacts,omitempty"`

	// PressReleaseContent - <nil>
	PressReleaseContent *SoftLayer_Auxiliary_Press_Release_Content `json:"pressReleaseContent,omitempty"`

	// About - <nil>
	About []*SoftLayer_Auxiliary_Press_Release_About_Press_Release `json:"about,omitempty"`

	// MediaPartners - <nil>
	MediaPartners []*SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners,omitempty"`
}

func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release_Extended) String() string {
	return "SoftLayer_Auxiliary_Press_Release"
}
