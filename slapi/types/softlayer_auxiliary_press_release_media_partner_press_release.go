package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release - <nil>
type SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// MediaPartnerId - no documentation
	MediaPartnerId int `json:"mediaPartnerId,omitempty"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId,omitempty"`
}

func (softlayer_auxiliary_press_release_media_partner_press_release *SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release"
}

// SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release_Extended is SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release with all maskable types.
type SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release_Extended struct {
	SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release

	// MediaPartnerCount - no documentation
	MediaPartnerCount uint64 `json:"mediaPartnerCount,omitempty"`

	// PressReleaseCount - no documentation
	PressReleaseCount uint64 `json:"pressReleaseCount,omitempty"`

	// MediaPartners - <nil>
	MediaPartners []*SoftLayer_Auxiliary_Press_Release_Media_Partner `json:"mediaPartners,omitempty"`

	// PressReleases - <nil>
	PressReleases []*SoftLayer_Auxiliary_Press_Release `json:"pressReleases,omitempty"`
}

func (softlayer_auxiliary_press_release_media_partner_press_release *SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release_Extended) String() string {
	return "SoftLayer_Auxiliary_Press_Release_Media_Partner_Press_Release"
}
