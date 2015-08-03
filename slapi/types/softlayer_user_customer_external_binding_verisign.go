package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_External_Binding_Verisign - The
// SoftLayer_User_Customer_External_Binding_Verisign data type contains information about a single
// VeriSign external binding. The external binding information is used when a SoftLayer customer logs
// into the SoftLayer customer portal to authenticate them against a 3rd party, in this case VeriSign.
// The information provided by the VeriSign external binding data type includes: * The type of
// credential * The current state of the credential ** Enabled ** Disabled ** Locked * The credential's
// expiration date * The last time the credential was updated SoftLayer users with an active external
// binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Verisign struct {
}

// SoftLayer_User_Customer_External_Binding_Verisign_Extended is SoftLayer_User_Customer_External_Binding_Verisign with all maskable types.
type SoftLayer_User_Customer_External_Binding_Verisign_Extended struct {
	SoftLayer_User_Customer_External_Binding_Verisign

	// CredentialExpirationDate - no documentation
	CredentialExpirationDate string `json:"credentialExpirationDate"`

	// CredentialLastUpdateDate - no documentation
	CredentialLastUpdateDate string `json:"credentialLastUpdateDate"`

	// CredentialState - The current state of a VeriSign credential. This can be 'Enabled', 'Disabled', or
	// 'Locked'.
	CredentialState string `json:"credentialState"`

	// CredentialType - The type of VeriSign credential. This can be either 'Hardware' or 'Software'.
	CredentialType string `json:"credentialType"`
}

func (softlayer_user_customer_external_binding_verisign *SoftLayer_User_Customer_External_Binding_Verisign) String() string {
	return "SoftLayer_User_Customer_External_Binding_Verisign"
}
