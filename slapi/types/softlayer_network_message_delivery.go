package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Message_Delivery - <nil>
type SoftLayer_Network_Message_Delivery struct {

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Password - <nil>
	Password string `json:"password"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// TypeId - <nil>
	TypeId int `json:"typeId"`

	// Username - <nil>
	Username string `json:"username"`

	// VendorId - <nil>
	VendorId int `json:"vendorId"`
}

func (softlayer_network_message_delivery *SoftLayer_Network_Message_Delivery) String() string {
	return "SoftLayer_Network_Message_Delivery"
}

// SoftLayer_Network_Message_Delivery_Extended is SoftLayer_Network_Message_Delivery with all maskable types.
type SoftLayer_Network_Message_Delivery_Extended struct {
	SoftLayer_Network_Message_Delivery

	// Type - The message delivery type of a network message delivery account.
	Type *SoftLayer_Network_Message_Delivery_Type `json:"type"`

	// Vendor - no documentation
	Vendor *SoftLayer_Network_Message_Delivery_Vendor `json:"vendor"`

	// Account - The SoftLayer customer account that a network message delivery account belongs to.
	Account *SoftLayer_Account `json:"account"`

	// BillingItem - The billing item for a network message delivery account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`
}

func (softlayer_network_message_delivery *SoftLayer_Network_Message_Delivery_Extended) String() string {
	return "SoftLayer_Network_Message_Delivery"
}
