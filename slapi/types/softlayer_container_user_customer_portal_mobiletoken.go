package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Customer_Portal_MobileToken - Container classed used to hold mobile portal
// token
type SoftLayer_Container_User_Customer_Portal_MobileToken struct {

	// UserId - no documentation
	UserId int `json:"userId,omitempty"`

	// HasExternalBinding - True if this user login required an external binding.
	HasExternalBinding bool `json:"hasExternalBinding,omitempty"`

	// Hash - no documentation
	Hash string `json:"hash,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_container_user_customer_portal_mobiletoken *SoftLayer_Container_User_Customer_Portal_MobileToken) String() string {
	return "SoftLayer_Container_User_Customer_Portal_MobileToken"
}
