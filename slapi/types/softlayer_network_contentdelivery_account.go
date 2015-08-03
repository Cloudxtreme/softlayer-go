package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_ContentDelivery_Account - The SoftLayer_Network_ContentDelivery_Account data type
// models an individual CDN account. CDN accounts contain references to the SoftLayer customer account
// they belong to, login credentials for upload services, and a CDN account's status. Please contact
// SoftLayer sales to purchase or cancel a CDN account
type SoftLayer_Network_ContentDelivery_Account struct {

	// AccountId - The internal identifier of the customer account that a CDN account belongs to.
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_Network_ContentDelivery_Account_Extended is SoftLayer_Network_ContentDelivery_Account with all maskable types.
type SoftLayer_Network_ContentDelivery_Account_Extended struct {
	SoftLayer_Network_ContentDelivery_Account

	// AuthenticationIpAddressCount - A count of the IP addresses that are used for the content
	// authentication service.
	AuthenticationIpAddressCount uint64 `json:"authenticationIpAddressCount"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CdnAccountName - no documentation
	CdnAccountName string `json:"cdnAccountName"`

	// DependantServiceFlag - Indicates if CDN account is dependent on other service. If set, this CDN
	// account is limited to these services: createOriginPullMapping, deleteOriginPullRule,
	// getOriginPullMappingInformation, getCdnUrls, purgeCache, loadContent, manageHttpCompression
	DependantServiceFlag bool `json:"dependantServiceFlag"`

	// LegacyCdnFlag - no documentation
	LegacyCdnFlag bool `json:"legacyCdnFlag"`

	// AuthenticationIpAddresses - The IP addresses that are used for the content authentication service.
	AuthenticationIpAddresses []*SoftLayer_Network_ContentDelivery_Authentication_Address `json:"authenticationIpAddresses"`

	// CdnSolutionName - no documentation
	CdnSolutionName string `json:"cdnSolutionName"`

	// ProviderPortalAccessFlag - Indicates if customer is allowed to access the CDN provider's management
	// portal.
	ProviderPortalAccessFlag bool `json:"providerPortalAccessFlag"`

	// Status - A CDN account's status presented in a more detailed data type.
	Status *SoftLayer_Network_ContentDelivery_Account_Status `json:"status"`

	// Account - The customer account that a CDN account belongs to.
	Account *SoftLayer_Account `json:"account"`

	// AssociatedCdnAccountId - The CDN account id that this CDN account is associated with.
	AssociatedCdnAccountId string `json:"associatedCdnAccountId"`

	// CdnAccountNote - no documentation
	CdnAccountNote string `json:"cdnAccountNote"`

	// LogEnabledFlag - no documentation
	LogEnabledFlag string `json:"logEnabledFlag"`

	// TokenAuthenticationEnabledFlag - Indicates if the token authentication service is enabled or not.
	TokenAuthenticationEnabledFlag bool `json:"tokenAuthenticationEnabledFlag"`
}

func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) String() string {
	return "SoftLayer_Network_ContentDelivery_Account"
}
