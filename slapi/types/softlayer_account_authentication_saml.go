package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Authentication_Saml - <nil>
type SoftLayer_Account_Authentication_Saml struct {

	// CertificateFingerprint - <nil>
	CertificateFingerprint string `json:"certificateFingerprint"`

	// ServiceProviderPublicKey - <nil>
	ServiceProviderPublicKey string `json:"serviceProviderPublicKey"`

	// ServiceProviderSingleSignOnUrl - <nil>
	ServiceProviderSingleSignOnUrl string `json:"serviceProviderSingleSignOnUrl"`

	// ServiceProviderCertificate - <nil>
	ServiceProviderCertificate string `json:"serviceProviderCertificate"`

	// ServiceProviderSingleLogoutUrl - <nil>
	ServiceProviderSingleLogoutUrl string `json:"serviceProviderSingleLogoutUrl"`

	// SingleSignOnUrl - <nil>
	SingleSignOnUrl string `json:"singleSignOnUrl"`

	// EntityId - <nil>
	EntityId string `json:"entityId"`

	// Id - <nil>
	Id int `json:"id"`

	// ServiceProviderEntityId - <nil>
	ServiceProviderEntityId string `json:"serviceProviderEntityId"`

	// ServiceProviderSingleLogoutEncoding - <nil>
	ServiceProviderSingleLogoutEncoding string `json:"serviceProviderSingleLogoutEncoding"`

	// ServiceProviderSingleSignOnEncoding - <nil>
	ServiceProviderSingleSignOnEncoding string `json:"serviceProviderSingleSignOnEncoding"`

	// AccountId - <nil>
	AccountId string `json:"accountId"`

	// Certificate - <nil>
	Certificate string `json:"certificate"`

	// SingleLogoutEncoding - <nil>
	SingleLogoutEncoding string `json:"singleLogoutEncoding"`

	// SingleLogoutUrl - <nil>
	SingleLogoutUrl string `json:"singleLogoutUrl"`

	// SingleSignOnEncoding - <nil>
	SingleSignOnEncoding string `json:"singleSignOnEncoding"`
}

// SoftLayer_Account_Authentication_Saml_Extended is SoftLayer_Account_Authentication_Saml with all maskable types.
type SoftLayer_Account_Authentication_Saml_Extended struct {
	SoftLayer_Account_Authentication_Saml

	// Account - The account associated with this saml configuration.
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) String() string {
	return "SoftLayer_Account_Authentication_Saml"
}
