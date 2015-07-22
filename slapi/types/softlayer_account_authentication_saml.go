package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Authentication_Saml - <nil>
type SoftLayer_Account_Authentication_Saml struct {

	// Account - The account associated with this saml configuration.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId string `json:"accountId"`

	// Certificate - <nil>
	Certificate string `json:"certificate"`

	// CertificateFingerprint - <nil>
	CertificateFingerprint string `json:"certificateFingerprint"`

	// EntityId - <nil>
	EntityId string `json:"entityId"`

	// Id - <nil>
	Id int `json:"id"`

	// ServiceProviderCertificate - <nil>
	ServiceProviderCertificate string `json:"serviceProviderCertificate"`

	// ServiceProviderEntityId - <nil>
	ServiceProviderEntityId string `json:"serviceProviderEntityId"`

	// ServiceProviderPublicKey - <nil>
	ServiceProviderPublicKey string `json:"serviceProviderPublicKey"`

	// ServiceProviderSingleLogoutEncoding - <nil>
	ServiceProviderSingleLogoutEncoding string `json:"serviceProviderSingleLogoutEncoding"`

	// ServiceProviderSingleLogoutUrl - <nil>
	ServiceProviderSingleLogoutUrl string `json:"serviceProviderSingleLogoutUrl"`

	// ServiceProviderSingleSignOnEncoding - <nil>
	ServiceProviderSingleSignOnEncoding string `json:"serviceProviderSingleSignOnEncoding"`

	// ServiceProviderSingleSignOnUrl - <nil>
	ServiceProviderSingleSignOnUrl string `json:"serviceProviderSingleSignOnUrl"`

	// SingleLogoutEncoding - <nil>
	SingleLogoutEncoding string `json:"singleLogoutEncoding"`

	// SingleLogoutUrl - <nil>
	SingleLogoutUrl string `json:"singleLogoutUrl"`

	// SingleSignOnEncoding - <nil>
	SingleSignOnEncoding string `json:"singleSignOnEncoding"`

	// SingleSignOnUrl - <nil>
	SingleSignOnUrl string `json:"singleSignOnUrl"`
}

// CreateObject - <nil>
func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Account_Authentication_Saml) (*SoftLayer_Account_Authentication_Saml, error) {
	var returnValue *SoftLayer_Account_Authentication_Saml
	return returnValue, nil
}

// DeleteObject - <nil>
func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit the object by passing in a modified instance of the object
func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Account_Authentication_Saml) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetMetadata - <nil>
func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) GetMetadata(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_authentication_saml *SoftLayer_Account_Authentication_Saml) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Account_Authentication_Saml, error) {
	var returnValue *SoftLayer_Account_Authentication_Saml
	return returnValue, nil
}
