package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Authentication_Saml - <nil>
type SoftLayer_Account_Authentication_Saml struct {

	// AccountId - <nil>
	AccountId string `json:"accountId,omitempty"`

	// Certificate - <nil>
	Certificate string `json:"certificate,omitempty"`

	// EntityId - <nil>
	EntityId string `json:"entityId,omitempty"`

	// SingleLogoutEncoding - <nil>
	SingleLogoutEncoding string `json:"singleLogoutEncoding,omitempty"`

	// SingleSignOnEncoding - <nil>
	SingleSignOnEncoding string `json:"singleSignOnEncoding,omitempty"`

	// ServiceProviderCertificate - <nil>
	ServiceProviderCertificate string `json:"serviceProviderCertificate,omitempty"`

	// ServiceProviderPublicKey - <nil>
	ServiceProviderPublicKey string `json:"serviceProviderPublicKey,omitempty"`

	// ServiceProviderSingleLogoutUrl - <nil>
	ServiceProviderSingleLogoutUrl string `json:"serviceProviderSingleLogoutUrl,omitempty"`

	// ServiceProviderSingleSignOnEncoding - <nil>
	ServiceProviderSingleSignOnEncoding string `json:"serviceProviderSingleSignOnEncoding,omitempty"`

	// ServiceProviderSingleSignOnUrl - <nil>
	ServiceProviderSingleSignOnUrl string `json:"serviceProviderSingleSignOnUrl,omitempty"`

	// SingleLogoutUrl - <nil>
	SingleLogoutUrl string `json:"singleLogoutUrl,omitempty"`

	// SingleSignOnUrl - <nil>
	SingleSignOnUrl string `json:"singleSignOnUrl,omitempty"`

	// CertificateFingerprint - <nil>
	CertificateFingerprint string `json:"certificateFingerprint,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ServiceProviderEntityId - <nil>
	ServiceProviderEntityId string `json:"serviceProviderEntityId,omitempty"`

	// ServiceProviderSingleLogoutEncoding - <nil>
	ServiceProviderSingleLogoutEncoding string `json:"serviceProviderSingleLogoutEncoding,omitempty"`

	// Account - The account associated with this saml configuration.
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) String() string {
	return "SoftLayer_Account_Authentication_Saml"
}
