package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Customer_Security_Answer - The SoftLayer_User_Customer_Security_Answer type contains
// user's answers to security questions.
type SoftLayer_User_Customer_Security_Answer struct {

	// Answer - no documentation
	Answer string `json:"answer"`

	// Id - no documentation
	Id int `json:"id"`

	// Question - The question the security answer is associated with.
	Question *SoftLayer_User_Security_Question `json:"question"`

	// QuestionId - no documentation
	QuestionId int `json:"questionId"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - no documentation
	UserId int `json:"userId"`
}

// GetObject - getObject retrieves the SoftLayer_User_Customer_Security_Answer object whose ID number
// corresponds to the ID number of the init parameter passed to the
// SoftLayer_User_Customer_Security_Answer service.
func (softlayer_user_customer_security_answer *SoftLayer_User_Customer_Security_Answer) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_User_Customer_Security_Answer, error) {
	var returnValue *SoftLayer_User_Customer_Security_Answer
	return returnValue, nil
}
