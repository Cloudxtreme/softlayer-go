package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release struct {

	// About - <nil>
	About []*SoftLayer_Auxiliary_Press_Release_About_Press_Release `json:"about"`

	// AboutCount - no documentation
	AboutCount uint64 `json:"aboutCount"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`

	// Contacts - <nil>
	Contacts []*SoftLayer_Auxiliary_Press_Release_Contact_Press_Release `json:"contacts"`

	// Id - no documentation
	Id int `json:"id"`

	// MediaPartnerCount - no documentation
	MediaPartnerCount uint64 `json:"mediaPartnerCount"`

	// MediaPartners - <nil>
	MediaPartners []*SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release `json:"mediaPartners"`

	// PressReleaseContent - <nil>
	PressReleaseContent *SoftLayer_Auxiliary_Press_Release_Content `json:"pressReleaseContent"`

	// PublishDate - no documentation
	PublishDate *time.Time `json:"publishDate"`

	// ReleaseLocation - no documentation
	ReleaseLocation string `json:"releaseLocation"`

	// SubTitle - no documentation
	SubTitle string `json:"subTitle"`

	// Title - no documentation
	Title string `json:"title"`

	// WebsiteHighlightFlag - Whether or not a press release is highlighted on the SoftLayer Website.
	WebsiteHighlightFlag bool `json:"websiteHighlightFlag"`
}

// GetAllObjects - Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which contain all
// press releases.
func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Press_Release, error) {
	var returnValue []*SoftLayer_Auxiliary_Press_Release
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release
// service.
func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Press_Release, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release
	return returnValue, nil
}

// GetRenderedPressRelease - Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which
// contain all press releases.
func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) GetRenderedPressRelease(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Press_Release, error) {
	var returnValue []*SoftLayer_Auxiliary_Press_Release
	return returnValue, nil
}

// GetRenderedPressReleases - Retrieve an array of SoftLayer_Auxiliary_Press_Release data types, which
// contain all press releases for a given year and or result limit.
func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) GetRenderedPressReleases(ctx *slapi.RequestContext, resultLimit string, year string) ([]*SoftLayer_Auxiliary_Press_Release, error) {
	var returnValue []*SoftLayer_Auxiliary_Press_Release
	return returnValue, nil
}

// GetWebsiteHighlightPressReleases - Retrieve an array of SoftLayer_Auxiliary_Press_Release data
// types, which have the website highlight flag set.
func (softlayer_auxiliary_press_release *SoftLayer_Auxiliary_Press_Release) GetWebsiteHighlightPressReleases(ctx *slapi.RequestContext) ([]*SoftLayer_Auxiliary_Press_Release, error) {
	var returnValue []*SoftLayer_Auxiliary_Press_Release
	return returnValue, nil
}
