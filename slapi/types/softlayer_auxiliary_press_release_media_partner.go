package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Press_Release_Media_Partner - <nil>
type SoftLayer_Auxiliary_Press_Release_Media_Partner struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_Contact object whose contact
// id number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Auxiliary_Press_Release service.
func (softlayer_auxiliary_press_release_media_partner *SoftLayer_Auxiliary_Press_Release_Media_Partner) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Press_Release_Media_Partner, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_Media_Partner
	return returnValue, nil
}
