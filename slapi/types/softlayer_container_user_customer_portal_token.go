package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Customer_Portal_Token - no documentation
type SoftLayer_Container_User_Customer_Portal_Token struct {

	// Hash - no documentation
	Hash string `json:"hash"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - no documentation
	UserId int `json:"userId"`
}

func (softlayer_container_user_customer_portal_token *SoftLayer_Container_User_Customer_Portal_Token) String() string {
	return "SoftLayer_Container_User_Customer_Portal_Token"
}
