package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_External_Binding - The SoftLayer_User_External_Binding data type contains general
// information for a single external binding. This includes the 3rd party vendor, type of binding, and
// a unique identifier and password that is used to authenticate against the 3rd party service.
type SoftLayer_User_External_Binding struct {

	// VendorId - The [[SoftLayer_User_External_Binding_Vendor|vendor]] identifier of an external
	// authentication binding.
	VendorId int `json:"vendorId"`

	// Active - The flag that determines whether the external binding is active will be used for
	// authentication or not.
	Active bool `json:"active"`

	// Password - The password used to authenticate the external id at an external authentication source.
	Password string `json:"password"`

	// TypeId - The [[SoftLayer_User_External_Binding_Type|type]] identifier of an external authentication
	// binding.
	TypeId int `json:"typeId"`

	// Id - An external authentication binding's internal identifier.
	Id int `json:"id"`

	// UserId - An external authentication binding's associated [[SoftLayer_User_Customer|user account]]
	// id.
	UserId int `json:"userId"`

	// CreateDate - The date that the external authentication binding was created.
	CreateDate *time.Time `json:"createDate"`

	// ExternalId - The identifier used to identify this binding to an external authentication source.
	ExternalId string `json:"externalId"`
}

// SoftLayer_User_External_Binding_Extended is SoftLayer_User_External_Binding with all maskable types.
type SoftLayer_User_External_Binding_Extended struct {
	SoftLayer_User_External_Binding

	// Vendor - no documentation
	Vendor *SoftLayer_User_External_Binding_Vendor `json:"vendor"`

	// AttributeCount - A count of attributes of an external authentication binding.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - no documentation
	Attributes []*SoftLayer_User_External_Binding_Attribute `json:"attributes"`

	// BillingItem - Information regarding the billing item for external authentication.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// Note - An optional note for identifying the external binding.
	Note string `json:"note"`

	// Type - no documentation
	Type *SoftLayer_User_External_Binding_Type `json:"type"`
}

func (softlayer_user_external_binding *SoftLayer_User_External_Binding) String() string {
	return "SoftLayer_User_External_Binding"
}
