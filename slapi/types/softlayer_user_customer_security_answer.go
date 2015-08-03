package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Security_Answer - The SoftLayer_User_Customer_Security_Answer type contains
// user's answers to security questions.
type SoftLayer_User_Customer_Security_Answer struct {

	// Answer - no documentation
	Answer string `json:"answer"`

	// Id - no documentation
	Id int `json:"id"`

	// QuestionId - no documentation
	QuestionId int `json:"questionId"`

	// UserId - no documentation
	UserId int `json:"userId"`
}

// SoftLayer_User_Customer_Security_Answer_Extended is SoftLayer_User_Customer_Security_Answer with all maskable types.
type SoftLayer_User_Customer_Security_Answer_Extended struct {
	SoftLayer_User_Customer_Security_Answer

	// Question - The question the security answer is associated with.
	Question *SoftLayer_User_Security_Question `json:"question"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_security_answer *SoftLayer_User_Customer_Security_Answer) String() string {
	return "SoftLayer_User_Customer_Security_Answer"
}
