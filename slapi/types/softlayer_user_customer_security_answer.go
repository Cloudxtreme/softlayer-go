package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_user_customer_security_answer *SoftLayer_User_Customer_Security_Answer) String() string {
	return "SoftLayer_User_Customer_Security_Answer"
}
