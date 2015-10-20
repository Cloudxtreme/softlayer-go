package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Auxiliary_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release struct {

	// PublishDate - no documentation
	PublishDate *time.Time `json:"publishDate,omitempty"`

	// ReleaseLocation - no documentation
	ReleaseLocation string `json:"releaseLocation,omitempty"`

	// Title - no documentation
	Title string `json:"title,omitempty"`

	// WebsiteHighlightFlag - Whether or not a press release is highlighted on the SoftLayer Website.
	WebsiteHighlightFlag bool `json:"websiteHighlightFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// SubTitle - no documentation
	SubTitle string `json:"subTitle,omitempty"`

	// PressReleaseContent - <nil>
	PressReleaseContent *SoftLayer_Auxiliary_Press_Release_Content `json:"pressReleaseContent,omitempty"`

	// AboutCount - no documentation
	AboutCount uint64 `json:"aboutCount,omitempty"`

	// About - <nil>
	About []*SoftLayer_Auxiliary_Press_Release_About_Press_Release `json:"about,omitempty"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact_Press_Release `json:"contacts,omitempty"`

	// MediaPartners - <nil>
	MediaPartners []*SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners,omitempty"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount,omitempty"`

	// MediaPartnerCount - no documentation
	MediaPartnerCount uint64 `json:"mediaPartnerCount,omitempty"`
}

func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release"
}
