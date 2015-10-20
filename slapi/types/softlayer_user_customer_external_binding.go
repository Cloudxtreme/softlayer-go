package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_External_Binding - The SoftLayer_User_Customer_External_Binding data type
// contains general information for a single external binding. This includes the 3rd party vendor, type
// of binding, and a unique identifier and password that is used to authenticate against the 3rd party
// service.
type SoftLayer_User_Customer_External_Binding struct {

	// CreateDate - The date that the external authentication binding was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// TypeId - The [[SoftLayer_User_External_Binding_Type|type]] identifier of an external authentication
	// binding.
	TypeId int `json:"typeId,omitempty"`

	// VendorId - The [[SoftLayer_User_External_Binding_Vendor|vendor]] identifier of an external
	// authentication binding.
	VendorId int `json:"vendorId,omitempty"`

	// UserId - An external authentication binding's associated [[SoftLayer_User_Customer|user account]]
	// id.
	UserId int `json:"userId,omitempty"`

	// Active - The flag that determines whether the external binding is active will be used for
	// authentication or not.
	Active bool `json:"active,omitempty"`

	// ExternalId - The identifier used to identify this binding to an external authentication source.
	ExternalId string `json:"externalId,omitempty"`

	// Password - The password used to authenticate the external id at an external authentication source.
	Password string `json:"password,omitempty"`

	// Id - An external authentication binding's internal identifier.
	Id int `json:"id,omitempty"`

	// Type - no documentation
	Type *SoftLayer_User_External_Binding_Type `json:"type,omitempty"`

	// Vendor - no documentation
	Vendor *SoftLayer_User_External_Binding_Vendor `json:"vendor,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_User_External_Binding_Attribute `json:"attributes,omitempty"`

	// Note - An optional note for identifying the external binding.
	Note string `json:"note,omitempty"`

	// User - The SoftLayer user that the external authentication binding belongs to.
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// BillingItem - Information regarding the billing item for external authentication.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// AttributeCount - A count of attributes of an external authentication binding.
	AttributeCount uint64 `json:"attributeCount,omitempty"`
}

func (softlayer_user_customer_external_binding *SoftLayer_User_Customer_External_Binding) String() string {
	return "SoftLayer_User_Customer_External_Binding"
}
