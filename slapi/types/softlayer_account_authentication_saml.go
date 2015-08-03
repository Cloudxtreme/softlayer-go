package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Authentication_Saml - <nil>
type SoftLayer_Account_Authentication_Saml struct {

	// SingleSignOnEncoding - <nil>
	SingleSignOnEncoding string `json:"singleSignOnEncoding"`

	// AccountId - <nil>
	AccountId string `json:"accountId"`

	// ServiceProviderCertificate - <nil>
	ServiceProviderCertificate string `json:"serviceProviderCertificate"`

	// Certificate - <nil>
	Certificate string `json:"certificate"`

	// Id - <nil>
	Id int `json:"id"`

	// SingleLogoutUrl - <nil>
	SingleLogoutUrl string `json:"singleLogoutUrl"`

	// ServiceProviderSingleSignOnUrl - <nil>
	ServiceProviderSingleSignOnUrl string `json:"serviceProviderSingleSignOnUrl"`

	// SingleLogoutEncoding - <nil>
	SingleLogoutEncoding string `json:"singleLogoutEncoding"`

	// SingleSignOnUrl - <nil>
	SingleSignOnUrl string `json:"singleSignOnUrl"`

	// EntityId - <nil>
	EntityId string `json:"entityId"`

	// ServiceProviderPublicKey - <nil>
	ServiceProviderPublicKey string `json:"serviceProviderPublicKey"`

	// ServiceProviderSingleSignOnEncoding - <nil>
	ServiceProviderSingleSignOnEncoding string `json:"serviceProviderSingleSignOnEncoding"`

	// ServiceProviderSingleLogoutUrl - <nil>
	ServiceProviderSingleLogoutUrl string `json:"serviceProviderSingleLogoutUrl"`

	// CertificateFingerprint - <nil>
	CertificateFingerprint string `json:"certificateFingerprint"`

	// ServiceProviderEntityId - <nil>
	ServiceProviderEntityId string `json:"serviceProviderEntityId"`

	// ServiceProviderSingleLogoutEncoding - <nil>
	ServiceProviderSingleLogoutEncoding string `json:"serviceProviderSingleLogoutEncoding"`
}

func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) String() string {
	return "SoftLayer_Account_Authentication_Saml"
}

// SoftLayer_Account_Authentication_Saml_Extended is SoftLayer_Account_Authentication_Saml with all maskable types.
type SoftLayer_Account_Authentication_Saml_Extended struct {
	SoftLayer_Account_Authentication_Saml

	// Account - The account associated with this saml configuration.
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml_Extended) String() string {
	return "SoftLayer_Account_Authentication_Saml"
}
