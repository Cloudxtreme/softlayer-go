package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Press_Release_Content - <nil>
type SoftLayer_Auxiliary_Press_Release_Content struct {

	// Id - no documentation
	Id int `json:"id"`

	// PressReleaseId - no documentation
	PressReleaseId int `json:"pressReleaseId"`

	// Text - no documentation
	Text string `json:"text"`
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_Content object whose ID number
// corresponds to the ID number of the init parameter passed to the SoftLayer_Auxiliary_Press_Release
// service.
func (softlayer_auxiliary_press_release_content *SoftLayer_Auxiliary_Press_Release_Content) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Auxiliary_Press_Release_Content, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_Content
	return returnValue, nil
}
