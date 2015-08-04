package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Message_Delivery_Email_Sendgrid - <nil>
type SoftLayer_Network_Message_Delivery_Email_Sendgrid struct {

	// Password - <nil>
	Password string `json:"password,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Username - <nil>
	Username string `json:"username,omitempty"`

	// VendorId - <nil>
	VendorId int `json:"vendorId,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// TypeId - <nil>
	TypeId int `json:"typeId,omitempty"`
}

func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid) String() string {
	return "SoftLayer_Network_Message_Delivery_Email_Sendgrid"
}

// SoftLayer_Network_Message_Delivery_Email_Sendgrid_Extended is SoftLayer_Network_Message_Delivery_Email_Sendgrid with all maskable types.
type SoftLayer_Network_Message_Delivery_Email_Sendgrid_Extended struct {
	SoftLayer_Network_Message_Delivery_Email_Sendgrid

	// EmailAddress - no documentation
	EmailAddress string `json:"emailAddress,omitempty"`

	// Account - The SoftLayer customer account that a network message delivery account belongs to.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItem - The billing item for a network message delivery account.
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// Vendor - no documentation
	Vendor *SoftLayer_Network_Message_Delivery_Vendor `json:"vendor,omitempty"`

	// SmtpAccess - A flag that determines if a SendGrid e-mail delivery account has access to send mail
	// through the SendGrid server.
	SmtpAccess string `json:"smtpAccess,omitempty"`

	// Type - The message delivery type of a network message delivery account.
	Type *SoftLayer_Network_Message_Delivery_Type `json:"type,omitempty"`
}

func (softlayer_network_message_delivery_email_sendgrid *SoftLayer_Network_Message_Delivery_Email_Sendgrid_Extended) String() string {
	return "SoftLayer_Network_Message_Delivery_Email_Sendgrid"
}
