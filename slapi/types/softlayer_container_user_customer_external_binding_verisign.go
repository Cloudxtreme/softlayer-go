package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Customer_External_Binding_Verisign - no documentation
type SoftLayer_Container_User_Customer_External_Binding_Verisign struct {

	// SecondSecurityCode - A second security code that is only required if your credential has become
	// unsynchronized.
	SecondSecurityCode string `json:"secondSecurityCode"`

	// SecurityCode - The security code used to validate a VeriSign credential.
	SecurityCode string `json:"securityCode"`
}

func (softlayer_container_user_customer_external_binding_verisign *SoftLayer_Container_User_Customer_External_Binding_Verisign) String() string {
	return "SoftLayer_Container_User_Customer_External_Binding_Verisign"
}
