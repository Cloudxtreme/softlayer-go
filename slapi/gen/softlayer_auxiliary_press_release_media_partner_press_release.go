package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release struct {

	// Id - no documentation
	Id int `json:"id"`

	// MediaPartnerCount - no documentation
	MediaPartnerCount uint64 `json:"mediaPartnerCount"`

	// MediaPartnerId - no documentation
	MediaPartnerId int `json:"mediaPartnerId"`

	// MediaPartners - <nil>
	MediaPartners []*SoftLayer_Auxiliary_Press_Release_Media_Partner `json:"mediaPartners"`

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases"`
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release
// object whose media partner id number corresponds to the ID number of the init parameter passed to
// the SoftLayer_Auxiliary_Press_Release service.
func (softlayer_auxiliary_press_release_media_partner_press_release *SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release) GetObject() (*SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release
	return returnValue, nil
}
