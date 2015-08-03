package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

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

func (softlayer_user_security_question *SoftLayer_User_Security_Question) String() string {
	return "SoftLayer_User_Security_Question"
}
