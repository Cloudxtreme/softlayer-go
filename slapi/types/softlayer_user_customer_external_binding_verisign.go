package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_External_Binding_Verisign - The
// SoftLayer_User_Customer_External_Binding_Verisign data type contains information about a single
// VeriSign external binding. The external binding information is used when a SoftLayer customer logs
// into the SoftLayer customer portal to authenticate them against a 3rd party, in this case VeriSign.
// The information provided by the VeriSign external binding data type includes: * The type of
// credential * The current state of the credential ** Enabled ** Disabled ** Locked * The credential's
// expiration date * The last time the credential was updated SoftLayer users with an active external
// binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Verisign struct {

	// ExternalId - The identifier used to identify this binding to an external authentication source.
	ExternalId string `json:"externalId,omitempty"`

	// VendorId - The [[SoftLayer_User_External_Binding_Vendor|vendor]] identifier of an external
	// authentication binding.
	VendorId int `json:"vendorId,omitempty"`

	// CreateDate - The date that the external authentication binding was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// TypeId - The [[SoftLayer_User_External_Binding_Type|type]] identifier of an external authentication
	// binding.
	TypeId int `json:"typeId,omitempty"`

	// Password - The password used to authenticate the external id at an external authentication source.
	Password string `json:"password,omitempty"`

	// UserId - An external authentication binding's associated [[SoftLayer_User_Customer|user account]]
	// id.
	UserId int `json:"userId,omitempty"`

	// Active - The flag that determines whether the external binding is active will be used for
	// authentication or not.
	Active bool `json:"active,omitempty"`

	// Id - An external authentication binding's internal identifier.
	Id int `json:"id,omitempty"`

	// CredentialExpirationDate - no documentation
	CredentialExpirationDate string `json:"credentialExpirationDate,omitempty"`

	// CredentialLastUpdateDate - no documentation
	CredentialLastUpdateDate string `json:"credentialLastUpdateDate,omitempty"`

	// Vendor - no documentation
	Vendor *SoftLayer_User_External_Binding_Vendor `json:"vendor,omitempty"`

	// AttributeCount - A count of attributes of an external authentication binding.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// Type - no documentation
	Type *SoftLayer_User_External_Binding_Type `json:"type,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_User_External_Binding_Attribute `json:"attributes,omitempty"`

	// User - The SoftLayer user that the external authentication binding belongs to.
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// CredentialState - The current state of a VeriSign credential. This can be 'Enabled', 'Disabled', or
	// 'Locked'.
	CredentialState string `json:"credentialState,omitempty"`

	// CredentialType - The type of VeriSign credential. This can be either 'Hardware' or 'Software'.
	CredentialType string `json:"credentialType,omitempty"`

	// Note - An optional note for identifying the external binding.
	Note string `json:"note,omitempty"`

	// BillingItem - Information regarding the billing item for external authentication.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`
}

func (softlayer_user_customer_external_binding_verisign *SoftLayer_User_Customer_External_Binding_Verisign) String() string {
	return "SoftLayer_User_Customer_External_Binding_Verisign"
}
