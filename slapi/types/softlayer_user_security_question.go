package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Security_Question - The SoftLayer_User_Security_Question data type contains
// questions.
type SoftLayer_User_Security_Question struct {

	// DisplayOrder - no documentation
	DisplayOrder int `json:"displayOrder"`

	// Id - no documentation
	Id int `json:"id"`

	// Question - no documentation
	Question string `json:"question"`

	// Viewable - no documentation
	Viewable int `json:"viewable"`
}

// GetAllObjects - no documentation
func (softlayer_user_security_question *SoftLayer_User_Security_Question) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_User_Security_Question, error) {
	var returnValue []*SoftLayer_User_Security_Question
	return returnValue, nil
}

// GetObject - getAllObjects retrieves all the SoftLayer_User_Security_Question objects where it is set
// to be viewable.
func (softlayer_user_security_question *SoftLayer_User_Security_Question) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Security_Question, error) {
	var returnValue *SoftLayer_User_Security_Question
	return returnValue, nil
}
