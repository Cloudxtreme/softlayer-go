package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Customer_External_Binding_Totp - no documentation
type SoftLayer_Container_User_Customer_External_Binding_Totp struct {

	// Password - Your SoftLayer customer portal user's portal password.
	Password string `json:"password,omitempty"`

	// SecurityQuestionAnswer - no documentation
	SecurityQuestionAnswer string `json:"securityQuestionAnswer,omitempty"`

	// SecurityQuestionId - A security question you wish to answer when authenticating to the SoftLayer
	// customer portal. This parameter isn't required if no security questions are set on your portal
	// account or if your account is configured to not require answering a security account upon login.
	SecurityQuestionId int `json:"securityQuestionId,omitempty"`

	// Username - The username you wish to authenticate to the SoftLayer customer portal with.
	Username string `json:"username,omitempty"`

	// Vendor - The name of the vendor that will be used for external authentication
	Vendor string `json:"vendor,omitempty"`

	// SecurityCode - The security code used to validate a Totp credential.
	SecurityCode string `json:"securityCode,omitempty"`

	// AuthenticationToken - The unique token that is created by an external authentication request.
	AuthenticationToken string `json:"authenticationToken,omitempty"`
}

func (softlayer_container_user_customer_external_binding_totp *SoftLayer_Container_User_Customer_External_Binding_Totp) String() string {
	return "SoftLayer_Container_User_Customer_External_Binding_Totp"
}
