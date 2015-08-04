package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Authentication_Saml - <nil>
type SoftLayer_Account_Authentication_Saml struct {

	// Certificate - <nil>
	Certificate string `json:"certificate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ServiceProviderSingleLogoutEncoding - <nil>
	ServiceProviderSingleLogoutEncoding string `json:"serviceProviderSingleLogoutEncoding,omitempty"`

	// SingleSignOnUrl - <nil>
	SingleSignOnUrl string `json:"singleSignOnUrl,omitempty"`

	// AccountId - <nil>
	AccountId string `json:"accountId,omitempty"`

	// ServiceProviderPublicKey - <nil>
	ServiceProviderPublicKey string `json:"serviceProviderPublicKey,omitempty"`

	// SingleLogoutUrl - <nil>
	SingleLogoutUrl string `json:"singleLogoutUrl,omitempty"`

	// SingleSignOnEncoding - <nil>
	SingleSignOnEncoding string `json:"singleSignOnEncoding,omitempty"`

	// EntityId - <nil>
	EntityId string `json:"entityId,omitempty"`

	// ServiceProviderEntityId - <nil>
	ServiceProviderEntityId string `json:"serviceProviderEntityId,omitempty"`

	// ServiceProviderSingleSignOnUrl - <nil>
	ServiceProviderSingleSignOnUrl string `json:"serviceProviderSingleSignOnUrl,omitempty"`

	// ServiceProviderSingleSignOnEncoding - <nil>
	ServiceProviderSingleSignOnEncoding string `json:"serviceProviderSingleSignOnEncoding,omitempty"`

	// SingleLogoutEncoding - <nil>
	SingleLogoutEncoding string `json:"singleLogoutEncoding,omitempty"`

	// CertificateFingerprint - <nil>
	CertificateFingerprint string `json:"certificateFingerprint,omitempty"`

	// ServiceProviderCertificate - <nil>
	ServiceProviderCertificate string `json:"serviceProviderCertificate,omitempty"`

	// ServiceProviderSingleLogoutUrl - <nil>
	ServiceProviderSingleLogoutUrl string `json:"serviceProviderSingleLogoutUrl,omitempty"`
}

func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) String() string {
	return "SoftLayer_Account_Authentication_Saml"
}

// SoftLayer_Account_Authentication_Saml_Extended is SoftLayer_Account_Authentication_Saml with all maskable types.
type SoftLayer_Account_Authentication_Saml_Extended struct {
	SoftLayer_Account_Authentication_Saml

	// Account - The account associated with this saml configuration.
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml_Extended) String() string {
	return "SoftLayer_Account_Authentication_Saml"
}
