package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Auxiliary_Press_Release_Contact - <nil>
type SoftLayer_Auxiliary_Press_Release_Contact struct {

	// Email - no documentation
	Email string `json:"email"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// Id - no documentation
	Id int `json:"id"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// Phone - no documentation
	Phone string `json:"phone"`

	// ProfessionalTitle - no documentation
	ProfessionalTitle string `json:"professionalTitle"`
}

// GetObject - getObject retrieves the SoftLayer_Auxiliary_Press_Release_Contact object whose contact
// id number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Auxiliary_Press_Release service.
func (softlayer_auxiliary_press_release_contact *SoftLayer_Auxiliary_Press_Release_Contact) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Auxiliary_Press_Release_Contact, error) {
	var returnValue *SoftLayer_Auxiliary_Press_Release_Contact
	return returnValue, nil
}
