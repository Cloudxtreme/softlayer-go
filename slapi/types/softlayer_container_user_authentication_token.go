package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Authentication_Token - Container class used to hold user authentication
// token
type SoftLayer_Container_User_Authentication_Token struct {

	// Hash - no documentation
	Hash string `json:"hash,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`
}

func (softlayer_container_user_authentication_token *SoftLayer_Container_User_Authentication_Token) String() string {
	return "SoftLayer_Container_User_Authentication_Token"
}
