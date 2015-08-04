package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Message_Delivery - <nil>
type SoftLayer_Network_Message_Delivery struct {

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// TypeId - <nil>
	TypeId int `json:"typeId,omitempty"`

	// VendorId - <nil>
	VendorId int `json:"vendorId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Password - <nil>
	Password string `json:"password,omitempty"`

	// Username - <nil>
	Username string `json:"username,omitempty"`

	// Account - The SoftLayer customer account that a network message delivery account belongs to.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - The billing item for a network message delivery account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Type - The message delivery type of a network message delivery account.
	Type *SoftLayer_Network_Message_Delivery_Type `json:"type,omitempty"`

	// Vendor - no documentation
	Vendor *SoftLayer_Network_Message_Delivery_Vendor `json:"vendor,omitempty"`
}

func (softlayer_network_message_delivery *SoftLayer_Network_Message_Delivery) String() string {
	return "SoftLayer_Network_Message_Delivery"
}
