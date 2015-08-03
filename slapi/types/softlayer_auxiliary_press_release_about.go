package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Auxiliary_Press_Release_About - <nil>
type SoftLayer_Auxiliary_Press_Release_About struct {

	// Content - no documentation
	Content string `json:"content"`

	// Id - no documentation
	Id int `json:"id"`

	// Title - no documentation
	Title string `json:"title"`
}

func (softlayer_auxiliary_press_release_about *SoftLayer_Auxiliary_Press_Release_About) String() string {
	return "SoftLayer_Auxiliary_Press_Release_About"
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_About object whose about id
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Auxiliary_Press_Release service.
func (softlayer_auxiliary_press_release_about *SoftLayer_Auxiliary_Press_Release_About) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Press_Release_About, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_About
	return returnValue, nil
}
