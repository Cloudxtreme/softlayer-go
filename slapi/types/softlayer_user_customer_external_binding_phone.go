package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer_External_Binding_Phone - The SoftLayer_User_Customer_External_Binding_Phone
// data type contains information about an external binding that uses a phone call, SMS or mobile app
// for 2 form factor authentication. The external binding information is used when a SoftLayer customer
// logs into the SoftLayer customer portal or VPN to authenticate them against a trusted 3rd party, in
// this case using a mobile phone, mobile phone application or land-line phone. SoftLayer users with an
// active external binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Phone struct {

	// UserId - An external authentication binding's associated [[SoftLayer_User_Customer|user account]]
	// id.
	UserId int `json:"userId,omitempty"`

	// Password - The password used to authenticate the external id at an external authentication source.
	Password string `json:"password,omitempty"`

	// TypeId - The [[SoftLayer_User_External_Binding_Type|type]] identifier of an external authentication
	// binding.
	TypeId int `json:"typeId,omitempty"`

	// Id - An external authentication binding's internal identifier.
	Id int `json:"id,omitempty"`

	// ExternalId - The identifier used to identify this binding to an external authentication source.
	ExternalId string `json:"externalId,omitempty"`

	// CreateDate - The date that the external authentication binding was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// VendorId - The [[SoftLayer_User_External_Binding_Vendor|vendor]] identifier of an external
	// authentication binding.
	VendorId int `json:"vendorId,omitempty"`

	// Active - The flag that determines whether the external binding is active will be used for
	// authentication or not.
	Active bool `json:"active,omitempty"`
}

func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) String() string {
	return "SoftLayer_User_Customer_External_Binding_Phone"
}

// SoftLayer_User_Customer_External_Binding_Phone_Extended is SoftLayer_User_Customer_External_Binding_Phone with all maskable types.
type SoftLayer_User_Customer_External_Binding_Phone_Extended struct {
	SoftLayer_User_Customer_External_Binding_Phone

	// AttributeCount - A count of attributes of an external authentication binding.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// User - The SoftLayer user that the external authentication binding belongs to.
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// BindingStatus - The current external binding status. It can be or
	BindingStatus string `json:"bindingStatus,omitempty"`

	// Type - no documentation
	Type *SoftLayer_User_External_Binding_Type `json:"type,omitempty"`

	// Attributes - no documentation
	Attributes []*SoftLayer_User_External_Binding_Attribute `json:"attributes,omitempty"`

	// PinLength - <nil>
	PinLength string `json:"pinLength,omitempty"`

	// BillingItem - Information regarding the billing item for external authentication.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Note - An optional note for identifying the external binding.
	Note string `json:"note,omitempty"`

	// Vendor - no documentation
	Vendor *SoftLayer_User_External_Binding_Vendor `json:"vendor,omitempty"`
}

func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone_Extended) String() string {
	return "SoftLayer_User_Customer_External_Binding_Phone"
}
