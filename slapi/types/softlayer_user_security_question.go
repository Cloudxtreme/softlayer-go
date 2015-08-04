package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Security_Question - The SoftLayer_User_Security_Question data type contains
// questions.
type SoftLayer_User_Security_Question struct {

	// Question - no documentation
	Question string `json:"question,omitempty"`

	// Viewable - no documentation
	Viewable int `json:"viewable,omitempty"`

	// DisplayOrder - no documentation
	DisplayOrder int `json:"displayOrder,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`
}

func (softlayer_user_security_question *SoftLayer_User_Security_Question) String() string {
	return "SoftLayer_User_Security_Question"
}
