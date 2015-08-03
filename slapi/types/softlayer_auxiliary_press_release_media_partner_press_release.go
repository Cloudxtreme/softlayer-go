package types

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

func (softlayer_auxiliary_press_release_media_partner_press_release *SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release"
}
